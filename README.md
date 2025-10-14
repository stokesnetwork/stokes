# 🚀 Stokes (STKS)

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/stokesnetwork/stokes)
[![Testnet](https://img.shields.io/badge/testnet-live-green.svg)](https://stokesnetwork.github.io/stokes)

**Stokes - also known as StokesCoin in its early release** is a proof-of-work, fair, decentralized cryptocurrency combining the speed of Kaspa's blockDAG with Bitcoin's proven halving economics.

> **"Kaspa’s speed with Bitcoin’s halving."**

Built on the PHANTOM protocol (a generalization of Nakamoto consensus), Stokes delivers instant confirmations with sub-second block times while maintaining a predictable, deflationary supply schedule.

---

## 🎯 What is Stokes?

Stokes is a **fair-launch cryptocurrency** designed for long-term value accrual through:

### Core Features

- ⚡ **Instant Confirmations** - Sub-second block times via blockDAG architecture
- 💎 **Bitcoin-Style Halving** - 50 STKS → 25 → 12.5... every ~4 years
- 🔒 **Fixed Supply Cap** - 12.6 billion STKS maximum (vs Bitcoin's 21M)
- 🎯 **Fair Launch** - No premine, no ICO, no VC allocation
- 🌐 **High Throughput** - Scalable parallel block processing
- 🔐 **Proven Security** - Battle-tested PHANTOM consensus

### Why Stokes?

**The Problem:** Most cryptocurrencies choose between speed OR scarcity:

- Bitcoin: Scarce but slow (10 min blocks)
- Kaspa: Fast but smooth emission (no halving events)

**The Solution:** Stokes combines both:

- ✅ Fast like Kaspa (1 block/second)
- ✅ Scarce like Bitcoin (halving every 4 years)
- ✅ Fair distribution (pure PoW mining)

### Key Differences from Kaspa

| Feature        | Stokes                | Kaspa                     |
| -------------- | --------------------- | ------------------------- |
| **Emission**   | Bitcoin-style halving | Smooth chromatic emission |
| **Supply**     | 12.6B fixed cap       | 28.7B fixed cap           |
| **Halvings**   | Every ~4 years        | No halvings               |
| **Philosophy** | Deflationary scarcity | Smooth distribution       |
| **Network**    | Independent genesis   | Original Kaspa network    |

## 📊 Emission Schedule

| Block Range       | Reward    | Duration | Total STKS |
| ----------------- | --------- | -------- | ---------- |
| 0 - 126.23M       | 50 STKS   | ~4 years | 6.31B      |
| 126.23M - 252.46M | 25 STKS   | ~4 years | 3.16B      |
| 252.46M - 378.69M | 12.5 STKS | ~4 years | 1.58B      |
| ...               | ...       | ...      | ...        |
| **Total**         |           |          | **12.6B**  |

## 🚀 Current Status

**✅ Testnet Live** - Public testing now available!

- ✅ Bitcoin-style halving implemented and tested
- ✅ Solo and multi-node mining working
- ✅ Wallet functionality complete
- ✅ Network isolation verified
- ✅ Genesis blocks generated for all networks
- ⏳ Public seed nodes deploying soon
- ⏳ Block explorer in development

## 📚 Documentation

- **[Installation Guide](#-installation)** - Get started quickly
- **[Mining Guide](#-mining)** - Start earning STKS
- **[Wallet Guide](#-wallet-usage)** - Manage your coins
- **[Troubleshooting](#-troubleshooting)** - Common issues

## 🌐 Genesis Blocks

Stokes has unique genesis blocks for all networks:

- **Mainnet:** `dca25f14ec25b37efcb1ba767154ef9f2473e596a518d1f1c0be19b87d786949`
- **Testnet:** `4240b4dbce9f3a5d1483fd19146051264571494700ed3818b46cd84b0817ef38`
- **Simnet:** `3d7f1715e6f7c2744730462226a37c196d879f7391cdbcf8d28efe68e2655c779`
- **Devnet:** `90d169c8336ca62088271141a5d2c610fd20ac913f678e1711542bfd9a732058`

**Genesis Message:** "Stokes - Fair Launch 2025 - Bitcoin-style Halving"

## 💻 Installation

### Option 1: Download Pre-Built Binaries (Recommended)

**Coming Soon:** Pre-built binaries for Linux, macOS, and Windows will be available on the [Releases](https://github.com/stokesnetwork/stokes/releases) page.

### Option 2: Build from Source

**Requirements:**

- Go 1.23 or later
- Git
- 2GB RAM minimum
- 10GB disk space

**Steps:**

1. **Install Go** (if not already installed):

   ```bash
   # Visit: https://go.dev/doc/install
   # Or use your package manager:

   # macOS
   brew install go

   # Ubuntu/Debian
   sudo apt install golang-go

   # Verify installation
   go version  # Should show 1.23 or later
   ```

2. **Clone the repository:**

   ```bash
   git clone https://github.com/stokesnetwork/stokes
   cd stokes
   ```

3. **Build all binaries:**

   ```bash
   # Build node
   go build -o stokesd .

   # Build CLI tool
   go build -o stokesctl ./cmd/stokesctl

   # Build miner
   go build -o stokesminer ./cmd/stokesminer

   # Build wallet
   go build -o stokeswallet ./cmd/stokeswallet
   ```

4. **Verify build:**
   ```bash
   ./stokesd --version
   ```

**Build time:** 2-5 minutes depending on your system

## 🚀 Quick Start (Testnet)

### Step 1: Run a Node

```bash
# Start testnet node
./stokesd --testnet \
  --utxoindex \
  --nodnsseed \
  --nolisten

# For public node (accepts connections):
./stokesd --testnet \
  --utxoindex \
  --listen=0.0.0.0:17711 \
  --rpclisten=0.0.0.0:17710
```

**What you'll see:**

```
[INF] KASD: Version 0.12.22
[INF] KASD: UTXO index started
[INF] TXMP: RPC Server listening on 127.0.0.1:17710
```

### Step 2: Create a Wallet

```bash
# Create new wallet
./stokeswallet --testnet create -f ~/stokes-wallet/keys.json

# You'll be prompted for a password
# SAVE YOUR SEED PHRASE - This is your backup!

# Start wallet daemon (in new terminal)
./stokeswallet --testnet start-daemon \
  -f ~/stokes-wallet/keys.json \
  -s 127.0.0.1:17710

# Get your mining address (in new terminal)
./stokeswallet --testnet new-address
```

**Example output:**

```
New address:
stokestest:qpkpllexmwjpdfru335psxssd3v8hs2l2gp8qv74gm8ajrrpxv2e7mmsr4ucq
```

### Step 3: Start Mining

```bash
# Replace YOUR_ADDRESS with address from Step 2
./stokesminer --testnet \
  --miningaddr=stokestest:YOUR_ADDRESS \
  --rpcserver=127.0.0.1:17710 \
  --mine-when-not-synced
```

**What you'll see:**

```
[INF] KSMN: Found block 761dc037... with parents [97b58165...]
[INF] KSMN: Submitting block 761dc037... to 127.0.0.1:17710
[INF] KSMN: Current hash rate is 160.02 Khash/s
```

### Step 4: Check Your Balance

```bash
# Check balance (after mining 100+ blocks)
./stokeswallet --testnet balance

# Check block count
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetBlockCount
```

**Congratulations! You're now mining Stokes!** 🎉

## 💰 Wallet Usage

### Creating a Wallet

```bash
# Create wallet with custom location
./stokeswallet --testnet create -f ~/my-wallet/keys.json

# You'll be prompted for:
# 1. Password (choose a strong one!)
# 2. Confirm password
#
# IMPORTANT: Write down your seed phrase shown on screen!
# This is the ONLY way to recover your wallet if you lose the password.
```

### Managing Addresses

```bash
# Generate new receiving address
./stokeswallet --testnet new-address

# Show all your addresses
./stokeswallet --testnet show-addresses

# Check balance
./stokeswallet --testnet balance
```

### Sending STKS

```bash
# Send coins to another address
./stokeswallet --testnet send \
  --send-amount 100 \
  --to-address stokestest:RECIPIENT_ADDRESS_HERE

# The wallet will:
# 1. Create the transaction
# 2. Sign it with your private key
# 3. Broadcast to the network
# 4. Show you the transaction ID
```

### Wallet Security

**⚠️ CRITICAL: Protect Your Wallet**

1. **Backup your seed phrase** - Write it down on paper, store securely
2. **Use a strong password** - At least 12 characters, mix of letters/numbers/symbols
3. **Never share your keys** - Not your keys, not your coins!
4. **Keep keys.json safe** - This file contains your encrypted private keys

## ⛏️ Mining

### Solo Mining (Testnet)

```bash
# Basic solo mining
./stokesminer --testnet \
  --miningaddr=stokestest:YOUR_ADDRESS \
  --rpcserver=127.0.0.1:17710 \
  --mine-when-not-synced
```

### Mining Performance

**Expected hash rates:**

- CPU (4 cores): ~100-200 Khash/s
- CPU (8 cores): ~200-400 Khash/s
- CPU (16 cores): ~400-800 Khash/s

**Mining rewards:**

- Current: **50 STKS** per block
- Coinbase maturity: **100 blocks** (coins spendable after 100 confirmations)
- Block time: ~1 second average

### Checking Mining Progress

```bash
# Check how many blocks you've mined
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetBlockCount

# Check your balance (includes pending rewards)
./stokeswallet --testnet balance

# Monitor miner output for:
# [INF] KSMN: Found block ...
# [INF] KSMN: Current hash rate is X Khash/s
```

## 🔧 Advanced Configuration

### Running a Public Node

To help the network, run a public node that accepts connections:

```bash
./stokesd --testnet \
  --utxoindex \
  --listen=0.0.0.0:17711 \
  --rpclisten=0.0.0.0:17710 \
  --externalip=YOUR_PUBLIC_IP
```

**Firewall configuration:**

```bash
# Allow P2P connections
sudo ufw allow 17711/tcp

# Allow RPC (only if you want remote access)
# WARNING: Secure your RPC with authentication!
# sudo ufw allow 17710/tcp
```

### Connecting to Seed Nodes

```bash
# Connect to a specific seed node
./stokesd --testnet \
  --utxoindex \
  --connect=SEED_NODE_IP:17711
```

### Custom Data Directory

```bash
# Use custom directory for blockchain data
./stokesd --testnet \
  --appdir=/path/to/custom/directory \
  --utxoindex
```

## 🐛 Troubleshooting

### Node Won't Start

**Problem:** Node fails to start or crashes immediately

**Solutions:**

```bash
# 1. Check if port is already in use
lsof -i :17710

# 2. Kill existing process
pkill stokesd

# 3. Reset database (WARNING: deletes blockchain data)
rm -rf ~/Library/Application\ Support/stokes-testnet  # macOS
rm -rf ~/.stokes-testnet  # Linux

# 4. Start fresh
./stokesd --testnet --utxoindex
```

### Wallet Can't Connect

**Problem:** Wallet daemon shows connection errors

**Solutions:**

```bash
# 1. Verify node is running
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetBlockCount

# 2. Check node started with --utxoindex flag
# Restart node with:
./stokesd --testnet --utxoindex

# 3. Restart wallet daemon
pkill stokeswallet
./stokeswallet --testnet start-daemon \
  -f ~/stokes-wallet/keys.json \
  -s 127.0.0.1:17710
```

### Miner Not Finding Blocks

**Problem:** Miner runs but no blocks found

**This is normal!** Mining is probabilistic. Factors:

- Your hash rate (higher = more blocks)
- Network difficulty (adjusts automatically)
- Other miners on network

**Tips:**

- Be patient - blocks will come
- Check hash rate is reasonable (~100+ Khash/s)
- Ensure node is synced
- Try mining for at least 10-15 minutes

### Balance Shows "Pending"

**Problem:** Balance shows but marked as "pending"

**This is normal!** Coinbase rewards require **100 block confirmations** before they're spendable.

**Check maturity:**

```bash
# Get current block count
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetBlockCount

# Your first rewards become spendable at block 100
# Second rewards at block 101, etc.
```

### "Block rejected - node is in IBD"

**Problem:** Miner shows blocks rejected with "node is in IBD"

**Solution:** This was fixed in recent updates. If you still see this:

```bash
# 1. Pull latest code
git pull origin master

# 2. Rebuild
go build -o stokesd .

# 3. Restart node
./stokesd --testnet --utxoindex --nodnsseed --nolisten
```

## 📊 Network Information

### Network Ports

| Network | P2P Port | RPC Port |
| ------- | -------- | -------- |
| Mainnet | 17611    | 17610    |
| Testnet | 17711    | 17710    |
| Simnet  | 17511    | 17510    |
| Devnet  | 17611    | 17610    |

### Address Prefixes

| Network | Prefix        | Example                    |
| ------- | ------------- | -------------------------- |
| Mainnet | `stokes:`     | `stokes:qpkpllexmw...`     |
| Testnet | `stokestest:` | `stokestest:qpkpllexmw...` |
| Simnet  | `stokessim:`  | `stokessim:qpkpllexmw...`  |
| Devnet  | `stokesdev:`  | `stokesdev:qpkpllexmw...`  |

### Useful Commands

```bash
# Get node info
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetInfo

# Get block DAG info
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetBlockDagInfo

# Get peer info
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetConnectedPeerInfo

# Get mempool info
./stokesctl --testnet --rpcserver=127.0.0.1:17710 GetMempoolEntries
```

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **Test the network** - Run nodes, mine, report bugs
2. **Improve documentation** - Fix typos, add examples
3. **Submit bug reports** - Use GitHub Issues
4. **Propose features** - Open discussions
5. **Code contributions** - Submit pull requests

### Development Setup

```bash
# Clone repo
git clone https://github.com/stokesnetwork/stokes
cd stokes

# Run tests
go test ./...

# Run with race detector
go test -race ./...

# Format code
go fmt ./...
```

## 📞 Support & Community

- **GitHub Issues:** [Report bugs](https://github.com/stokesnetwork/stokes/issues)
- **Documentation:** This README and [docs/](docs/)
- **Website:** [https://stokesnetwork.github.io/stokes](https://stokesnetwork.github.io/stokes)

**Coming Soon:**

- Discord server
- Telegram group
- Block explorer
- Mining pools

## License

Stokes is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).

## Acknowledgments

Stokes is built on the foundation of [Kaspa](https://github.com/kaspanet/kaspad), implementing the PHANTOM protocol for blockDAG consensus.

## Contributing

Contributions are welcome! Please see our documentation for guidelines.

## Support

- **Issues:** [GitHub Issues](https://github.com/stokesnetwork/stokes/issues)
- **Documentation:** See docs in this repository
- **Community:** Coming in Phase 2
