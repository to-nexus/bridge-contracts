// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {MockERC20} from "./mocks/MockERC20.sol";

contract TestToken is MockERC20 {
    error TestTokenRevertTransfer();

    bool public isStop = false;
    bool public revertTransfer = false;

    constructor(string memory name, string memory symbol, uint8 decimals) MockERC20() {
        initialize(name, symbol, decimals);
    }

    function stop() external {
        isStop = true;
    }

    function start() external {
        isStop = false;
    }

    function setRevertTransfer(bool revertTransfer_) external {
        revertTransfer = revertTransfer_;
    }

    function mint(address to, uint value) public {
        _mint(to, value);
    }

    function transfer(address to, uint amount) public override returns (bool) {
        require(!revertTransfer, TestTokenRevertTransfer());
        if (isStop) return false;
        return super.transfer(to, amount);
    }

    function transferFrom(address from, address to, uint amount) public override returns (bool) {
        require(!revertTransfer, TestTokenRevertTransfer());
        if (isStop) return false;
        return super.transferFrom(from, to, amount);
    }
}
