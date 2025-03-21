// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {Test} from "forge-std/Test.sol";

import {ValidatorManager} from "../src/abstract/ValidatorManager.sol";
import {CrossCheckBlock, CrossCheckStorage} from "../src/abstract/CrossCheckStorage.sol";
import {ICrossCheck} from "../src/interface/ICrossCheck.sol";
import {CrossCheck} from "../src/CrossCheck.sol";
import {CrossCheckMixin} from "../src/lib/CrossCheckMixin.sol";
import {Const} from "../src/lib/Const.sol";

// forge coverage --match-contract CrossCheck --report lcov

contract CrossCheckTest is Test {
    uint256 constant N_VALIDATORS = 3;
    uint256 constant N_BLOCKS = CrossCheckMixin.blocksPerCheck;
    uint256 constant CHAIN_ID = CrossCheckMixin.crossChainID;

    bytes32 constant SUBMIT_TYPEHASH = keccak256("CheckBlockArg(uint256 nonce,uint256 start,uint256 end,bytes32 rootHash,uint256 chainID)");

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
        crossCheck.initialize(uint8(N_VALIDATORS));

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

        verifyCheckBlock(_block, proposer, timestamp);
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

    function test_verifyBlock() public {
        // use the pre-computed values, generated by bridge-validator from a random merkle tree
        // root
        bytes32 rootHash = 0xcf2d1648acc4a788513d22ce2b63d58535af57dca09772c0beebfb9a7654b779;
        // leaf
        bytes32 blockHash = 0x4a03938a8241a2d3e3167e7bf1c4e6527b6c44c82142a99e46c5a26fc879db35;
        // path to root from leaf
        bytes32[] memory proof = new bytes32[](17);
        proof[0] = 0xec1c8ac11be24af989145f162bb483e0fd877ef85520e05289894569b75c7d7a;
        proof[1] = 0xc29fd136b1fdd8684b6f5d96fbfa615b5f9385b4249968b4640370a7b17ac9ea;
        proof[2] = 0x0ff54c19f610939d215b10c2f419a9526ec2e57028135bf2a8fcdacf8618f4db;
        proof[3] = 0x6aac385fe32734b9fbe7695f9d9c7379232116ade861643a042e17076e9c6e45;
        proof[4] = 0xa1a21d591c54c7853a74a4855ce85c34541677f86f8fc095d9105429343fa91e;
        proof[5] = 0xf8339e08566995e240044970c7836ac627ce259c54951ee469abc7d414e73b2b;
        proof[6] = 0x63e1b24d32d1cfdea2df710111e3338c3e3a1a1226537bf4bc448b08b99fede8;
        proof[7] = 0xb2da1354d32d3c4fe2059504f08190705a06d8cc4ec8a3ebde15075380b8e9d0;
        proof[8] = 0x3659b14fc60b9091c86f631cfe81f72d72150829937f3358a2d7886d1858fc0d;
        proof[9] = 0xe19ae0e94c9cc9b3175ec036eb70ff0ee6164cfbbac1ad9340b21fd48fb1b378;
        proof[10] = 0x0fe7113ee0e43f527ef63a6ffdfc63d1d4bc403ed14f35cc45ceb85b4abd4c78;
        proof[11] = 0xa565f09f59abd90a89c53bc5bfab972866f9544e70037acb5521701b7986145d;
        proof[12] = 0x3775bc9c496be3e634db63f31ca729381ecc9a402780308765d86a31f430cc20;
        proof[13] = 0xe3d11cf0dcd226d6dc15fe8a4ce10296a3121e6e01fa266911f980a625d3b9af;
        proof[14] = 0x9e62112efd3685401f32d71b5b2beb25d0ac98da64c1feb0a6500895db8fe141;
        proof[15] = 0x7ad88818d9a0b939699d0cece39d997f4b5b286579ec09d4bd4849d42d62e195;
        proof[16] = 0x443e2a4db47de56d57be355d7b8290dd6f5823da61a68da8e772a7e3c73e89f3;

        address proposer = getProposer();
        vm.startPrank(proposer);

        // submit a check block with the root hash
        ICrossCheck.CheckBlockArg memory _block = createBlockArg(0, 0, rootHash);
        (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);
        crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);

        assertEq(crossCheck.latestBlock(), 0);

        // the leaf should be proved
        bool result = crossCheck.verifyBlock(0, proof, blockHash);
        assertTrue(result);

        // wrong block hash
        result = crossCheck.verifyBlock(0, proof, bytes32(vm.randomBytes(32)));
        assertFalse(result);
    }

    function test_verifyBlock_RevertIf_BlockNotFound() public {
        bytes32 blockHash = bytes32(vm.randomBytes(32));
        bytes32[] memory proof = new bytes32[](1);
        proof[0] = bytes32(vm.randomBytes(32));

        vm.expectRevert(abi.encodeWithSelector(CrossCheck.CrossCheckBlockNotFound.selector, 0));

        // should revert because the check block does not exist
        crossCheck.verifyBlock(0, proof, blockHash);
    }

    function test_getCheckBlock() public {
        uint256 blockNumber = vm.randomUint() % (N_BLOCKS * 3);
        if (blockNumber % N_BLOCKS == 0) {
            blockNumber += 1;
        }

        address proposer = getProposer();
        vm.startPrank(proposer);

        // submit 3 check blocks
        for (uint256 i = 0; i < 3; ++i) {
            uint256 timestamp = 1641080000 + i * 60 * N_BLOCKS;
            vm.warp(timestamp);

            ICrossCheck.CheckBlockArg memory _block = createBlockArg(i, i * N_BLOCKS);
            (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) = signBlock(_block);
            crossCheck.submitCheckpoint(abi.encode(_block), vs, rs, ss);
        }

        // getCheckBlockByNonce
        for (uint256 i = 0; i < 3; ++i) {
            ICrossCheck.CheckBlockArg memory _block;
            uint256 timestamp;
            (_block.nonce, _block.start, _block.end, timestamp, _block.rootHash, ) = crossCheck.getCheckBlockByNonce(i);
            verifyCheckBlock(_block, proposer, timestamp);
        }

        // getCheckBlockByBlockNumber
        {
            ICrossCheck.CheckBlockArg memory _block;
            uint256 timestamp;
            (_block.nonce, _block.start, _block.end, timestamp, _block.rootHash, ) = crossCheck.getCheckBlockByBlockNumber(blockNumber);
            verifyCheckBlock(_block, proposer, timestamp);
        }

        // edge cases
        {
            // nonce = 0
            ICrossCheck.CheckBlockArg memory _block;
            uint256 timestamp;
            (_block.nonce, _block.start, _block.end, timestamp, _block.rootHash, ) = crossCheck.getCheckBlockByNonce(0);
            verifyCheckBlock(_block, proposer, timestamp);

            // block number = 0
            (_block.nonce, _block.start, _block.end, timestamp, _block.rootHash, ) = crossCheck.getCheckBlockByBlockNumber(0);
            verifyCheckBlock(_block, proposer, timestamp);

            // block number = _block.start
            (_block.nonce, _block.start, _block.end, timestamp, _block.rootHash, ) = crossCheck.getCheckBlockByBlockNumber(N_BLOCKS * 2);
            verifyCheckBlock(_block, proposer, timestamp);

            // block number = _block.end
            (_block.nonce, _block.start, _block.end, timestamp, _block.rootHash, ) = crossCheck.getCheckBlockByBlockNumber(N_BLOCKS * 2 - 1);
            verifyCheckBlock(_block, proposer, timestamp);
        }

        // not found
        {
            uint256 nonce;
            uint256 start;
            uint256 end;
            uint256 timestamp;
            bytes32 rootHash;
            address _proposer;

            (nonce, start, end, timestamp, rootHash, _proposer) = crossCheck.getCheckBlockByNonce(N_BLOCKS + 5);
            assertEq(nonce, 0);
            assertEq(start, 0);
            assertEq(end, 0);
            assertEq(timestamp, 0);
            assertEq(rootHash, bytes32(0));
            assertEq(_proposer, address(0));

            (nonce, start, end, timestamp, rootHash, _proposer) = crossCheck.getCheckBlockByBlockNumber(N_BLOCKS * 5);
            assertEq(nonce, 0);
            assertEq(start, 0);
            assertEq(end, 0);
            assertEq(timestamp, 0);
            assertEq(rootHash, bytes32(0));
            assertEq(_proposer, address(0));
        }
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

    function createBlockArg(uint256 nonce, uint256 start) internal view returns (ICrossCheck.CheckBlockArg memory _block) {
        return createBlockArg(nonce, start, bytes32(vm.randomBytes(32)));
    }

    function createBlockArg(uint256 nonce, uint256 start, bytes32 rootHash) internal pure returns (ICrossCheck.CheckBlockArg memory _block) {
        _block = ICrossCheck.CheckBlockArg({nonce: nonce, start: start, end: start + N_BLOCKS - 1, rootHash: rootHash, chainID: CHAIN_ID});
    }

    function signBlock(ICrossCheck.CheckBlockArg memory _block) internal view returns (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) {
        (vs, rs, ss) = signBlock(_block, valKeys);
    }

    function signBlock(
        ICrossCheck.CheckBlockArg memory _block,
        uint256[] memory keys
    ) internal view returns (uint8[] memory vs, bytes32[] memory rs, bytes32[] memory ss) {
        bytes32 dataHash = MessageHashUtils.toTypedDataHash(
            crossCheck.domainSeparator(),
            keccak256(abi.encode(SUBMIT_TYPEHASH, _block.nonce, _block.start, _block.end, _block.rootHash, _block.chainID))
        );

        (vs, rs, ss) = (new uint8[](keys.length), new bytes32[](keys.length), new bytes32[](keys.length));
        for (uint256 i = 0; i < keys.length; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(keys[i], dataHash);
            vs[i] = v;
            rs[i] = r;
            ss[i] = s;
        }
    }

    function verifyCheckBlock(ICrossCheck.CheckBlockArg memory _block, address _proposer, uint256 timestamp) internal view {
        (uint256 nonce, uint256 start, uint256 end, uint256 createdAt, bytes32 rootHash, address proposer) = crossCheck.getCheckBlock(_block.start);
        assertEq(nonce, _block.nonce);
        assertEq(start, _block.start);
        assertEq(end, _block.end);
        assertEq(createdAt, timestamp);
        assertEq(rootHash, _block.rootHash);
        assertEq(proposer, _proposer);
    }
}
