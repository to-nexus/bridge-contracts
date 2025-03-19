// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {Test} from "forge-std/Test.sol";

import {ValidatorManager} from "../src/abstract/ValidatorManager.sol";
import {CrossCheckBlock, CrossCheckStorage} from "../src/abstract/CrossCheckStorage.sol";
import {ICrossCheck} from "../src/interface/ICrossCheck.sol";
import {CrossCheck} from "../src/CrossCheck.sol";
import {Const} from "../src/lib/Const.sol";

// forge coverage --match-contract CrossCheck --report lcov

contract CrossCheckTest is Test {
    uint256 constant N_VALIDATORS = 3;
    uint256 constant N_BLOCKS = 1000;
    uint256 constant CHAIN_ID = 612044;

    bytes32 constant SUBMIT_TYPEHASH =
        keccak256("CheckBlockArg(uint256 nonce,uint256 start,uint256 end,bytes32 rootHash,uint256 chainID)");

    address depl;
    address[] vals;
    uint256[] valKeys;

    CrossCheck crossCheck;

    function setUp() public {
        // setup accounts
        depl = vm.randomAddress();

        vals = new address[](N_VALIDATORS);
        valKeys = new uint256[](N_VALIDATORS);
        for (uint256 i = 0; i < N_VALIDATORS; i++) {
            uint salt = vm.randomUint() % 1000;
            (vals[i], valKeys[i]) = makeAddrAndKey(string.concat("validator-", vm.toString(salt)));
        }

        sortValidators();

        // deploy
        vm.startPrank(depl);

        CrossCheck impl = new CrossCheck();
        ERC1967Proxy proxy = new ERC1967Proxy(address(impl), "");
        crossCheck = CrossCheck(address(proxy));
        crossCheck.initialize(CHAIN_ID, N_BLOCKS, uint8(N_VALIDATORS));

        // setup validators
        crossCheck.grantRoleBatch(Const.VALIDATOR_ROLE, vals);
        crossCheck.grantRoleBatch(crossCheck.VALIDATOR_ROLE(), vals);
    }

    function test_initialize_Initialized() public view {
        assertEq(crossCheck.latestBlock(), 0);

        (uint256 nextNonce, uint256 nextStart) = crossCheck.getNextBlockInfo();
        assertEq(nextNonce, 0);
        assertEq(nextStart, 0);

        (, , , uint256 createdAt, , ) = crossCheck.getCheckBlock(0);
        assertEq(createdAt, 0);
    }

    function testFuzz_submitCheckpoint(uint256 timemod) public {
        uint256 timestamp = 1641080000 + (timemod % 100000);
        vm.warp(timestamp);

        // proposer will submit a checkpoint
        address proposer = getProposer();
        vm.startPrank(proposer);

        ICrossCheck.CheckBlockArg memory _block = createBlockArg(0, 0);

        // NewCheckBlock should be emitted
        vm.expectEmit(true, true, true, true);
        emit CrossCheckBlock.NewCheckBlock(proposer, _block.nonce, _block.start, _block.end, _block.rootHash);

        (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);

        assertEq(crossCheck.latestBlock(), 0);

        (uint256 nextNonce, uint256 nextStart) = crossCheck.getNextBlockInfo();
        assertEq(nextNonce, 1);
        assertEq(nextStart, N_BLOCKS);

        verifyBlock(_block, proposer, timestamp);
    }

    function test_submitCheckpoint_RevertIf_BlockIsNotNext() public {
        address proposer = getProposer();

        // wrong nonce
        ICrossCheck.CheckBlockArg memory _block = createBlockArg(0, 0);
        _block.nonce = 1;
        (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);

        vm.expectPartialRevert(CrossCheck.CrossCheckNotNextBlock.selector);
        vm.startPrank(proposer);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);

        // wrong end
        _block.nonce = 0;
        _block.end = _block.start + N_BLOCKS + 1;
        (vs, rs, ss) = signBlock(_block);

        vm.expectPartialRevert(CrossCheck.CrossCheckNotNextBlock.selector);
        vm.startPrank(proposer);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);
    }

    function test_submitCheckpoint_RevertIf_InvalidData() public {
        address proposer = getProposer();

        ICrossCheck.CheckBlockArg memory _block = createBlockArg(0, 0);
        _block.rootHash = bytes32(0); // invalid root hash
        (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);

        vm.expectPartialRevert(CrossCheckStorage.CrossCheckInvalidData.selector);
        vm.startPrank(proposer);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);
    }

    function test_submitCheckpoint_RevertIf_ChainIDMismatch() public {
        address proposer = getProposer();
        ICrossCheck.CheckBlockArg memory _block = createBlockArg(0, 0);
        _block.chainID = CHAIN_ID + 1; // wrong chain ID
        (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);

        vm.expectPartialRevert(CrossCheck.CrossCheckInvalidChainID.selector);
        vm.startPrank(proposer);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);
    }

    function test_submitCheckpoint_RevertIf_Paused() public {
        address proposer = getProposer();
        ICrossCheck.CheckBlockArg memory _block = createBlockArg(0, 0);
        (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);

        vm.startPrank(depl);
        crossCheck.grantRole(crossCheck.OPERATOR_ROLE(), depl);

        // pause
        vm.expectEmit(false, false, false, true);
        emit PausableUpgradeable.Paused(depl);
        crossCheck.pause();

        // should revert
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);

        vm.startPrank(proposer);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);

        // unpause
        vm.startPrank(depl);
        vm.expectEmit(false, false, false, true);
        emit PausableUpgradeable.Unpaused(depl);
        crossCheck.unpause();

        // should succeed
        vm.startPrank(proposer);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);
    }

    // utility functions

    function sortValidators() internal {
        // insertion sort in ascending order
        for (uint256 i = 0; i < vals.length; ++i) {
            address current = vals[i];
            uint256 currentKey = valKeys[i];
            uint256 j = i;
            while (j > 0 && vals[j - 1] > current) {
                vals[j] = vals[j - 1];
                valKeys[j] = valKeys[j - 1];
                --j;
            }
            vals[j] = current;
            valKeys[j] = currentKey;
        }
    }

    function getProposer() internal returns (address) {
        return vals[vm.randomUint() % N_VALIDATORS];
    }

    function createBlockArg(
        uint256 nonce,
        uint256 start
    ) internal view returns (ICrossCheck.CheckBlockArg memory _block) {
        _block = ICrossCheck.CheckBlockArg({
            nonce: nonce,
            start: start,
            end: start + N_BLOCKS - 1,
            rootHash: keccak256(abi.encodePacked(vm.randomBytes(32))),
            chainID: CHAIN_ID
        });
    }

    function signBlock(
        ICrossCheck.CheckBlockArg memory _block
    ) internal view returns (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) {
        (vs, rs, ss) = signBlock(_block, valKeys);
    }

    function signBlock(
        ICrossCheck.CheckBlockArg memory _block,
        uint256[] memory keys
    ) internal view returns (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) {
        bytes32 dataHash = MessageHashUtils.toTypedDataHash(
            crossCheck.domainSeparator(),
            keccak256(
                abi.encode(SUBMIT_TYPEHASH, _block.nonce, _block.start, _block.end, _block.rootHash, _block.chainID)
            )
        );

        (vs, rs, ss) = (new uint8[](keys.length), new bytes32[](keys.length), new bytes32[](keys.length));
        for (uint256 i = 0; i < keys.length; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(keys[i], dataHash);
            vs[i] = v;
            rs[i] = r;
            ss[i] = s;
        }
    }

    function verifyBlock(ICrossCheck.CheckBlockArg memory _block, address _proposer, uint256 timestamp) internal view {
        (uint256 nonce, uint256 start, uint256 end, uint256 createdAt, bytes32 rootHash, address proposer) = crossCheck
            .getCheckBlock(_block.start);
        assertEq(nonce, _block.nonce);
        assertEq(start, _block.start);
        assertEq(end, _block.end);
        assertEq(createdAt, timestamp);
        assertEq(rootHash, _block.rootHash);
        assertEq(proposer, _proposer);
    }
}
