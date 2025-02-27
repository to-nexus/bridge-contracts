// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {IChainManager} from "../interface/IChainManager.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

library ChainLib {
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Math for uint;

    error errChainLibNotExistingIndex(uint index);
    error errChainLibDuplicateIndex(uint index);
    error errChainLibInsufficientBalance(uint remoteChainID, address token, uint value);
    error errChainLibNotExistToken(address token);
    error errChainLibAleadyExistToken(address token);
    error errChainLibCanNotZeroAddress(string name);
    error errChainLibNotPaused(address token);
    error errChainLibAleadyPaused(address token);

    function getNextInitiateIndex(IChainManager.Chain storage chain) internal view returns (uint) {
        return chain.index.initiate + 1;
    }

    function getNextFinalizeIndex(IChainManager.Chain storage chain) internal view returns (uint) {
        return chain.index.finalize + 1;
    }

    function incrementInitiateIndex(IChainManager.Chain storage chain) internal {
        unchecked {
            ++chain.index.initiate;
        }
    }

    function incrementFinalizeIndex(IChainManager.Chain storage chain) internal {
        unchecked {
            ++chain.index.finalize;
        }
    }

    function getRevertedArguments(IChainManager.Chain storage chain, uint index)
        internal
        view
        returns (IChainManager.FinalizeArguments memory)
    {
        return chain.reverted.data[index];
    }

    function getRevertedReason(IChainManager.Chain storage chain, uint index) internal view returns (string memory) {
        return chain.reverted.reason[index];
    }

    function getRevertedData(IChainManager.Chain storage chain, uint index)
        internal
        view
        returns (IChainManager.FinalizeArguments memory args)
    {
        require(chain.reverted.index.contains(index), errChainLibNotExistingIndex(index));
        return getRevertedArguments(chain, index);
    }

    function setRevertedData(
        IChainManager.Chain storage chain,
        IChainManager.FinalizeArguments memory args,
        string memory reason
    ) internal {
        uint index = args.index;
        require(chain.reverted.index.add(index), errChainLibDuplicateIndex(index));
        chain.reverted.data[index] = args;
        chain.reverted.reason[index] = reason;
    }

    function removeRevertedData(IChainManager.Chain storage chain, uint index) internal {
        require(chain.reverted.index.remove(index), errChainLibNotExistingIndex(index));
        delete (chain.reverted.reason[index]);
        delete (chain.reverted.data[index]);
    }

    function registerToken(IChainManager.Chain storage chain, bool isOrigin, address localToken, address remoteToken)
        internal
    {
        require(localToken != address(0), errChainLibCanNotZeroAddress("localToken"));
        require(chain.tokens.add(localToken), errChainLibAleadyExistToken(localToken));
        chain.tokenPairs[localToken] = IChainManager.TokenPair({
            localToken: localToken,
            remoteToken: remoteToken,
            isOrigin: isOrigin,
            paused: false,
            deposited: 0
        });
    }

    function unregisterToken(IChainManager.Chain storage chain, address localToken) internal {
        require(localToken != address(0), errChainLibCanNotZeroAddress("localToken"));
        require(chain.tokens.remove(localToken), errChainLibNotExistToken(localToken));
        delete (chain.tokenPairs[localToken]);
    }

    function pauseToken(IChainManager.Chain storage chain, address token) internal {
        require(chain.tokens.contains(token), errChainLibNotExistToken(token));
        require(!chain.tokenPairs[token].paused, errChainLibAleadyPaused(token));
        chain.tokenPairs[token].paused = true;
    }

    function unpauseToken(IChainManager.Chain storage chain, address token) internal {
        require(chain.tokens.contains(token), errChainLibNotExistToken(token));
        require(chain.tokenPairs[token].paused, errChainLibNotPaused(token));
        chain.tokenPairs[token].paused = false;
    }

    function depositToken(IChainManager.Chain storage chain, address token, uint value) internal {
        chain.tokenPairs[token].deposited += value;
    }

    function withdrawToken(IChainManager.Chain storage chain, address token, uint value) internal {
        (bool ok, uint deposited) = chain.tokenPairs[token].deposited.trySub(value);
        require(ok, errChainLibInsufficientBalance(chain.remoteChainID, token, value));
        chain.tokenPairs[token].deposited = deposited;
    }
}
