// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {ICrossBridge} from "./interface/ICrossBridge.sol";
import {Const} from "./lib/Const.sol";
import {ForwardLib} from "./lib/ForwardLib.sol";

/**
 * @title CrossBridge
 * @notice CROSS-chain relay-hub bridge: keeps the bridge's operational surface
 * (crossSupply accounting) and adds an executor-only forwarded entrypoint
 * (`bridgeTokenForwarded`) that enables multi-hop bridging (A -> B -> C).
 * @dev This implementation is the upgrade target for the `CrossBridge` proxy already
 * deployed and initialized on every CROSS-family chain. The currently-live bytecode on
 * chains not yet upgraded (which predates `bridgeTokenForwarded` and has no multi-hop
 * support) is frozen for reference at `test/legacy/CrossBridgeLegacy.sol`, and
 * `test/CrossBridgeUpgrade.t.sol` proves this contract is a layout-compatible upgrade
 * target for it.
 *
 * Inherits `BaseBridge` DIRECTLY, not the legacy contract preserved in
 * `test/legacy/CrossBridgeLegacy.sol`. That legacy contract's initializer, predeploy
 * address, and initializer-only errors are deployment-only concerns that would otherwise
 * be linked into this implementation for no runtime benefit, and removing them frees up
 * the code size this feature needs. This implementation is intended ONLY as an upgrade
 * target for an already-initialized `CrossBridge` proxy — there is no initialization path
 * here, `initialize` is permanently disabled below.
 *
 * Storage layout is redeclared IDENTICALLY to the legacy contract (same order, same
 * types, same gap size) so an upgrade from it to this contract is layout-compatible.
 * Do not reorder or resize these declarations.
 */
