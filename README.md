# STOKES

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/Sam-Stokes/stokes)

**STOKES** is a proof-of-work cryptocurrency with instant confirmations and sub-second block times, featuring **Bitcoin-style halving** for predictable, deflationary emission.

Built on Kaspa's PHANTOM protocol (a generalization of Nakamoto consensus), STOKES combines the speed and scalability of a blockDAG with the proven economic model of Bitcoin.

## 🎯 What is STOKES?

STOKES is a fair-launch cryptocurrency with:
- **Bitcoin-style halving:** 50 STKS → 25 → 12.5 → ... every ~4 years
- **Fixed supply cap:** 12.6 billion STKS maximum
- **Instant confirmations:** Sub-second block times via blockDAG
- **High throughput:** Scalable architecture
- **Fair distribution:** No premine, no ICO, pure proof-of-work

### Key Differences from Kaspa
- **Emission:** Bitcoin-style halving vs Kaspa's smooth emission
- **Supply:** 12.6B fixed cap vs Kaspa's 28.7B
- **Network:** Completely independent genesis and network
- **Philosophy:** Deflationary scarcity model

## 📊 Emission Schedule

| Block Range | Reward | Duration | Total STKS |
|-------------|--------|----------|------------|
| 0 - 126.23M | 50 STKS | ~4 years | 6.31B |
| 126.23M - 252.46M | 25 STKS | ~4 years | 3.16B |
| 252.46M - 378.69M | 12.5 STKS | ~4 years | 1.58B |
| ... | ... | ... | ... |
| **Total** | | | **12.6B** |

## 🚀 Status

**Phase 1: COMPLETE ✅**
- ✅ Bitcoin-style halving implemented
- ✅ Network isolation complete
- ✅ New genesis blocks generated
- ✅ Full rebranding complete
- ✅ Compilation successful
- ✅ Mining functionality verified

**Phase 2: In Planning**
- See [PHASE2_PLAN.md](PHASE2_PLAN.md) for details

## 📚 Documentation

- **[Phase 1 Complete](PHASE1_COMPLETE.md)** - What we've accomplished
- **[Phase 2 Plan](PHASE2_PLAN.md)** - Testnet launch roadmap
- **[Testing Guide](TESTING_GUIDE.md)** - How to test STOKES

## 🌐 Genesis Blocks

STOKES has unique genesis blocks for all networks:

- **Mainnet:** `dca25f14ec25b37efcb1ba767154ef9f2473e596a518d1f1c0be19b87d786949`
- **Testnet:** `4240b4dbce9f3a5d1483fd19146051264571494700ed3818b46cd84b0817ef38`
- **Simnet:** `3d7f1715e6f7c2744730462226a37c196d879f7391cdbcf8d28efe68e2655c779`
- **Devnet:** `90d169c8336ca62088271141a5d2c610fd20ac913f678e1711542bfd9a732058`

**Genesis Message:** "STOKES - Fair Launch 2025 - Bitcoin-style Halving"

## Requirements

Go 1.23 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
go version
```

- Clone and build STOKES:

```bash
git clone https://github.com/Sam-Stokes/stokes
cd stokes
go build -o stokesd .
go build -o stokesctl ./cmd/stokesctl
go build -o stokesminer ./cmd/stokesminer
go build -o stokeswallet ./cmd/stokeswallet
```

- Binaries will be created in the current directory

## Quick Start

### Running a Node (Simnet for testing)

```bash
./stokesd --simnet --appdir=./data --utxoindex
```

### Creating a Wallet

```bash
./stokeswallet create --simnet
./stokeswallet new-address --simnet
```

### Mining

```bash
./stokesminer --simnet --miningaddr=<YOUR_ADDRESS> --rpcserver=localhost:17510 --mine-when-not-synced
```

For detailed testing instructions, see [TESTING_GUIDE.md](TESTING_GUIDE.md)

## Network Ports

| Network | P2P Port | RPC Port |
|---------|----------|----------|
| Mainnet | 17611 | 17610 |
| Testnet | 17711 | 17710 |
| Simnet | 17511 | 17510 |
| Devnet | 17611 | 17610 |

## License

STOKES is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).

## Acknowledgments

STOKES is built on the foundation of [Kaspa](https://github.com/kaspanet/kaspad), implementing the PHANTOM protocol for blockDAG consensus.

## Contributing

Contributions are welcome! Please see our documentation for guidelines.

## Support

- **Issues:** [GitHub Issues](https://github.com/Sam-Stokes/stokes/issues)
- **Documentation:** See docs in this repository
- **Community:** Coming in Phase 2
