// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// import {ERC20, IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20.sol";

import {IBridgeStandard} from "./interface/IBridgeStandard.sol";
import {ICrossMintableERC20, ICrossMintableERC20Code} from "./interface/ICrossMintableERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC20, ERC20Permit, IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import {IERC20, IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

contract CrossMintableERC20Code is ICrossMintableERC20Code {
    function code() external pure returns (bytes memory) {
        return type(CrossMintableERC20).creationCode;
    }
}

contract CrossMintableERC20 is Ownable, Pausable, ERC20, ERC20Permit, ICrossMintableERC20 {
    error CrossMintableERC20NotMinter(address account);

    uint8 private immutable _decimals;
    address public bridge;

    constructor(string memory name_, string memory symbol_, uint8 decimals_)
        Ownable(_msgSender())
        ERC20(name_, symbol_)
        ERC20Permit(name_)
    {
        bridge = _msgSender();
        transferOwnership(Ownable(bridge).owner());

        _decimals = decimals_;
    }

    modifier onlyBridge() {
        require(isBridge(_msgSender()), CrossMintableERC20NotMinter(_msgSender()));
        _;
    }

    function isBridge(address _account) public view returns (bool) {
        return address(bridge) == _account;
    }

    function mint(address _account, uint _amount) external onlyBridge returns (bool) {
        _mint(_account, _amount);
        return true;
    }

    function burn(address _account, uint _amount) external onlyBridge returns (bool) {
        _burn(_account, _amount);
        return true;
    }

    function transferFrom(address _sender, address _recipient, uint _amount)
        public
        override(ERC20, IERC20)
        returns (bool)
    {
        if (allowance(_sender, _msgSender()) != type(uint).max) super.transferFrom(_sender, _recipient, _amount);
        else _transfer(_sender, _recipient, _amount);
        return true;
    }

    function decimals() public view override(ERC20, IERC20Metadata) returns (uint8) {
        return _decimals;
    }

    function nonces(address owner) public view override(ERC20Permit, IERC20Permit) returns (uint) {
        return super.nonces(owner);
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function _update(address from, address to, uint value) internal override whenNotPaused {
        super._update(from, to, value);
    }
}
