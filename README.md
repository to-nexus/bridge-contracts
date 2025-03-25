## Bridge Contracts

This repository manages all contracts which works with the bridge backends.

## Setup

Dependencies must be updated.

```shell
$ forge install
```

## Usage

### Mixin
At first, `lib/CrossCheckMixin.sol` must be generated.  
This contains some constant variables and they are injected based on selected network.

Because upgradeable contracts cannot initialize immutables, this mixin declares constants instead of immutables.  
It is used for some gas savings.

```shell
# generate a mixin
$ ./scripts/gen-mixin [devnet|devnet3|testnet|mainnet]
# generate a mixin and run forge command
$ ./scripts/forge [devnet|devnet3|testnet|mainnet] <forge-command> ...
```

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test [-match-contract contract_name]
```

### Gas Snapshots

```shell
$ forge snapshot [--match-contract contract_name]
```

### Coverage Report

```shell
$ forge coverage [--match-contract contract_name] [--report lcov]
```