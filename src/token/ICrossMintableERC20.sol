// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

interface ICrossMintableERC20 is IERC20Metadata, IERC20Permit {
    function mint(address _account, uint _amount) external returns (bool);
    function burn(address _account, uint _amount) external returns (bool);
    function decimals() external view returns (uint8);
    function nonces(address owner) external view returns (uint);
}
