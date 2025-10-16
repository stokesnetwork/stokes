# 🚀 Stokes (STKS)

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/stokesnetwork/stokes)
[![Testnet](https://img.shields.io/badge/testnet-live-green.svg)](https://stokesnetwork.github.io/stokes)
[![Discord](https://img.shields.io/badge/discord-join-7289da.svg)](https://discord.gg/eMAcvMev)

> **🧪 Currently in Public Testnet - [Join Testing Now!](#-quick-start-testnet)**

**Stokes** is a proof-of-work, fair-launch cryptocurrency combining the speed of Kaspa's blockDAG with Bitcoin's proven halving economics.

> **"Kaspa's speed with Bitcoin's halving."**

Built on the PHANTOM protocol (a generalization of Nakamoto consensus), Stokes delivers instant confirmations with sub-second block times while maintaining a predictable, deflationary supply schedule.

**🎯 Fair Launch Commitment:** Mainnet genesis will be generated on launch day to ensure no pre-mining and equal opportunity for all participants.

---

## 🎯 What is Stokes?

Stokes is a **fair-launch cryptocurrency** designed for long-term value accrual through:

### Core Features

- ⚡ **Instant Confirmations** - Sub-second block times via blockDAG architecture
- 💎 **Bitcoin-Style Halving** - 0.0016667 STKS → 0.00083335 → ... every ~4 years
- 🔒 **Fixed Supply Cap** - 21 million STKS maximum (Bitcoin-equivalent)
- 🎯 **Fair Launch** - No premine, no ICO, no VC allocation
- 🌐 **High Throughput** - 50 blocks/sec parallel processing (500x Bitcoin)
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
| **Supply**     | 21M fixed cap         | 28.7B fixed cap           |
| **Halvings**   | Every ~4 years        | No halvings               |
| **Philosophy** | Bitcoin scarcity      | Smooth distribution       |
| **Network**    | Independent genesis   | Original Kaspa network    |

## 📊 Emission Schedule

**Bitcoin-Equivalent Emission:** 21M total supply, 4-year halvings

| Era | Years | Reward/Block | STKS This Era | Cumulative |
| --- | ----- | ------------ | ------------- | ---------- |
| 1   | 0-4   | 0.0016667    | 10.52M        | 10.52M (50%) |
| 2   | 4-8   | 0.00083335   | 5.26M         | 15.78M (75%) |
| 3   | 8-12  | 0.000416675  | 2.63M         | 18.41M (87.7%) |
| 4   | 12-16 | 0.0002083375 | 1.31M         | 19.72M (93.9%) |
| ... | ...   | ...          | ...           | ...        |
| **Total** |   |              |               | **21M STKS** |

**Daily Issuance:** ~7,200 STKS/day (first era) - identical to Bitcoin's 7,200 BTC/day

See [EMISSION_SCHEDULE.md](EMISSION_SCHEDULE.md) for complete details.

## 🚀 Current Status

**🧪 TESTNET PHASE - Help Us Test!**

We're in public testnet to ensure a fair, bug-free mainnet launch. Join us in testing:

**What's Working:**

- ✅ Bitcoin-equivalent emission (21M supply, 4-year halvings)
- ✅ Sub-second block times via blockDAG (50 blocks/sec)
- ✅ Solo and multi-node mining
- ✅ Wallet creation and transactions
- ✅ Full node synchronization

**What We're Testing:**

- 🧪 Network stability under load
- 🧪 Halving mechanism accuracy
- 🧪 Mining difficulty adjustment
- 🧪 Transaction validation
- 🧪 Wallet security and usability

**Coming Soon:**

- 📅 Mainnet launch date announcement
- 🌐 Public seed nodes
- 📊 Block explorer
- 💬 Community Discord/Telegram

## 📚 Documentation

- **[Installation Guide](#-installation)** - Get started quickly
- **[Troubleshooting](#-troubleshooting)** - Common issues

## ⚠️ TESTNET PHASE - Mainnet Not Yet Launched

**Current Status:** Public Testnet Testing

Stokes is currently in **testnet phase** to:

- 🧪 Test the Bitcoin-style halving implementation
- 🐛 Identify and fix bugs before mainnet
- 👥 Build community and gather feedback
- 📊 Validate network performance and security
- 🎯 Ensure a truly fair launch

**Mainnet Launch:** TBA - Will be announced with advance notice

### Why Testnet First?

We're committed to a **fair launch** where everyone starts together. By removing the mainnet genesis until launch day, we ensure:

- ✅ No pre-mining possible
- ✅ No insider advantage
- ✅ Transparent development process
- ✅ Community-tested software

## 🌐 Testnet Genesis

- **Testnet:** `4240b4dbce9f3a5d1483fd19146051264571494700ed3818b46cd84b0817ef38`
- **Devnet:** `90d169c8336ca62088271141a5d2c610fd20ac913f678e1711542bfd9a732058`
- **Simnet:** `3d7f1715e6f7c2744730462226a37c196d879f7391cdbcf8d28efe68e2655c779`

**Mainnet genesis will be generated on launch day to ensure fair distribution.**

## 💻 Installation

### Option 1: Download Pre-Built Binaries (Recommended)

Download the latest release for your platform from the [Releases](https://github.com/stokesnetwork/stokes/releases) page.

**Available platforms:**

- Linux (x64, ARM64)
- macOS (Intel, Apple Silicon)
- Windows (x64)

**Quick install:**

```bash
# Linux/macOS: Extract and move to PATH
tar -xzf stokes-v*-*.tar.gz
cd stokes-v*-*
sudo mv stokesd stokesctl stokesminer stokeswallet /usr/local/bin/

# Verify installation
stokesd --version
```

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

### Step 1: Run a Node (Terminal 1 - Keep Open)

**Option A: Connect to Public Testnet (Recommended)**

Try connecting to the public testnet seed nodes:

```bash
./stokesd --testnet \
  --utxoindex \
  --connect=95.216.155.253:17711 \
  --connect=46.62.218.114:17711
```

**Seed Nodes:**

- `95.216.155.253:17711` (Germany)
- `46.62.218.114:17711` (Germany)

**What you'll see if successful:**

```
[INF] STKS: Version 0.12.22
[INF] STKS: UTXO index started
[INF] TXMP: P2P Connected to 95.216.155.253:17711
[INF] TXMP: P2P Connected to 46.62.218.114:17711
[INF] TXMP: RPC Server listening on [::]:17210
```

**Option B: Standalone Mode (For Local Testing)**

If you want to test locally without connecting to the network, or if connection fails:

```bash
./stokesd --testnet \
  --utxoindex \
  --nodnsseed \
  --nolisten
```

**What you'll see:**

```
[INF] STKS: Version 0.12.22
[INF] STKS: UTXO index started
[INF] TXMP: RPC Server listening on [::]:17210
```

**Note:** Standalone mode creates your own local chain. Your blocks won't sync with the public testnet, but all features work identically for testing.

**⚠️ Keep this terminal open! The node must stay running.**

### Step 2: Create a Wallet (Terminal 2 - One Time Setup)

```bash
# Create new wallet
./stokeswallet --testnet create -f ~/stokes-wallet/keys.json

# You'll be prompted for a password
# SAVE YOUR SEED PHRASE - This is your backup!

# Start wallet daemon (keep this terminal open too)
./stokeswallet --testnet start-daemon \
  -f ~/stokes-wallet/keys.json \
  -s 127.0.0.1:17210
```

**⚠️ Keep Terminal 2 open! The wallet daemon must stay running.**

### Step 3: Get Mining Address (Terminal 3)

```bash
# Get a new address to receive mining rewards
./stokeswallet --testnet new-address
```

**Example output:**

```
New address:
stokestest:qpkpllexmwjp...
```

**Copy this address - you'll need it for mining!**

### Step 4: Start Mining (Terminal 3 - Keep Open)

**If connected to seed nodes (Option A):**

```bash
# Replace YOUR_ADDRESS with address from Step 3
./stokesminer --testnet \
  --miningaddr=stokestest:YOUR_ADDRESS
```

**If running standalone (Option B):**

```bash
# Replace YOUR_ADDRESS with address from Step 3
./stokesminer --testnet \
  --miningaddr=stokestest:YOUR_ADDRESS \
  --mine-when-not-synced
```

**Note:** The `--mine-when-not-synced` flag is only needed for standalone/local mining. When connected to seed nodes, your node will sync existing blocks first, then start mining safely.

**What you'll see:**

```
[INF] KSMN: Found block 761dc037... with parents [97b58165...]
[INF] KSMN: Submitting block 761dc037... to 127.0.0.1:17210
[INF] KSMN: Current hash rate is 160.02 Khash/s
```

### Step 5: Check Your Balance (Terminal 4 - Run When Needed)

```bash
# Check wallet balance
./stokeswallet --testnet balance

# Check block count
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetBlockCount
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
  --rpcserver=127.0.0.1:17210 \
  --mine-when-not-synced
```

### Mining Performance

**Expected hash rates:**

- CPU (4 cores): ~100-200 Khash/s
- CPU (8 cores): ~200-400 Khash/s
- CPU (16 cores): ~400-800 Khash/s

**Mining rewards:**

- Current: **0.0016667 STKS** per block (Bitcoin-equivalent emission)
- Daily emission: **~7,200 STKS/day** (same as Bitcoin's 7,200 BTC/day)
- Coinbase maturity: **100 blocks** (coins spendable after 100 confirmations)
- Block time: ~1 second average (50 blocks/sec throughput)

### Checking Mining Progress

````bash
# Check how many blocks you've mined
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetBlockCount

# Check your balance (includes pending rewards)
Ensure your wallet daemon is running:

```bash
./stokeswallet --testnet start-daemon \
   ~/stokes-wallet-testnet/keys.json \
  -s 127.0.0.1:17210
````

./stokeswallet --testnet balance

# Monitor miner output for:

# [INF] KSMN: Found block ...

# [INF] KSMN: Current hash rate is X Khash/s

````

## 🔧 Advanced Configuration

### Running a Public Node

To help the network, run a public node that accepts connections:

```bash
./stokesd --testnet \
  --utxoindex \
  --listen=0.0.0.0:17711 \
  --rpclisten=0.0.0.0:17210 \
  --externalip=YOUR_PUBLIC_IP
````

**Firewall configuration:**

```bash
# Allow P2P connections
sudo ufw allow 17711/tcp

# Allow RPC (only if you want remote access)
# WARNING: Secure your RPC with authentication!
# sudo ufw allow 17210/tcp
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

### Can't Connect to Seed Nodes

**Problem:** Connection errors or timeouts when trying to connect to seed nodes

**Error messages you might see:**

```
[ERR] TXMP: status error from connectionLoops for 95.216.155.253:17711
error reading from server: read tcp ... connection timeout
```

**Cause:** Your firewall, ISP, or corporate network may be blocking port 17711.

**Solution:** Use standalone mode instead:

```bash
./stokesd --testnet --utxoindex --nodnsseed --nolisten
```

Standalone mode works identically for testing and mining - you just won't sync blocks with other nodes.

### Node Won't Start

**Problem:** Node fails to start or crashes immediately

**Solutions:**

```bash
# 1. Check if port is already in use
lsof -i :17210

# 2. Kill existing process
pkill stokesd

# 3. Reset database (WARNING: deletes blockchain data)
rm -rf ~/Library/Application\ Support/Kaspad/stokes-testnet  # macOS
rm -rf ~/.kaspad/stokes-testnet  # Linux

# 4. Start fresh
./stokesd --testnet --utxoindex --nodnsseed --nolisten
```

### Wallet Can't Connect

**Problem:** Wallet daemon shows connection errors

**Solutions:**

```bash
# 1. Verify node is running
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetBlockCount

# 2. Check node started with --utxoindex flag
# Restart node with:
./stokesd --testnet --utxoindex

# 3. Restart wallet daemon
pkill stokeswallet
./stokeswallet --testnet start-daemon \
  -f ~/stokes-wallet-testnet/keys.json \
  -s 127.0.0.1:17210
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
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetBlockCount

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
| Testnet | 17711    | 17210    |
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
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetInfo

# Get block DAG info
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetBlockDagInfo

# Get peer info
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetConnectedPeerInfo

# Get mempool info
./stokesctl --testnet --rpcserver=127.0.0.1:17210 GetMempoolEntries
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
- **Website:** [https://stokesnetwork.github.io/stokes](https://stokesnetwork.github.io/stokes)
- **Discord:** [https://discord.gg/stokes](https://discord.gg/stokes)

**Coming Soon:**

- Telegram group
- Block explorer
- Mining pools

## License

Stokes is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).

## Acknowledgments

Stokes is built on the foundation of [Kaspa](https://github.com/kaspanet/kaspad), implementing the PHANTOM protocol for blockDAG consensus.