contract CrossBridge is BaseBridge, ICrossBridge {
    error Disabled();

    /**
     * @notice Emitted when the cross-chain supply limit is updated
     * @param crossSupplyLimit The new maximum supply limit for CROSS native token transfers
     */
    event CrossSupplyLimitSet(uint crossSupplyLimit);

    // --- storage layout identical to the legacy CrossBridge — do not reorder/resize ---

    /// @dev Maximum issuance limit for CROSS native token on the Cross chain
    uint public crossSupplyLimit;

    /// @dev BSC chain ID (56 for BNB Smart Chain mainnet, 97 for its testnet)
    uint private _bscChainID;

    /// @dev Storage gap for future upgrades
    uint[48] private __gap;

    // --- end legacy-CrossBridge-identical layout ---

    /**
     * @notice Disabled. `CrossBridge` is deployed only as an upgrade target for an
     * already-initialized `CrossBridge` proxy; there is no initialization path here.
     */
    function initialize(address, address payable, uint8) external pure override {
        revert Disabled();
    }

    /*
     * ---------------------------------------------------------------------------------
     * PRESERVED FOR REFERENCE ONLY — NOT CALLABLE. Do not uncomment without redeploying
     * a genesis CROSS chain: this selector is intentionally dropped from the runtime
     * because every CROSS-family chain (612044 / 612055 / 612088) is already deployed
     * and initialized, and there is no plan to deploy a new one.
     *
     * If a new CROSS chain genesis deploy is ever needed, this initializer MUST be
     * restored (uncommented, with `Const`, `_registerToken`, `_withdrawToken`,
     * `__BaseBridge_init` and the two errors below all still in scope) before that
     * genesis deployment, and removed/commented out again before the first upgrade off
     * of it — see the legacy `CrossBridge` contract's own warning at
     * `test/legacy/CrossBridgeLegacy.sol`.
     *
     * error CrossBridgeCanNotZeroAddress();
     * error CrossBridgeCanNotZero();
     *
     * /// @notice Initializes the CrossBridge contract
     * /// @dev Sets up the contract with initial configuration
     * /// - Calls the base initialization in BaseBridge
     * /// - Records the initial balance for CROSS token supply tracking
     * /// - Sets the initial CROSS token supply limit to 0
     * /// - Pairs native CROSS token with the CROSS ERC20 token on BSC chain
     * /// @param owner Address that will receive admin role
     * /// @param dev_ Address of the developer account for receiving fees
     * /// @param threshold_ Minimum number of validators required for validation
     * /// @param bscChainID The chain ID of the BSC Network
     * /// @param cross Address of the CROSS ERC20 token on BSC Network
     * /// @param crossInitialSupply Pre-minted supply of CROSS tokens for the CROSS Foundation
     * function initializeCrossBridge(
     *     address owner,
     *     address payable dev_,
     *     uint8 threshold_,
     *     uint bscChainID,
     *     address cross,
     *     uint crossInitialSupply
     * ) external initializer {
     *     require(bscChainID != 0, CrossBridgeCanNotZero());
     *     require(cross != address(0), CrossBridgeCanNotZeroAddress());
     *
     *     __BaseBridge_init(owner, dev_, threshold_);
     *
     *     // Register CROSS token as a token pair
     *     // This pairs the native CROSS token on this chain with the CROSS ERC20 token on BSC Network
     *     _registerToken(bscChainID, false, Const.NATIVE_TOKEN, cross);
     *     if (crossInitialSupply > 0) _withdrawToken(bscChainID, Const.NATIVE_TOKEN, crossInitialSupply);
     *
     *     _bscChainID = bscChainID;
     * }
     * ---------------------------------------------------------------------------------
     */

    /**
     * @notice Returns the net amount of CROSS bridged in from BSC and not yet returned
     * @dev Intentionally also sums amounts parked on a further hop's destination chain
     * (BSC -> CROSS -> chain C): those funds were minted against the BSC leg and have not
     * been repaid there yet, so they still count as outstanding CROSS supply until they are.
     */
    function crossSupply() public view returns (uint) {
        return _tokenPairs[_bscChainID][Const.NATIVE_TOKEN].minted;
    }

    /**
     * @notice Sets the maximum issuance limit for CROSS native token
     * @dev Only callable by admin role
     * @param _crossSupplyLimit New maximum issuance limit for CROSS native token
     */
    function setCrossSupplyLimit(uint _crossSupplyLimit) external onlyRole(Const.ADMIN_ROLE) {
        crossSupplyLimit = _crossSupplyLimit;
        emit CrossSupplyLimitSet(crossSupplyLimit);
    }

    /**
     * @notice Verifies if a finalization amount is within allowed limits
     * @dev Extends the base implementation with CROSS token issuance limit checks
     */
    function _checkFinalizeAmount(uint fromChainID, IERC20 token, uint value, bool retry)
        internal
        override
        returns (Const.FinalizeStatus status, bool delay)
    {
        if (address(token) == Const.NATIVE_TOKEN && fromChainID == _bscChainID) {
            if (crossSupply() + value > crossSupplyLimit) return (Const.FinalizeStatus.CrossSupplyLimitExceeded, true);
        }

        return super._checkFinalizeAmount(fromChainID, token, value, retry);
    }

    /**
     * @notice Executor-only entrypoint used to initiate hop-2+ of a multi-hop bridge
     * (A -> B -> C), invoked from within `BridgeExecutor.executeExtraCall` during this
     * chain's finalize of the prior hop.
     * @dev No modifiers: `fromToken`/`fromChainID` are never parameters — a malicious
     * origin-chain user could forge them if they were, but the verifier's signature only
     * attests to what chain A emitted, not that the values a caller supplies here are
     * self-consistent with it. They are only obtained from the trusted, transient forward
     * context staged by `_forwardBegin` earlier in the SAME finalize call.
     * Since the context is only known after consuming it, and Solidity modifiers run
     * before the function body, all validation happens in-body via internal helpers,
     * in the exact order below:
     *   1. `msg.sender == bridgeExecutor`
     *   2. acquire forward reentrancy guard
     *   3. consume forward context (reverts if inactive)
     *   4. not paused
     *   5. token registered / chain+token not paused
     *   6. `to` non-zero, extraData length bound
     *   7. native-pass-through invariant (native always allowed; ERC20 requires both-origin)
     *   8. exact amount equality against the staged context value
     *   9. initiate-amount / fee checks
     *   10. execute the hop-2+ bridge (from = bridgeExecutor)
     *   11. emit ForwardInitiated, release guard
     */
    function bridgeTokenForwarded(
        uint toChainID,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes calldata extraData
    ) external payable {
        require(
            address(bridgeExecutor) != address(0) && msg.sender == address(bridgeExecutor),
            BaseBridgeForwardNotExecutor()
        );

        ForwardLib.acquireGuard();

        (uint ctxFromChainID, address ctxToken, uint ctxValue, uint ctxIndex) = ForwardLib.consumeCtx();

        _requireNotPaused();
        _validateToken(toChainID, ctxToken);

        require(to != address(0), BaseBridgeCanNotZeroAddress());
        require(_maxExtraDataLength == 0 || extraData.length <= _maxExtraDataLength, BaseBridgeExtraDataTooLong());

        // Native-pass-through invariant: native tokens never mint/burn on any leg, so
        // pass-through is always safe. ERC20 wrapped tokens do mint/burn, so both the
        // from-chain and to-chain pairs must be origin pairs to avoid a mint -> burn
        // round trip.
        if (ctxToken != Const.NATIVE_TOKEN) {
            require(
                _tokenPairs[ctxFromChainID][ctxToken].isOrigin && _tokenPairs[toChainID][ctxToken].isOrigin,
                BaseBridgeForwardNotOrigin()
            );
        }

        // Exact equality (not `<=`): for native, the executor structurally forwards the
        // full staged value and `_initiateBridge` requires msg.value == value + fee, so
        // this is enforced either way. For ERC20, exact equality prevents a silent
        // partial forward and guarantees `remaining == 0` for the executor's accounting.
        require(
            value + networkFee + exFee == ctxValue,
            BaseBridgeForwardAmountMismatch(value + networkFee + exFee, ctxValue)
        );

        (networkFee, exFee) = _checkInitiateAmount(toChainID, IERC20(ctxToken), value, networkFee, exFee);

        _executeBridge(
            BridgeTokenArguments({
                toChainID: toChainID,
                fromToken: IERC20(ctxToken),
                from: address(bridgeExecutor),
                to: to,
                value: value,
                networkFee: networkFee,
                exFee: exFee,
                extraData: extraData
            })
        );

        ForwardLib.emitInitiated(
            ctxFromChainID, ctxIndex, toChainID, ctxToken, to, value, networkFee, exFee, keccak256(extraData)
        );

        ForwardLib.releaseGuard();
    }

    /**
     * @notice Stages the forward context when a finalize's extraData targets this
     * contract's own `bridgeTokenForwarded` entrypoint.
     * @dev See `BaseBridge._forwardBegin` for the general contract.
     */
    function _forwardBegin(
        uint fromChainID,
        uint index,
        address token,
        uint value,
        address targetContract,
        bytes4 methodID
    ) internal override returns (bool) {
        if (targetContract != address(this) || methodID != this.bridgeTokenForwarded.selector) return false;
        ForwardLib.setCtx(fromChainID, token, value, index);
        return true;
    }

    /**
     * @notice Resolves the staged forward context: classifies the outcome, clears any
     * residual context, and reports via `ForwardFailed` when it did not succeed.
     * @dev See `BaseBridge._forwardEnd` for the general contract.
     */
    function _forwardEnd(uint fromChainID, uint index, bool executorInvoked, bool ok, bytes memory result)
        internal
        override
    {
        ForwardLib.endAndReport(fromChainID, index, executorInvoked, ok, result);
    }
}
