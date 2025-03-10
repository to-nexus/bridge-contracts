// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// import {ERC20, IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20.sol";

import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";
import {ERC20, ERC20Permit, IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

contract CrossMintableERC20 is ERC20, ERC20Permit, ICrossMintableERC20 {
    error ERC20NotBridge(address account);
    error ERC20CanNotTransferToBridge();

    uint8 private immutable _decimals;
    address public immutable bridge;

    constructor(address _bridge, string memory name_, string memory symbol_, uint8 decimals_)
        ERC20(name_, symbol_)
        ERC20Permit(name_)
    {
        bridge = _bridge;
        _decimals = decimals_;
    }

    modifier onlyBridge() {
        require(bridge == msg.sender, ERC20NotBridge(msg.sender));
        _;
    }

    function mint(address _account, uint _amount) external onlyBridge returns (bool) {
        _mint(_account, _amount);
        return true;
    }

    function burn(address _account, uint _amount) external onlyBridge returns (bool) {
        _burn(_account, _amount);
        return true;
    }

    function decimals() public view override(ERC20, ICrossMintableERC20) returns (uint8) {
        return _decimals;
    }

    function nonces(address owner) public view override(ERC20Permit, ICrossMintableERC20) returns (uint) {
        return super.nonces(owner);
    }

    function _update(address from, address to, uint value) internal override {
        if (_msgSender() != address(bridge) && to == address(bridge)) revert ERC20CanNotTransferToBridge();
        super._update(from, to, value);
    }
}
