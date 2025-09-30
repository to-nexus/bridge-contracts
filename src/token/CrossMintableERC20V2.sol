// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";
import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {ERC20, ERC20Permit, IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract CrossMintableERC20V2 is ERC20, ERC20Permit, ICrossMintableERC20, AccessControlDefaultAdminRules {
    error CrossMintableERC20V2MinterNotGranted();

    uint8 private immutable _decimals;

    constructor(
        address initialOwner,
        address initialMinter,
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) ERC20(name_, symbol_) ERC20Permit(name_) AccessControlDefaultAdminRules(0, initialOwner) {
        if (initialMinter != address(0)) {
            // Grant the minter role to the bridge
            _grantRole(Const.MINTER_ROLE, initialMinter);
        }

        _decimals = decimals_;
    }

    function mint(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _mint(_account, _amount);
        return true;
    }

    function burn(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _burn(_account, _amount);
        return true;
    }

    function decimals() public view override(ERC20, ICrossMintableERC20) returns (uint8) {
        return _decimals;
    }

    function nonces(address owner_) public view override(ERC20Permit, ICrossMintableERC20) returns (uint) {
        return super.nonces(owner_);
    }
}
