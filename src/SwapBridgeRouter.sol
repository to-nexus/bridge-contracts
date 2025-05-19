// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {IPancakeFactory} from "pancake-smart-contracts/contracts/interfaces/IPancakeFactory.sol";
import {IPancakePair} from "pancake-smart-contracts/contracts/interfaces/IPancakePair.sol";
import {IPancakeRouter02} from "pancake-smart-contracts/contracts/interfaces/IPancakeRouter02.sol";

import {IBaseBridge} from "./interface/IBaseBridge.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";
import {ISwapBridgeRouter} from "./interface/ISwapBridgeRouter.sol";
import {Const} from "./lib/Const.sol";

/// @title SwapBridgeRouter
/// @notice A contract that combines token swapping and cross-chain bridging functionalities
/// @dev Uses PancakeSwap for token swapping and Bridge for cross-chain token transfers
contract SwapBridgeRouter is ISwapBridgeRouter, ReentrancyGuardTransient, Ownable, Pausable {
    using Math for uint;
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.AddressSet;

    error SwapBridgeBridgeFailed();
    error SwapBridgeInsufficientBalance();
    error SwapBridgeInvalidPath();
    error SwapBridgeInvalidBridgeToken();
    error SwapBridgeNetworkFeeTooHigh();
    error SwapBridgeExFeeTooHigh();
    error SwapBridgeBridgeValueTooLow();
    error SwapBridgeOnlySwapRouter();
    error SwapBridgeCanNotZeroAddress();
    error SwapBridgeCanNotZero();
    error SwapBridgePairDoesNotExist();
    error SwapBridgeInsufficientReserve();
    error SwapBridgeInvalidSourceToken();

    /// @notice Emitted when a swap and bridge operation is initiated
    /// @param toChainID The destination chain ID
    /// @param index The index of the bridge transaction
    /// @param from The address initiating the transaction
    /// @param to The recipient address on the destination chain
    /// @param path The swap path used for the transaction
    /// @param amounts The amounts of tokens used and received in the swap
    event SwapBridgeInitiated(
        uint indexed toChainID, uint indexed index, address indexed from, address to, address[] path, uint[] amounts
    );

    /// @notice Emitted when a token path is set or updated
    /// @param sourceToken The token for which the path is set
    /// @param path The swap path from the token to the bridge token
    event TokenPathSet(address indexed sourceToken, address[] path);

    /// @notice Emitted when a token path is removed
    /// @param sourceToken The token for which the path is removed
    event TokenPathRemoved(address indexed sourceToken);

    /// @notice The bridge contract used for cross-chain transfers
    IBaseBridge public immutable BRIDGE;

    /// @notice The swap router contract used for token swapping
    IPancakeRouter02 public immutable SWAP_ROUTER;

    /// @notice The address of the cross token
    address public immutable CROSS_TOKEN;

    /// @notice The ID of the destination chain for cross operations
    uint public immutable CROSS_CHAIN_ID;

    /// @notice The address of the wrapped native token
    address public immutable WETH;

    /// @notice Mapping from token address to its swap path to the cross token
    mapping(address => address[]) private _tokenToPath;
    EnumerableSet.AddressSet private _sourceTokens;

    /// @notice Constructor for the SwapBridgeRouter contract
    /// @param owner_ The owner address of the contract
    /// @param bridge_ The address of the bridge contract
    /// @param swapRouter_ The address of the swap router contract
    /// @param crossToken_ The address of the cross token
    /// @param crossChainID_ The ID of the destination chain for cross operations
    constructor(address owner_, address bridge_, address swapRouter_, address crossToken_, uint crossChainID_)
        Ownable(owner_)
    {
        require(owner_ != address(0), SwapBridgeCanNotZeroAddress());
        require(bridge_ != address(0), SwapBridgeCanNotZeroAddress());
        require(swapRouter_ != address(0), SwapBridgeCanNotZeroAddress());
        require(crossToken_ != address(0), SwapBridgeCanNotZeroAddress());
        require(crossChainID_ != 0, SwapBridgeCanNotZero());

        BRIDGE = IBaseBridge(bridge_);
        SWAP_ROUTER = IPancakeRouter02(swapRouter_);
        CROSS_TOKEN = crossToken_;
        CROSS_CHAIN_ID = crossChainID_;
        WETH = SWAP_ROUTER.WETH();
    }

    /// @notice Receives ETH from the swap router
    /// @dev Only allows the swap router to send ETH to this contract
    receive() external payable {
        require(msg.sender == address(SWAP_ROUTER), SwapBridgeOnlySwapRouter());
    }

    // ----------------
    // --- modifier ---
    // ----------------
    /// @notice Validates that the swap path is valid for the given destination chain
    /// @param toChainID The destination chain ID
    /// @param path The swap path to validate
    modifier validPath(uint toChainID, address[] memory path) {
        require(path.length >= 2, SwapBridgeInvalidPath());
        address tokenForBridge = path[path.length - 1];
        // Convert WETH to native token if necessary since bridge handles native token differently from wrapped version
        if (tokenForBridge == WETH) tokenForBridge = Const.NATIVE_TOKEN;
        // Verify the token pair exists in the bridge for the destination chain
        IBaseBridge.TokenPair memory tokenPair = BRIDGE.getTokenPair(toChainID, tokenForBridge);
        require(tokenPair.localToken == tokenForBridge, SwapBridgeInvalidBridgeToken());
        _;
    }

    /// @notice Handles token transfers for swap and bridge operations
    /// @param path The swap path
    /// @param amount The amount of tokens to handle
    modifier handleToken(address[] memory path, uint amount) {
        // First receive tokens from user to this contract
        _receiveToken(IERC20(path[0]), amount);
        _;
        // After the operation, refund any unused tokens back to the user
        // Refund both input and output tokens to handle cases where not all tokens were used
        _refundToken(IERC20(path[0]));
        _refundToken(IERC20(path[path.length - 1]));
    }

    /// @notice Handles token transfers for swap and bridge operations to cross tokens
    /// @param sourceToken The token to handle
    /// @param amount The amount of tokens to handle
    modifier handleTokenCross(address sourceToken, uint amount) {
        address[] memory path = _tokenToPath[sourceToken];
        require(path.length >= 2, SwapBridgeInvalidPath());
        _receiveToken(IERC20(sourceToken), amount);
        _;
        _refundToken(IERC20(sourceToken));
        _refundToken(IERC20(CROSS_TOKEN));
    }

    // -------------------------
    // --- execute functions ---
    // -------------------------

    // ------------- Functions with predefined route paths for swapping to Cross tokens and bridging -------------
    /// @notice Swaps exact tokens for cross tokens and bridges them to the destination chain
    /// @param sourceToken The token to swap
    /// @param to The recipient address on the destination chain
    /// @param amountIn The amount of tokens to swap
    /// @param amountOutMin The minimum amount of tokens to receive from the swap
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param deadline The deadline for the swap
    function swapBridgeExactTokensForCrossTokens(
        address sourceToken,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external nonReentrant handleTokenCross(sourceToken, amountIn) {
        address[] memory path = _tokenToPath[sourceToken];
        _executeSwapAndBridge(
            CROSS_CHAIN_ID,
            to,
            amountIn,
            amountOutMin,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.EXACT_TOKEN_FOR_TOKEN
        );
    }

    /// @notice Swaps tokens for exact cross tokens and bridges them to the destination chain
    /// @param sourceToken The token to swap
    /// @param to The recipient address on the destination chain
    /// @param amountOut The exact amount of tokens to receive
    /// @param amountInMax The maximum amount of tokens to spend
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param deadline The deadline for the swap
    function swapBridgeTokensForExactCrossTokens(
        address sourceToken,
        address to,
        uint amountOut,
        uint amountInMax,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external nonReentrant handleTokenCross(sourceToken, amountInMax) {
        address[] memory path = _tokenToPath[sourceToken];
        _executeSwapAndBridge(
            CROSS_CHAIN_ID,
            to,
            amountInMax,
            amountOut,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.TOKEN_FOR_EXACT_TOKEN
        );
    }

    /// @notice Swaps exact ETH for cross tokens and bridges them to the destination chain
    /// @param to The recipient address on the destination chain
    /// @param amountOutMin The minimum amount of tokens to receive from the swap
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param deadline The deadline for the swap
    function swapBridgeExactETHForCrossTokens(
        address to,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external payable nonReentrant handleTokenCross(WETH, msg.value) {
        address[] memory path = _tokenToPath[WETH];
        _executeSwapAndBridge(
            CROSS_CHAIN_ID,
            to,
            msg.value,
            amountOutMin,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.EXACT_ETH_FOR_TOKEN
        );
    }

    /// @notice Swaps ETH for exact cross tokens and bridges them to the destination chain
    /// @param to The recipient address on the destination chain
    /// @param amountOut The exact amount of tokens to receive
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param deadline The deadline for the swap
    function swapBridgeETHForExactCrossTokens(
        address to,
        uint amountOut,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external payable nonReentrant handleTokenCross(WETH, msg.value) {
        address[] memory path = _tokenToPath[WETH];
        _executeSwapAndBridge(
            CROSS_CHAIN_ID,
            to,
            msg.value,
            amountOut,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.ETH_FOR_EXACT_TOKEN
        );
    }

    // ------------- Basic swap bridge functions -------------
    /// @notice Swaps exact tokens for tokens and bridges them to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountIn The amount of tokens to swap
    /// @param amountOutMin The minimum amount of tokens to receive from the swap. This amount is the amount after deducting the network fee and exchange fee.
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param deadline The deadline for the swap
    function swapBridgeExactTokensForTokens(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) public nonReentrant validPath(toChainID, path) handleToken(path, amountIn) {
        _executeSwapAndBridge(
            toChainID,
            to,
            amountIn,
            amountOutMin,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.EXACT_TOKEN_FOR_TOKEN
        );
    }

    /// @notice Swaps tokens for exact tokens and bridges them to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountOut The exact amount of tokens to receive
    /// @param amountInMax The maximum amount of tokens to spend
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param deadline The deadline for the swap
    function swapBridgeTokensForExactTokens(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) public nonReentrant validPath(toChainID, path) handleToken(path, amountInMax) {
        _executeSwapAndBridge(
            toChainID,
            to,
            amountInMax,
            amountOut,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.TOKEN_FOR_EXACT_TOKEN
        );
    }

    /// @notice Swaps exact ETH for tokens and bridges them to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountOutMin The minimum amount of tokens to receive from the swap
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param deadline The deadline for the swap
    function swapBridgeExactETHForTokens(
        uint toChainID,
        address to,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) public payable nonReentrant validPath(toChainID, path) handleToken(path, msg.value) {
        _executeSwapAndBridge(
            toChainID,
            to,
            msg.value,
            amountOutMin,
            networkFeeMax,
            exFeeMax,
            path,
            deadline,
            SwapType.EXACT_ETH_FOR_TOKEN
        );
    }

    /// @notice Swaps tokens for exact ETH and bridges them to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountOut The exact amount of ETH to receive
    /// @param amountInMax The maximum amount of tokens to spend
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param deadline The deadline for the swap
    function swapBridgeTokensForExactETH(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) public nonReentrant validPath(toChainID, path) handleToken(path, amountInMax) {
        _executeSwapAndBridge(
            toChainID, to, amountInMax, amountOut, networkFeeMax, exFeeMax, path, deadline, SwapType.TOKEN_FOR_EXACT_ETH
        );
    }

    /// @notice Swaps exact tokens for ETH and bridges them to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountIn The amount of tokens to swap
    /// @param amountOutMin The minimum amount of ETH to receive from the swap
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param deadline The deadline for the swap
    function swapBridgeExactTokensForETH(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) public payable nonReentrant validPath(toChainID, path) handleToken(path, amountIn) {
        _executeSwapAndBridge(
            toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline, SwapType.EXACT_TOKEN_FOR_ETH
        );
    }

    /// @notice Swaps ETH for exact tokens and bridges them to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountOut The exact amount of tokens to receive
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param deadline The deadline for the swap
    function swapBridgeETHForExactTokens(
        uint toChainID,
        address to,
        uint amountOut,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) public payable nonReentrant validPath(toChainID, path) handleToken(path, msg.value) {
        _executeSwapAndBridge(
            toChainID, to, msg.value, amountOut, networkFeeMax, exFeeMax, path, deadline, SwapType.ETH_FOR_EXACT_TOKEN
        );
    }

    // -----------------------
    // --- admin functions ---
    // -----------------------

    /// @notice Sets the swap path for a token
    /// @param sourceToken The token to set the path for
    /// @param path The swap path from the token to the cross token
    /// @dev Only the owner can call this function
    function setTokenPath(address sourceToken, address[] memory path) external onlyOwner {
        require(path.length >= 2, SwapBridgeInvalidPath());
        // First token in path must be the token we're starting with
        require(path[0] == sourceToken, SwapBridgeInvalidPath());
        // Last token in path must be the cross token used for bridging
        require(path[path.length - 1] == CROSS_TOKEN, SwapBridgeInvalidPath());

        _tokenToPath[sourceToken] = path;
        if (!_sourceTokens.contains(sourceToken)) _sourceTokens.add(sourceToken);

        emit TokenPathSet(sourceToken, path);
    }

    /// @notice Removes the swap path for a token
    /// @param sourceToken The token for which the path is removed
    /// @dev Only the owner can call this function
    function removeTokenPath(address sourceToken) external onlyOwner {
        require(_sourceTokens.contains(sourceToken), SwapBridgeInvalidSourceToken());
        delete _tokenToPath[sourceToken];
        _sourceTokens.remove(sourceToken);
        emit TokenPathRemoved(sourceToken);
    }

    /// @notice Pauses the contract
    /// @dev Only the owner can call this function
    function pause() external onlyOwner {
        _pause();
    }

    /// @notice Unpauses the contract
    /// @dev Only the owner can call this function
    function unpause() external onlyOwner {
        _unpause();
    }

    // ----------------------
    // --- view functions ---
    // ----------------------

    /// @notice Gets the list of source tokens
    /// @return The list of source tokens
    function getSourceTokens() external view returns (address[] memory) {
        return _sourceTokens.values();
    }

    /// @notice Gets the swap path for a token to the cross token
    /// @param token The token to get the path for
    /// @return The swap path from the token to the cross token
    function getPath(address token) external view returns (address[] memory) {
        return _tokenToPath[token];
    }

    /// @notice Gets output amounts, bridge value, network fee, and exchange fee for a cross token swap
    /// @param token The token to swap
    /// @param amountIn The amount of tokens to swap
    /// @return amounts The array of output amounts for each step in the swap
    /// @return bridgeValue The amount of tokens being bridged
    /// @return networkFee The network fee for the bridge
    /// @return exFee The exchange fee for the bridge
    /// @return priceImpactBps The price impact in basis points
    function getSwapBridgeOutCross(address token, uint amountIn)
        external
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee, uint priceImpactBps)
    {
        require(_sourceTokens.contains(token), SwapBridgeInvalidSourceToken());
        address[] memory path = _tokenToPath[token];
        (amounts, bridgeValue, networkFee, exFee, priceImpactBps) = getSwapBridgeOut(CROSS_CHAIN_ID, amountIn, path);
    }

    /// @notice Gets input amounts, network fee, and exchange fee for a cross token swap
    /// @param token The token to swap
    /// @param amountOut The amount of tokens to receive
    /// @return amounts The array of input amounts for each step in the swap
    /// @return networkFee The network fee for the bridge
    /// @return exFee The exchange fee for the bridge
    /// @return priceImpactBps The price impact in basis points
    function getSwapBridgeInCross(address token, uint amountOut)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee, uint priceImpactBps)
    {
        require(_sourceTokens.contains(token), SwapBridgeInvalidSourceToken());
        address[] memory path = _tokenToPath[token];
        (amounts, networkFee, exFee, priceImpactBps) = getSwapBridgeIn(CROSS_CHAIN_ID, amountOut, path);
    }

    /// @notice Calculates output amounts, bridge value, network fee, and exchange fee for a token swap and bridge
    /// @param toChainID The destination chain ID
    /// @param amountIn The amount of tokens to swap
    /// @param path The swap path
    /// @return amounts The array of output amounts for each step in the swap
    /// @return bridgeValue The amount of tokens being bridged
    /// @return networkFee The network fee for the bridge
    /// @return exFee The exchange fee for the bridge
    /// @return priceImpactBps The price impact in basis points
    function getSwapBridgeOut(uint toChainID, uint amountIn, address[] memory path)
        public
        view
        validPath(toChainID, path)
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee, uint priceImpactBps)
    {
        address token = path[path.length - 1];
        // Handle WETH as native token for bridge operations
        if (token == WETH) token = Const.NATIVE_TOKEN;

        // Get the output amounts for each step in the swap path
        amounts = getAmountsOut(amountIn, path);

        bool ok;
        // Calculate how much of the swapped token can actually be bridged after accounting for fees
        (ok, bridgeValue, networkFee, exFee) = _estimateMaxValue(toChainID, IERC20(token), amounts[amounts.length - 1]);
        require(ok, SwapBridgeInsufficientBalance());

        priceImpactBps = getPriceImpact(path, amountIn);
    }

    /// @notice Calculates input amounts, network fee, and exchange fee for a token swap and bridge
    /// @param toChainID The destination chain ID
    /// @param amountOut The amount of tokens to receive on the destination chain
    /// @param path The swap path
    /// @return amounts The array of input amounts for each step in the swap
    /// @return networkFee The network fee for the bridge
    /// @return exFee The exchange fee for the bridge
    /// @return priceImpactBps The price impact in basis points
    function getSwapBridgeIn(uint toChainID, uint amountOut, address[] memory path)
        public
        view
        validPath(toChainID, path)
        returns (uint[] memory amounts, uint networkFee, uint exFee, uint priceImpactBps)
    {
        address token = path[path.length - 1];
        // Handle WETH as native token for bridge operations
        if (token == WETH) token = Const.NATIVE_TOKEN;

        uint minimumAmount;
        // Get fee information for this token and destination chain
        (minimumAmount, networkFee, exFee) = BRIDGE.bridgeVerifier().calculateFee(toChainID, IERC20(token), amountOut);
        // Ensure the requested amount is more than the minimum bridgeable amount
        require(amountOut > minimumAmount, SwapBridgeInsufficientBalance());

        // Calculate how many input tokens are needed to get the desired output amount plus fees
        // This accounts for all fees to ensure the user receives exactly amountOut after bridging
        amounts = getAmountsIn(amountOut + networkFee + exFee, path);
        priceImpactBps = getPriceImpact(path, amounts[0]);
    }

    // ----------------------------------------------------------------------
    // --- PancakeSwap Router Integration and Price Calculation Utilities ---
    // ----------------------------------------------------------------------
    function getReserves(address tokenA, address tokenB) public view returns (uint112 reserveA, uint112 reserveB) {
        address pair = IPancakeFactory(SWAP_ROUTER.factory()).getPair(tokenA, tokenB);
        require(pair != address(0), SwapBridgePairDoesNotExist());

        IPancakePair pairContract = IPancakePair(pair);
        (uint112 r0, uint112 r1,) = pairContract.getReserves();

        (reserveA, reserveB) = tokenA == pairContract.token0() ? (r0, r1) : (r1, r0);
    }

    function quote(uint amountA, uint reserveA, uint reserveB) external view returns (uint amountB) {
        return SWAP_ROUTER.quote(amountA, reserveA, reserveB);
    }

    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external view returns (uint amountOut) {
        return SWAP_ROUTER.getAmountOut(amountIn, reserveIn, reserveOut);
    }

    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) external view returns (uint amountIn) {
        return SWAP_ROUTER.getAmountIn(amountOut, reserveIn, reserveOut);
    }

    function getAmountsOut(uint amountIn, address[] memory path) public view returns (uint[] memory amounts) {
        return SWAP_ROUTER.getAmountsOut(amountIn, path);
    }

    function getAmountsIn(uint amountOut, address[] memory path) public view returns (uint[] memory amounts) {
        return SWAP_ROUTER.getAmountsIn(amountOut, path);
    }

    /// @notice Calculates the price impact of a swap in basis points (BPS)
    /// @dev 1 BPS = 0.01%, 100 BPS = 1%
    /// @param path The swap path
    /// @param amountIn The amount of input tokens
    /// @return priceImpactBps The price impact in basis points
    function getPriceImpact(address[] memory path, uint amountIn) public view returns (uint priceImpactBps) {
        if (path.length < 2 || amountIn == 0) return 0;

        // Get the amount out for the swap
        uint[] memory amounts = getAmountsOut(amountIn, path);

        // Calculate price impact for each hop in the path
        for (uint i = 0; i < path.length - 1; i++) {
            address tokenIn = path[i];
            address tokenOut = path[i + 1];

            // Get current reserves
            (uint112 reserveIn, uint112 reserveOut) = getReserves(tokenIn, tokenOut);
            require(reserveIn > 0 && reserveOut > 0, SwapBridgeInsufficientReserve());

            // Price before swap (reserveOut/reserveIn)
            uint priceBeforeSwap = uint(reserveOut) * 1e18 / uint(reserveIn);

            // Input and output amounts for the swap
            uint hopAmountIn = amounts[i];
            uint hopAmountOut = amounts[i + 1];

            // Expected price after swap ((reserveOut - hopAmountOut) / (reserveIn + hopAmountIn))
            uint priceAfterSwap = uint(reserveOut - hopAmountOut) * 1e18 / uint(reserveIn + hopAmountIn);

            // Calculate price impact: (price before swap - price after swap) / price before swap * 10000
            uint hopImpact;
            if (priceBeforeSwap > priceAfterSwap) {
                hopImpact = ((priceBeforeSwap - priceAfterSwap) * 10000) / priceBeforeSwap;
            }

            // Accumulate price impact across hops
            priceImpactBps += hopImpact;
        }
    }

    // --------------------------
    // --- internal functions ---
    // --------------------------
    /// @notice Receives tokens from the sender
    /// @param token The token to receive
    /// @param amount The amount of tokens to receive
    /// @dev If amount is 0, receives the sender's entire balance
    function _receiveToken(IERC20 token, uint amount) private {
        // If amount is 0, use the sender's entire token balance
        amount = amount == 0 ? token.balanceOf(msg.sender) : amount;

        // Check if the token is WETH and if the transaction includes enough ETH
        // If it's WETH and enough ETH is provided, we don't need to transfer tokens
        if (!(address(token) == WETH && msg.value >= amount)) token.safeTransferFrom(msg.sender, address(this), amount);

        // Approve the swap router to spend the tokens if needed
        if (token.allowance(address(this), address(SWAP_ROUTER)) < amount) {
            token.approve(address(SWAP_ROUTER), type(uint).max);
        }
    }

    /// @notice Refunds any remaining tokens to the sender
    /// @param token The token to refund
    function _refundToken(IERC20 token) private {
        // For ETH, transfer any remaining balance back to the sender
        if (address(this).balance > 0) payable(msg.sender).transfer(address(this).balance);

        // refund any remaining balance
        uint refundAmount = token.balanceOf(address(this));
        if (refundAmount > 0) token.safeTransfer(msg.sender, refundAmount);
    }

    /// @notice Helper function to reduce code duplication across swap types
    /// @dev Executes appropriate swap based on SwapType and performs bridge operation
    function _executeSwapAndBridge(
        uint toChainID,
        address to,
        uint amount,
        uint amountOutOrMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline,
        SwapType swapType
    ) private {
        uint amountSwapOutMin;
        uint amountOut;
        uint[] memory amounts;

        // Execute the appropriate swap based on swap type
        if (swapType == SwapType.EXACT_TOKEN_FOR_TOKEN) {
            amountSwapOutMin = amountOutOrMin + networkFeeMax + exFeeMax;
            amounts = SWAP_ROUTER.swapExactTokensForTokens(amount, amountSwapOutMin, path, address(this), deadline);
        } else if (swapType == SwapType.EXACT_ETH_FOR_TOKEN) {
            amountSwapOutMin = amountOutOrMin + networkFeeMax + exFeeMax;
            amounts = SWAP_ROUTER.swapExactETHForTokens{value: amount}(amountSwapOutMin, path, address(this), deadline);
        } else if (swapType == SwapType.EXACT_TOKEN_FOR_ETH) {
            amountSwapOutMin = amountOutOrMin + networkFeeMax + exFeeMax;
            amounts = SWAP_ROUTER.swapExactTokensForETH(amount, amountSwapOutMin, path, address(this), deadline);
        } else if (swapType == SwapType.TOKEN_FOR_EXACT_TOKEN) {
            amountOut = _calculateTotalAmount(toChainID, amountOutOrMin, path);
            amounts = SWAP_ROUTER.swapTokensForExactTokens(amountOut, amount, path, address(this), deadline);
        } else if (swapType == SwapType.ETH_FOR_EXACT_TOKEN) {
            amountOut = _calculateTotalAmount(toChainID, amountOutOrMin, path);
            amounts = SWAP_ROUTER.swapETHForExactTokens{value: amount}(amountOut, path, address(this), deadline);
        } else if (swapType == SwapType.TOKEN_FOR_EXACT_ETH) {
            amountOut = _calculateTotalAmount(toChainID, amountOutOrMin, path);
            amounts = SWAP_ROUTER.swapTokensForExactETH(amountOut, amount, path, address(this), deadline);
        }

        // Bridge the tokens to the destination chain
        _bridgeToken(toChainID, to, amountOutOrMin, networkFeeMax, exFeeMax, path, amounts);
    }

    /// @notice Bridges tokens to the destination chain
    /// @param toChainID The destination chain ID
    /// @param to The recipient address on the destination chain
    /// @param amountOutMin The minimum amount to receive after bridging
    /// @param networkFeeMax The maximum network fee to pay
    /// @param exFeeMax The maximum exchange fee to pay
    /// @param path The swap path
    /// @param amounts The amounts of tokens received after each step in the swap
    function _bridgeToken(
        uint toChainID,
        address to,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint[] memory amounts
    ) private {
        bool isWETH;
        FeeData memory feeData;
        IERC20 tokenForBridge;

        {
            // Optimize local variables to reduce stack usage
            address lastToken = path[path.length - 1];
            isWETH = lastToken == WETH;
            uint swappedAmount;

            // Handle token type and balance
            (tokenForBridge, swappedAmount) = isWETH
                ? (IERC20(Const.NATIVE_TOKEN), address(this).balance)
                : (IERC20(lastToken), IERC20(lastToken).balanceOf(address(this)));

            // Approve bridge contract if needed
            if (!isWETH && tokenForBridge.allowance(address(this), address(BRIDGE)) < swappedAmount) {
                tokenForBridge.approve(address(BRIDGE), type(uint).max);
            }

            // Calculate fees and validate constraints
            feeData = _calculateFees(toChainID, tokenForBridge, swappedAmount, amountOutMin, networkFeeMax, exFeeMax);
        }

        // Execute bridge operation
        uint index = BRIDGE.getNextInitiateIndex(toChainID);

        require(
            BRIDGE.bridgeToken{value: isWETH ? feeData.bridgeValue : 0}(
                toChainID, tokenForBridge, to, feeData.bridgeValue, feeData.networkFee, feeData.exFee, new bytes(0)
            ),
            SwapBridgeBridgeFailed()
        );

        emit SwapBridgeInitiated(toChainID, index, msg.sender, to, path, amounts);
    }

    /// @notice Calculates and validates fee data for bridge operation
    /// @param toChainID Target chain ID for the bridge
    /// @param token Token to bridge
    /// @param amount Total token amount available
    /// @param amountOutMin Minimum output amount required
    /// @param networkFeeMax Maximum network fee allowed
    /// @param exFeeMax Maximum exchange fee allowed
    /// @return feeData Calculated fee data struct
    function _calculateFees(
        uint toChainID,
        IERC20 token,
        uint amount,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax
    ) private view returns (FeeData memory feeData) {
        bool ok;
        (ok, feeData.bridgeValue, feeData.networkFee, feeData.exFee) = _estimateMaxValue(toChainID, token, amount);

        require(ok, SwapBridgeInsufficientBalance());
        require(feeData.bridgeValue >= amountOutMin, SwapBridgeBridgeValueTooLow());
        require(feeData.networkFee <= networkFeeMax, SwapBridgeNetworkFeeTooHigh());
        require(feeData.exFee <= exFeeMax, SwapBridgeExFeeTooHigh());

        return feeData;
    }

    /// @notice Estimates the maximum value, network fee, and exchange fee for a bridge operation
    /// @param toChainID The destination chain ID
    /// @param token The token to bridge
    /// @param totalValue The total value to bridge
    /// @return ok Whether the estimation was successful
    /// @return value The amount of tokens being bridged
    /// @return networkFee The network fee for the bridge
    /// @return exFee The exchange fee for the bridge
    function _estimateMaxValue(uint toChainID, IERC20 token, uint totalValue)
        private
        view
        returns (bool ok, uint value, uint networkFee, uint exFee)
    {
        IBridgeVerifier bridgeVerifier = BRIDGE.bridgeVerifier();

        uint minimumAmount;
        uint exFeeRate;
        // Get token configuration for this chain including minimum amount, network fee, and exchange fee rate
        (minimumAmount, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);

        // If total value is less than or equal to the network fee, we can't bridge anything
        if (totalValue <= networkFee) return (false, 0, 0, 0);

        uint denominator = bridgeVerifier.denominator();
        // Calculate available value after network fee
        uint v = totalValue - networkFee;

        // Calculate actual bridge value after accounting for exchange fee
        // Formula: value = (totalValue - networkFee) * denominator / (denominator + exFeeRate)
        value = v.mulDiv(denominator, (denominator + exFeeRate));

        // Calculate exchange fee based on the bridge value
        // Formula: exFee = value * exFeeRate / denominator
        exFee = value.mulDiv(exFeeRate, denominator);

        // Verify that the bridge value exceeds minimum amount and total fees + value don't exceed available balance
        ok = value > minimumAmount && totalValue >= value + networkFee + exFee;
    }

    /// @notice Calculates the total amount required for a bridge operation
    /// @param toChainID The destination chain ID
    /// @param value The amount of tokens to bridge
    /// @param path The swap path
    /// @return totalValue The total amount required, including fees
    function _calculateTotalAmount(uint toChainID, uint value, address[] memory path)
        private
        view
        returns (uint totalValue)
    {
        IERC20 token = IERC20(path[path.length - 1]);
        // Convert WETH to native token if necessary
        if (address(token) == WETH) token = IERC20(Const.NATIVE_TOKEN);

        IBridgeVerifier bridgeVerifier = BRIDGE.bridgeVerifier();

        uint minimumAmount;
        uint exFeeRate;
        uint networkFee;

        // Get token configuration for this chain
        (minimumAmount, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);
        // Ensure the requested value exceeds the minimum bridgeable amount
        require(value >= minimumAmount, SwapBridgeInsufficientBalance());

        uint denominator = bridgeVerifier.denominator();

        // Calculate exchange fee based on the bridge value
        // Formula: exFee = value * exFeeRate / denominator
        uint exFee = value.mulDiv(exFeeRate, denominator);

        // Calculate total amount needed: the bridge value + exchange fee + network fee
        totalValue = value + exFee + networkFee;
    }
}
