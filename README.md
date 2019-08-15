# 0x-go
[![Build Status](https://travis-ci.org/linki/0x-go.svg?branch=master)](https://travis-ci.org/linki/0x-go)
[![Coverage Status](https://coveralls.io/repos/github/linki/0x-go/badge.svg?branch=master)](https://coveralls.io/github/linki/0x-go?branch=master)
[![GitHub release](https://img.shields.io/github/release/linki/0x-go.svg)](https://github.com/linki/0x-go/releases)
[![go-doc](https://godoc.org/github.com/linki/0x-go/0x-go?status.svg)](https://godoc.org/github.com/linki/0x-go/0x-go)

A collection of tools relating to [Ethereum](https://www.ethereum.org/)'s [0xProject](https://0xproject.com/).

### Overview

It currently has a single command that calculates the hash of a given order. More commands will be added.

Calculate an order's hash.

```console
$ go run main.go orders hash \
  --exchange-contract-address=0x12459c951127e0c374ff9105dda097662a027093 \
  --maker=0xc9b32e9563fe99612ce3a2695ac2a6404c111dde \
  --taker=0x0000000000000000000000000000000000000000 \
  --maker-token-address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 \
  --taker-token-address=0xe41d2489571d322189246dafa5ebde1f4699f498 \
  --fee-recipient=0xa258b39954cef5cb142fd567a46cddb31a670124 \
  --maker-token-amount=18981000000000000 \
  --taker-token-amount=19000000000000000000 \
  --maker-fee=0 \
  --taker-fee=0 \
  --expiration-unix-timestamp-sec=1518201120 \
  --salt=58600101225676680041453168589125977076540694791976419610199695339725548478315

=> 0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942
```

Compare it with the order on RadarRelay. The order below was created through RadarRelay which also calculated the hash, presumably using 0x.js. Since both hashes are the same we can assume that the Go version is correct as well.

```console
curl -Ss https://api.radarrelay.com/0x/v0/order/0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942 | jq .
{
  "ecSignature": {
    "v": 28,
    "r": "0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9",
    "s": "0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709"
  },
  "exchangeContractAddress": "0x12459c951127e0c374ff9105dda097662a027093",
  "expirationUnixTimestampSec": "1518201120",
  "feeRecipient": "0xa258b39954cef5cb142fd567a46cddb31a670124",
  "maker": "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
  "makerFee": "0",
  "makerTokenAddress": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
  "makerTokenAmount": "18981000000000000",
  "salt": "58600101225676680041453168589125977076540694791976419610199695339725548478315",
  "taker": "0x0000000000000000000000000000000000000000",
  "takerFee": "0",
  "takerTokenAddress": "0xe41d2489571d322189246dafa5ebde1f4699f498",
  "takerTokenAmount": "19000000000000000000"
}
```

### Requirements

* go 1.11
* go-ethereum (geth) 1.9

### Warning

**This is a trivial implementation, the author is not a cryptographer, and the code has not been reviewed. Use
at your own risk.**
