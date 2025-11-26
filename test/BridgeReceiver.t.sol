// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Test, console} from "forge-std/Test.sol";

import {IBridgeReceiver} from "../src/interface/IBridgeReceiver.sol";

/**
 * @title BridgeReceiverTest
 * @notice Tests for IBridgeReceiver callback functionality
 */
contract BridgeReceiverTest is Test {
    MockBridge public bridge;
    MockToken public token;

    SimpleReceiver public simpleReceiver;
    RevertingReceiver public revertingReceiver;
    GasHeavyReceiver public gasHeavyReceiver;

    address public user = address(0x1);
    uint public constant BRIDGE_AMOUNT = 1000 ether;

    function setUp() public {
        bridge = new MockBridge();
        token = new MockToken("Test Token", "TEST");

        simpleReceiver = new SimpleReceiver();
        revertingReceiver = new RevertingReceiver();
        gasHeavyReceiver = new GasHeavyReceiver();

        // Mint tokens to bridge
        token.mint(address(bridge), BRIDGE_AMOUNT * 10);
    }

    function testSimpleReceiverSuccess() public {
        bytes memory extraData = abi.encode("Hello from bridge!");

        // Simulate finalize with callback
        bridge.finalize(address(simpleReceiver), token, BRIDGE_AMOUNT, extraData);

        // Verify tokens transferred
        assertEq(token.balanceOf(address(simpleReceiver)), BRIDGE_AMOUNT);

        // Verify callback was called
        assertEq(simpleReceiver.callCount(), 1);
        assertEq(simpleReceiver.lastAmount(), BRIDGE_AMOUNT);
    }

    function testRevertingReceiverDoesNotRevertFinalize() public {
        bytes memory extraData = abi.encode("This will revert");

        // Even though callback reverts, finalize should succeed
        bridge.finalize(address(revertingReceiver), token, BRIDGE_AMOUNT, extraData);

        // Tokens should still be transferred
        assertEq(token.balanceOf(address(revertingReceiver)), BRIDGE_AMOUNT);
    }

    function testEOARecipientNoCallback() public {
        bytes memory extraData = abi.encode("EOA recipient");

        // EOA recipient should work without callback
        bridge.finalize(user, token, BRIDGE_AMOUNT, extraData);

        assertEq(token.balanceOf(user), BRIDGE_AMOUNT);
    }

    function testEmptyExtraDataNoCallback() public {
        bytes memory extraData = "";

        // No callback should be attempted with empty extraData
        bridge.finalize(address(simpleReceiver), token, BRIDGE_AMOUNT, extraData);

        assertEq(token.balanceOf(address(simpleReceiver)), BRIDGE_AMOUNT);
        assertEq(simpleReceiver.callCount(), 0); // Not called
    }

    function testGasHeavyReceiverWithLimit() public {
        bytes memory extraData = abi.encode(uint(100)); // 100 iterations

        // Should work within gas limit
        bridge.finalize(address(gasHeavyReceiver), token, BRIDGE_AMOUNT, extraData);

        assertEq(token.balanceOf(address(gasHeavyReceiver)), BRIDGE_AMOUNT);
    }

    function testInvalidSelectorReturned() public {
        InvalidSelectorReceiver invalidReceiver = new InvalidSelectorReceiver();
        bytes memory extraData = abi.encode("Invalid selector");

        // Should complete but emit failure event
        bridge.finalize(address(invalidReceiver), token, BRIDGE_AMOUNT, extraData);

        // Tokens still transferred
        assertEq(token.balanceOf(address(invalidReceiver)), BRIDGE_AMOUNT);
    }
}

/**
 * @notice Simple receiver that logs callback
 */
contract SimpleReceiver is IBridgeReceiver {
    uint public callCount;
    uint public lastAmount;
    bytes public lastExtraData;

    function onBridgeReceived(uint, uint, IERC20, uint amount, bytes calldata extraData)
        external
        override
        returns (bytes4)
    {
        callCount++;
        lastAmount = amount;
        lastExtraData = extraData;

        return IBridgeReceiver.onBridgeReceived.selector;
    }
}

/**
 * @notice Receiver that always reverts
 */
contract RevertingReceiver is IBridgeReceiver {
    function onBridgeReceived(uint, uint, IERC20, uint, bytes calldata) external pure override returns (bytes4) {
        revert("Callback revert");
    }
}

/**
 * @notice Receiver that consumes a lot of gas
 */
contract GasHeavyReceiver is IBridgeReceiver {
    uint public sum;

    function onBridgeReceived(uint, uint, IERC20, uint, bytes calldata extraData) external override returns (bytes4) {
        uint iterations = abi.decode(extraData, (uint));

        // Consume gas
        for (uint i = 0; i < iterations; i++) {
            sum += i;
        }

        return IBridgeReceiver.onBridgeReceived.selector;
    }
}

/**
 * @notice Receiver that returns wrong selector
 */
contract InvalidSelectorReceiver is IBridgeReceiver {
    function onBridgeReceived(uint, uint, IERC20, uint, bytes calldata) external pure override returns (bytes4) {
        return bytes4(keccak256("wrong()"));
    }
}

/**
 * @notice Mock bridge for testing
 */
contract MockBridge {
    event BridgeReceiverSuccess(address indexed recipient, uint fromChainID, uint index);
    event BridgeReceiverFailed(address indexed recipient, bytes reason);

    uint constant CALLBACK_GAS_LIMIT = 200_000;

    function finalize(address recipient, MockToken token, uint amount, bytes memory extraData) external {
        // 1. Transfer tokens (simulate finalize)
        token.transfer(recipient, amount);

        // 2. Callback if conditions met
        if (extraData.length > 0 && recipient.code.length > 0) {
            _callBridgeReceiver(1, 1, IERC20(address(token)), recipient, amount, extraData);
        }
    }

    function _callBridgeReceiver(
        uint fromChainID,
        uint index,
        IERC20 token,
        address recipient,
        uint amount,
        bytes memory extraData
    ) private {
        try IBridgeReceiver(recipient).onBridgeReceived{gas: CALLBACK_GAS_LIMIT}(
            fromChainID, index, token, amount, extraData
        ) returns (bytes4 selector) {
            if (selector == IBridgeReceiver.onBridgeReceived.selector) {
                emit BridgeReceiverSuccess(recipient, fromChainID, index);
            } else {
                emit BridgeReceiverFailed(recipient, "Invalid selector");
            }
        } catch (bytes memory reason) {
            emit BridgeReceiverFailed(recipient, reason);
        }
    }
}

/**
 * @notice Mock ERC20 token
 */
contract MockToken is IERC20 {
    string public name;
    string public symbol;
    uint8 public constant decimals = 18;
    uint public totalSupply;

    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

    constructor(string memory _name, string memory _symbol) {
        name = _name;
        symbol = _symbol;
    }

    function mint(address to, uint amount) external {
        balanceOf[to] += amount;
        totalSupply += amount;
    }

    function transfer(address to, uint amount) external returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        return true;
    }

    function transferFrom(address from, address to, uint amount) external returns (bool) {
        allowance[from][msg.sender] -= amount;
        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        return true;
    }

    function approve(address spender, uint amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        return true;
    }
}
