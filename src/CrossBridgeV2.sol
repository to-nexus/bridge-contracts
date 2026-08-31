// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {ICrossBridgeV2} from "./interface/ICrossBridgeV2.sol";
import {Const} from "./lib/Const.sol";
import {ForwardLib} from "./lib/ForwardLib.sol";

/**
 * @title CrossBridgeV2
 * @notice CROSS-chain relay-hub bridge: keeps `CrossBridge`'s operational surface
 * (crossSupply accounting) and adds an executor-only forwarded entrypoint
 * (`bridgeTokenForwarded`) that enables multi-hop bridging (A -> B -> C).
 * @dev Inherits `BaseBridge` DIRECTLY, NOT `CrossBridge`. `CrossBridge`'s initializer,
 * predeploy address, and initializer-only errors are deployment-only concerns that
 * would otherwise be linked into this implementation for no runtime benefit, and
 * removing them frees up the code size this feature needs (plan spec §6.2). This
 * implementation is intended ONLY as an upgrade target for an already-initialized
 * `CrossBridge` proxy — there is no initialization path here, `initialize` is
 * permanently disabled below.
 *
 * Storage layout is redeclared IDENTICALLY to `CrossBridge` (same order, same types,
 * same gap size) so an upgrade from `CrossBridge` to `CrossBridgeV2` is
 * layout-compatible. Do not reorder or resize these declarations.
 */
contract CrossBridgeV2 is BaseBridge, ICrossBridgeV2 {
    error Disabled();

    /**
     * @notice Emitted when the cross-chain supply limit is updated
     * @param crossSupplyLimit The new maximum supply limit for CROSS native token transfers
     */
    event CrossSupplyLimitSet(uint crossSupplyLimit);

    // --- storage layout identical to CrossBridge — do not reorder/resize ---

    /// @dev Maximum issuance limit for CROSS native token on the Cross chain
    uint public crossSupplyLimit;

    /// @dev BSC chain ID (e.g., 56 for mainnet, 97 for Sepolia testnet, or other chain IDs)
    uint private _bscChainID;

    /// @dev Storage gap for future upgrades
    uint[48] private __gap;

    // --- end CrossBridge-identical layout ---

    /**
     * @notice Disabled. `CrossBridgeV2` is deployed only as an upgrade target for an
     * already-initialized `CrossBridge` proxy; there is no initialization path here.
     */
    function initialize(address, address payable, uint8) external pure override {
        revert Disabled();
    }

    /**
     * @notice Returns the net amount of CROSS bridged in from BSC and not yet returned
     * @dev See BSC->CROSS->chainC pass-through accounting note in the plan spec (§6.1):
     * this intentionally also sums amounts parked on a further hop's destination chain,
     * since they were minted from the BSC leg and have not been repaid there yet.
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
     * origin-chain user could forge them if they were, but the verifier's signature
     * only attests to what chain A emitted, not that it is meaningfully self-consistent
     * (plan spec Assumptions). They are only obtained from the trusted, transient
     * forward context staged by `_forwardBegin` earlier in the SAME finalize call.
     * Since the context is only known after consuming it, and Solidity modifiers run
     * before the function body, all validation happens in-body via internal helpers,
     * in the exact order below (plan spec §7.3):
     *   1. `msg.sender == bridgeExecutor`
     *   2. acquire forward reentrancy guard
     *   3. consume forward context (reverts if inactive)
     *   4. not paused
     *   5. token registered / chain+token not paused
     *   6. `to` non-zero, extraData length bound
     *   7. D1 invariant (native always allowed; ERC20 requires both-origin)
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

        // D1 invariant: native tokens never mint/burn on any leg, so pass-through is
        // always safe. ERC20 wrapped tokens do mint/burn, so both the from-chain and
        // to-chain pairs must be origin pairs to avoid a mint -> burn round trip.
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
