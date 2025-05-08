// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {IPancakeRouter02} from "pancake-smart-contracts/contracts/interfaces/IPancakeRouter02.sol";

import {IBaseBridge} from "./interface/IBaseBridge.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";
import {ISwapBridgeRouter} from "./interface/ISwapBridgeRouter.sol";
import {Const} from "./lib/Const.sol";

contract SwapBridgeRouter is ISwapBridgeRouter, Ownable, Pausable {
    using Math for uint;

    error SwapBridgeBridgeFailed();
    error SwapBridgeInsufficientBalance();
    error SwapBridgeInvalidPath();
    error SwapBridgeInvalidBridgeToken();
    error SwapBridgeInvalidToken();
    error SwapBridgeNetworkFeeTooHigh();
    error SwapBridgeExFeeTooHigh();
    error SwapBridgeInvalidValue();
    error SwapBridgeOnlySwapRouter();

    event SwapBridgeInitiated(
        uint indexed toChainID, uint indexed index, address indexed from, address to, uint bridgeValue, address[] path
    );

    event TokenPathSet(address indexed token, address[] path);

    IBaseBridge public immutable bridge;
    IPancakeRouter02 public immutable swapRouter;
    address public immutable cross;
    uint public immutable crossChainID;
    mapping(address => address[]) public tokenToPath;

    constructor(address owner_, address bridge_, address swapRouter_, address cross_, uint crossChainID_)
        Ownable(owner_)
    {
        bridge = IBaseBridge(bridge_);
        swapRouter = IPancakeRouter02(swapRouter_);
        cross = cross_;
        crossChainID = crossChainID_;
    }

    receive() external payable {
        require(msg.sender == address(swapRouter), SwapBridgeOnlySwapRouter());
    }

    // ----------------
    // --- modifier ---
    // ----------------
    modifier validPath(uint toChainID, address[] memory path) {
        require(path.length > 1, SwapBridgeInvalidPath());
        address tokenForBridge = path[path.length - 1];
        if (tokenForBridge == swapRouter.WETH()) tokenForBridge = Const.NATIVE_TOKEN;
        IBaseBridge.TokenPair memory tokenPair = bridge.getTokenPair(toChainID, tokenForBridge);
        require(tokenPair.localToken == tokenForBridge, SwapBridgeInvalidBridgeToken());
        _;
    }

    modifier handleToken(address[] memory path, uint amount) {
        _receiveToken(IERC20(path[0]), amount);
        _;
        _refundToken(IERC20(path[0]));
        _refundToken(IERC20(path[path.length - 1]));
    }

    // -------------------------
    // --- execute functions ---
    // -------------------------
    // ------------- route 경로가 정해져있는 함수 -------------
    function swapExactTokensForCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external {
        address[] memory path = tokenToPath[tokenForSwap];
        require(path.length > 1, SwapBridgeInvalidPath());

        swapExactTokensForTokensBridge(
            crossChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline
        );
    }

    function swapTokensForExactCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external {
        address[] memory path = tokenToPath[tokenForSwap];
        require(path.length > 1, SwapBridgeInvalidPath());

        swapTokensForExactTokensBridge(
            crossChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline
        );
    }

    function swapExactETHForCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external payable {
        address[] memory path = tokenToPath[tokenForSwap];
        require(path.length > 1, SwapBridgeInvalidPath());

        swapExactETHForTokensBridge(crossChainID, to, amountOutMin, maxNetworkFee, maxExFee, path, deadline);
    }

    function swapETHForExactCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountOut,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external payable {
        address[] memory path = tokenToPath[tokenForSwap];
        require(path.length > 1, SwapBridgeInvalidPath());

        swapETHForExactTokensBridge(crossChainID, to, amountOut, maxNetworkFee, maxExFee, path, deadline);
    }

    // ------------- 기본 swap bridge 함수 -------------
    // 정확한 양의 토큰을 다른 토큰으로 교환 후 Bridge 요청
    function swapExactTokensForTokensBridge(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) public validPath(toChainID, path) handleToken(path, amountIn) {
        // swap 요청
        swapRouter.swapExactTokensForTokens(amountIn, amountOutMin, path, address(this), deadline);

        // bridge 요청
        _bridgeToken(toChainID, to, maxNetworkFee, maxExFee, path);
    }

    // 원하는 양의 토큰을 받기 위해 필요한 만큼의 토큰을 교환 후 Bridge 요청
    function swapTokensForExactTokensBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) public validPath(toChainID, path) handleToken(path, amountInMax) {
        amountOut = _calculateTotalAmount(toChainID, amountOut, path);

        // swap 요청
        swapRouter.swapTokensForExactTokens(amountOut, amountInMax, path, address(this), deadline);

        // bridge 요청
        _bridgeToken(toChainID, to, maxNetworkFee, maxExFee, path);
    }

    // ETH를 토큰으로 교환 후 Bridge 요청
    function swapExactETHForTokensBridge(
        uint toChainID,
        address to,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) public payable validPath(toChainID, path) handleToken(path, msg.value) {
        // swap 요청
        swapRouter.swapExactETHForTokens{value: msg.value}(amountOutMin, path, address(this), deadline);

        // bridge 요청
        _bridgeToken(toChainID, to, maxNetworkFee, maxExFee, path);
    }

    // 토큰을 원하는 양의 ETH로 교환 후 Bridge 요청
    function swapTokensForExactETHBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) public validPath(toChainID, path) handleToken(path, amountInMax) {
        amountOut = _calculateTotalAmount(toChainID, amountOut, path);

        // swap 요청
        swapRouter.swapTokensForExactETH(amountOut, amountInMax, path, address(this), deadline);

        // bridge 요청
        _bridgeToken(toChainID, to, maxNetworkFee, maxExFee, path);
    }

    // 토큰을 ETH로 교환 후 Bridge 요청
    function swapExactTokensForETHBridge(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) public payable validPath(toChainID, path) handleToken(path, amountIn) {
        // swap 요청
        swapRouter.swapExactTokensForETH(amountIn, amountOutMin, path, address(this), deadline);

        // bridge 요청
        _bridgeToken(toChainID, to, maxNetworkFee, maxExFee, path);
    }

    // ETH를 원하는 양의 토큰으로 교환 후 Bridge 요청
    function swapETHForExactTokensBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) public payable validPath(toChainID, path) handleToken(path, msg.value) {
        amountOut = _calculateTotalAmount(toChainID, amountOut, path);

        // swap 요청
        swapRouter.swapETHForExactTokens{value: msg.value}(amountOut, path, address(this), deadline);

        // bridge 요청
        _bridgeToken(toChainID, to, maxNetworkFee, maxExFee, path);
    }

    function setTokenPath(address token, address[] memory path) external onlyOwner {
        require(path.length > 1, SwapBridgeInvalidPath());
        require(path[0] == token, SwapBridgeInvalidPath());
        require(path[path.length - 1] == cross, SwapBridgeInvalidPath());

        tokenToPath[token] = path;
        emit TokenPathSet(token, path);
    }

    // ----------------------
    // --- view functions ---
    // ----------------------

    // Cross 토큰 스왑의 path 조회. 시작 token의 주소를 입력
    function getPath(address token) external view returns (address[] memory) {
        return tokenToPath[token];
    }

    // Cross 토큰 스왑의 출력 수량, 브릿지 가치, 네트워크 수수료, 추가 수수료 조회, 시작토큰주소와 amountIn을 입력
    function getSwapBridgeOutCross(address token, uint amountIn)
        external
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee)
    {
        address[] memory path = tokenToPath[token];
        (amounts, bridgeValue, networkFee, exFee) = getSwapBridgeOut(crossChainID, amountIn, path);
    }

    // Cross 토큰 스왑의 입력 수량, 네트워크 수수료, 추가 수수료 조회, 시작토큰주소와 amountOut을 입력
    function getSwapBridgeInCross(address token, uint amountOut)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee)
    {
        address[] memory path = tokenToPath[token];
        (amounts, networkFee, exFee) = getSwapBridgeIn(crossChainID, amountOut, path);
    }

    // bridge할 token의 수량 => swap + target chain에서 받을 token의 수량 계산
    function getSwapBridgeOut(uint toChainID, uint amountIn, address[] memory path)
        public
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee)
    {
        address token = path[path.length - 1];
        if (token == swapRouter.WETH()) token = Const.NATIVE_TOKEN;

        IBaseBridge.TokenPair memory tokenPair = bridge.getTokenPair(toChainID, token);
        require(tokenPair.localToken == token, SwapBridgeInvalidToken());

        amounts = swapRouter.getAmountsOut(amountIn, path);
        bool ok;
        (ok, bridgeValue, networkFee, exFee) = _estimateMaxValue(toChainID, IERC20(token), amounts[amounts.length - 1]);
        require(ok, SwapBridgeInsufficientBalance());
    }

    // // target chain에서 받을 token의 수량 => from chain에서 bridge할 token의 수량 계산
    function getSwapBridgeIn(uint toChainID, uint amountOut, address[] memory path)
        public
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee)
    {
        address token = path[path.length - 1];
        if (token == swapRouter.WETH()) token = Const.NATIVE_TOKEN;

        IBaseBridge.TokenPair memory tokenPair = bridge.getTokenPair(toChainID, token);
        require(tokenPair.localToken == token, SwapBridgeInvalidToken());

        uint minimumAmount;
        (minimumAmount, networkFee, exFee) = bridge.bridgeVerifier().calculateFee(toChainID, IERC20(token), amountOut);
        require(amountOut > minimumAmount, SwapBridgeInsufficientBalance());

        amounts = swapRouter.getAmountsIn(amountOut + networkFee + exFee, path);
    }

    function _receiveToken(IERC20 token, uint amount) private {
        amount = amount == 0 ? token.balanceOf(msg.sender) : amount;
        if (address(token) == swapRouter.WETH()) require(msg.value >= amount, SwapBridgeInvalidValue());
        else token.transferFrom(msg.sender, address(this), amount);

        if (token.allowance(address(this), address(swapRouter)) < amount) {
            token.approve(address(swapRouter), type(uint).max);
        }
    }

    function _refundToken(IERC20 token) private {
        if (address(token) == swapRouter.WETH()) {
            if (address(this).balance > 0) payable(msg.sender).transfer(address(this).balance);
        } else {
            uint refundAmount = token.balanceOf(address(this));
            if (refundAmount > 0) token.transfer(msg.sender, refundAmount);
        }
    }

    function _bridgeToken(uint toChainID, address to, uint maxNetworkFee, uint maxExFee, address[] memory path)
        private
    {
        IERC20 tokenForBridge = IERC20(path[path.length - 1]);
        uint swappedAmount;
        uint value;
        if (address(tokenForBridge) == swapRouter.WETH()) {
            tokenForBridge = IERC20(Const.NATIVE_TOKEN);
            swappedAmount = address(this).balance;
            value = swappedAmount;
        } else {
            swappedAmount = tokenForBridge.balanceOf(address(this));
            if (tokenForBridge.allowance(address(this), address(bridge)) < swappedAmount) {
                tokenForBridge.approve(address(bridge), type(uint).max);
            }
        }

        uint bridgeValue;
        uint networkFee;
        uint exFee;
        {
            bool ok;
            (ok, bridgeValue, networkFee, exFee) = _estimateMaxValue(toChainID, tokenForBridge, swappedAmount);
            require(ok, SwapBridgeInsufficientBalance());
            require(networkFee <= maxNetworkFee, SwapBridgeNetworkFeeTooHigh());
            require(exFee <= maxExFee, SwapBridgeExFeeTooHigh());
        }

        uint index = bridge.getNextInitiateIndex(toChainID);
        // bridge 요청
        require(
            bridge.bridgeToken{value: value}(
                toChainID, tokenForBridge, to, bridgeValue, networkFee, exFee, new bytes(0)
            ),
            SwapBridgeBridgeFailed()
        );

        emit SwapBridgeInitiated(toChainID, index, msg.sender, to, bridgeValue, path);
    }

    function _estimateMaxValue(uint toChainID, IERC20 token, uint totalValue)
        private
        view
        returns (bool ok, uint value, uint networkFee, uint exFee)
    {
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();

        uint minimumAmount;
        uint exFeeRate;
        (minimumAmount, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);
        if (totalValue <= networkFee) return (false, 0, 0, 0);

        uint denominator = bridgeVerifier.denominator();
        uint v = totalValue - networkFee;
        value = v.mulDiv(denominator, (denominator + exFeeRate));
        exFee = value.mulDiv(exFeeRate, denominator);
        ok = value > minimumAmount && totalValue >= value + networkFee + exFee;
    }

    function _calculateTotalAmount(uint toChainID, uint value, address[] memory path)
        private
        view
        returns (uint totalValue)
    {
        IERC20 token = IERC20(path[path.length - 1]);
        if (address(token) == swapRouter.WETH()) token = IERC20(Const.NATIVE_TOKEN);

        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();

        uint minimumAmount;
        uint exFeeRate;
        uint networkFee;

        (minimumAmount, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);
        require(value > minimumAmount, SwapBridgeInsufficientBalance());

        uint denominator = bridgeVerifier.denominator();

        // 계산 정확도 향상
        uint exFee = value.mulDiv(exFeeRate, denominator);
        totalValue = value + exFee + networkFee;

        return totalValue;
    }
}
