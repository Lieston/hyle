# Mini - A minimal Cosmos SDK chain

This repository contains an example of a tiny, but working Cosmos SDK chain.
It uses the least modules possible and is intended to be used as a starting point for building your own chain, without all the boilerplate that other tools generate. It is a simpler version of Cosmos SDK's [simapp](https://github.com/cosmos/cosmos-sdk/tree/main/simapp).

`Minid` uses the **latest** version of the [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk).

## How to use

In addition to learn how to build a chain thanks to `minid`, you can as well directly run `minid`.

### Installation

Install and run `minid`:

```sh
git clone git@github.com:hyle/hyle.git
cd chain-minimal
make install # install the minid binary
make init # initialize the chain
minid start # start the chain
```

## Useful links

* [Cosmos-SDK Documentation](https://docs.cosmos.network/)

## ZKTX module

Sending initial transactions
./minid tx zktx execute mini1s35tpv67eafejyvpxxdtn4e7dgm8whmm07a2x6 b AA== AQ==  
./minid tx zktx execute mini1s35tpv67eafejyvpxxdtn4e7dgm8whmm07a2x6 b AQ== Aq==  

Querying state
./minid query zktx contract mini1s35tpv67eafejyvpxxdtn4e7dgm8whmm07a2x6
