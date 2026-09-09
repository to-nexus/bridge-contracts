// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// ===========================================================================
// FROZEN COPY — DO NOT EDIT. This is a byte-for-byte copy (only the contract
// name and import paths changed to fit this file's location under `test/legacy/`)
// of `src/CrossBridge.sol` as it stood immediately before `CrossBridge` and
// `CrossBridgeV2` were unified into a single implementation. `src/CrossBridge.sol`
// now holds what was `CrossBridgeV2`'s content — a sibling contract with an
// IDENTICAL storage layout (`crossSupplyLimit` · `_bscChainID` · `uint[48] __gap`),
// confirmed by on-chain measurement at the time of the unification.
//
// CROSS mainnet (612055) is, at the time of this rename, STILL RUNNING the exact
// code frozen here — implementation address `0x2827865ad5330faf855af8b743c701ae9c818034`
// (`bridgeTokenForwarded` reverts on it; it predates the multi-hop feature). This
// file is the ONLY evidence that upgrading that live proxy to the renamed
// `src/CrossBridge.sol` is layout-safe — see `test/CrossBridgeUpgrade.t.sol`,
// which upgrades an instance of THIS contract to `CrossBridge` and asserts every
// storage-backed field survives unchanged.
//
// Delete this file once CROSS mainnet has actually been upgraded off this code
// (a separately-approved, future operation) — at that point it stops being
// evidence of anything live and becomes pure dead weight.
// ===========================================================================

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "../../src/BaseBridge.sol";
import {Const} from "../../src/lib/Const.sol";
import {ICrossMintableERC20} from "../../src/token/ICrossMintableERC20.sol";

/**
 * @title CrossBridgeLegacy
 * @notice Cross-chain bridge with predeploy functionality
 * @dev Extends BaseBridge with deterministic deployment features
 * - Uses predefined implementation address for easier verification
 * - Implements proxy security checks
 */
contract CrossBridgeLegacy is BaseBridge {
    error CrossBridgeCanNotZeroAddress();
    error CrossBridgeCanNotZero();
    /**
     * @notice Emitted when the cross-chain supply limit is updated
     * @param crossSupplyLimit The new maximum supply limit for CROSS native token transfers
     */

    event CrossSupplyLimitSet(uint crossSupplyLimit);

    /// @dev Predefined address for the predeployed implementation
    address private constant PREDEPLOYED_IMPLEMENTATION_ADDRESS = address(0xB81D6e000000000000000000000000000000AAaA);

    /// @dev Maximum issuance limit for CROSS native token on the Cross chain
    /// @notice Defines the maximum amount of CROSS native tokens that can be unlocked from the bridge contract.
    /// @notice This limit is designed to be dynamically adjustable to mitigate security risks from potential
    /// @notice bridge exploits, as most of the token supply is initially locked in the bridge contract.
    uint public crossSupplyLimit;

    /// @dev BSC chain ID (e.g., 56 for mainnet, 97 for Sepolia testnet, or other chain IDs)
    /// @dev This is the chain where the Cross ERC20 Token is deployed
    uint private _bscChainID;

    /// @dev Storage gap for future upgrades
    uint[48] private __gap;

    /**
     * @notice Initializes the CrossBridge contract
     * @dev Sets up the contract with initial configuration
     * - Calls the base initialization in BaseBridge
     * - Records the initial balance for CROSS token supply tracking
     * - Sets the initial CROSS token supply limit to 0
     * - Pairs native CROSS token with the CROSS ERC20 token on BSC chain
     * @param owner Address that will receive admin role
     * @param dev_ Address of the developer account for receiving fees
     * @param threshold_ Minimum number of validators required for validation
     * @param bscChainID The chain ID of the BSC Network
     * @param cross Address of the CROSS ERC20 token on BSC Network
     * @param crossInitialSupply Pre-minted supply of CROSS tokens for the CROSS Foundation
     */
    function initializeCrossBridge(
        address owner,
        address payable dev_,
        uint8 threshold_,
        uint bscChainID,
        address cross,
        uint crossInitialSupply
    ) external initializer {
        require(bscChainID != 0, CrossBridgeCanNotZero());
        require(cross != address(0), CrossBridgeCanNotZeroAddress());

        __BaseBridge_init(owner, dev_, threshold_);

        // Register CROSS token as a token pair
        // This pairs the native CROSS token on this chain with the CROSS ERC20 token on BSC Network
        _registerToken(bscChainID, false, Const.NATIVE_TOKEN, cross);
        if (crossInitialSupply > 0) _withdrawToken(bscChainID, Const.NATIVE_TOKEN, crossInitialSupply);

        _bscChainID = bscChainID;
    }

    function crossSupply() public view returns (uint) {
        return _tokenPairs[_bscChainID][Const.NATIVE_TOKEN].minted;
    }

    /**
     * @notice Sets the maximum issuance limit for CROSS native token
     * @dev Only callable by admin role
     * - Updates the issuance limit for CROSS native token on the Cross chain
     * - Emits CrossSupplyLimitSet event upon successful update
     * @param _crossSupplyLimit New maximum issuance limit for CROSS native token
     */
    function setCrossSupplyLimit(uint _crossSupplyLimit) external onlyRole(Const.ADMIN_ROLE) {
        crossSupplyLimit = _crossSupplyLimit;
        emit CrossSupplyLimitSet(crossSupplyLimit);
    }

    /**
     * @notice Verifies if a finalization amount is within allowed limits
     * @dev Extends the base implementation with CROSS token issuance limit checks
     * - First performs standard amount verification via parent contract
     * - For CROSS native token, enforces the issuance limit
     * - Calculates current supply by comparing current balance to initial balance
     * @param fromChainID Source chain ID of the transfer
     * @param token Address of the token being transferred (Const.NATIVE_TOKEN for CROSS native token)
     * @param value Amount of tokens to finalize
     * @param retry Whether this is a retry of a previous finalization attempt
     * @return status Status code indicating the result of the check
     * @return delay Boolean indicating if finalization should be delayed
     */
    function _checkFinalizeAmount(uint fromChainID, IERC20 token, uint value, bool retry)
        internal
        override
        returns (Const.FinalizeStatus status, bool delay)
    {
        if (address(token) == Const.NATIVE_TOKEN && fromChainID == _bscChainID) {
            // Check if the new transfer would exceed the configured CROSS token issuance limit
            // If limit is exceeded, return a specific error status and mark for delay
            if (crossSupply() + value > crossSupplyLimit) return (Const.FinalizeStatus.CrossSupplyLimitExceeded, true);
        }

        return super._checkFinalizeAmount(fromChainID, token, value, retry);
    }

    /**
     * @notice Verifies proxy delegation (PREDEPLOY ONLY)
     * @dev This function is ONLY used for the initial predeploy deployment.
     *      MUST be removed or commented out before the first upgrade.
     *      Failure to remove will permanently lock the contract from future upgrades.
     *
     *      [WARNING] Remove this override or uncomment below before the 1st upgrade.
     */
    // function _checkProxy() internal view override {
    //     if (
    //         address(this) == PREDEPLOYED_IMPLEMENTATION_ADDRESS // Must be called through delegatecall
    //             || ERC1967Utils.getImplementation() != PREDEPLOYED_IMPLEMENTATION_ADDRESS // Must be called through an active proxy
    //     ) revert UUPSUnauthorizedCallContext();
    // }
}
