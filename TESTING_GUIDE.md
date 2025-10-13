# STOKES Testing Guide

Complete guide for testing STOKES blockchain functionality.

---

## 🧪 Quick Start Testing

### Prerequisites
- Go 1.21+ installed
- STOKES repository cloned
- Binaries built

### Build Binaries
```bash
cd /path/to/stokes

# Build all binaries
go build -o stokesd .
go build -o stokesctl ./cmd/stokesctl
go build -o stokesminer ./cmd/stokesminer
go build -o stokeswallet ./cmd/stokeswallet
```

---

## 🔬 Test Scenarios

### Test 1: Compilation Test

**Objective:** Verify code compiles without errors

**Steps:**
```bash
# Clean build
go clean ./...

# Build all packages
go build ./...

# Check exit code
echo $?  # Should be 0
```

**Expected Result:**
- Exit code: 0
- No compilation errors
- All dependencies resolved

**Time:** 1-2 minutes

---

### Test 2: Single Node Test

**Objective:** Verify node starts and runs correctly

**Terminal 1 - Start Node:**
```bash
./stokesd --simnet \
  --appdir=./test-data \
  --logdir=./test-data/logs \
  --utxoindex
```

**Expected Output:**
```
[INF] KASD: Version 0.12.22
[INF] KASD: Loading database from 'test-data/stokes-simnet/datadir2'
[INF] KASD: UTXO index started
[INF] TXMP: P2P Server listening on [::]:17511
[INF] TXMP: RPC Server listening on [::]:17510
```

**Verification:**
- ✅ Node starts without errors
- ✅ P2P server listening on port 17511
- ✅ RPC server listening on port 17510
- ✅ UTXO index started

**Time:** 30 seconds

---

### Test 3: Wallet Creation Test

**Objective:** Verify wallet functionality

**Terminal 2 - Create Wallet:**
```bash
# Create wallet
./stokeswallet create --simnet

# Enter password when prompted
# Save the extended public key shown
```

**Expected Output:**
```
Enter password for the key file: ****
Confirm password: ****
Extended public key of mnemonic #1:
ksub8L39S3JnPGziQii6SpXt9HRvyigaWKbcZEu2NL4uvRwiA5urie8uXbJLxeTkt2URCFygQNENQikMjgZkNbUB3UY3EQi34XGtyGiAnofNELz

Wrote the keys into /Users/.../Kaspawallet/stokes-simnet/keys.json
```

**Generate Address:**
```bash
./stokeswallet new-address --simnet
```

**Expected Output:**
```
New address:
kaspasim:qpts2h78d9rpnsl8j0pvszrslh09ean3slzf89cvaumpw28hrmkdvv39jafyy
```

**Verification:**
- ✅ Wallet created successfully
- ✅ Keys file written
- ✅ Address generated
- ✅ Address starts with "kaspasim:"

**Time:** 1 minute

---

### Test 4: Two-Node Network Test

**Objective:** Verify nodes can connect to each other

**Terminal 1 - Node 1 (Primary):**
```bash
./stokesd --simnet \
  --appdir=./node1-data \
  --logdir=./node1-data/logs \
  --utxoindex \
  --nodnsseed \
  --listen=0.0.0.0:17511 \
  --rpclisten=127.0.0.1:17510
```

**Terminal 2 - Node 2 (Secondary):**
```bash
./stokesd --simnet \
  --appdir=./node2-data \
  --logdir=./node2-data/logs \
  --utxoindex \
  --nodnsseed \
  --listen=0.0.0.0:17512 \
  --rpclisten=127.0.0.1:17513 \
  --connect=127.0.0.1:17511
```

**Expected Output (Node 1):**
```
[INF] TXMP: P2P Incoming connection from 127.0.0.1:xxxxx #1
[INF] PROT: Registering p2p flows for peer <...> for protocol version 5
```

**Expected Output (Node 2):**
```
[INF] CMGR: Connecting to 127.0.0.1:17511
[INF] TXMP: P2P Connected to 127.0.0.1:17511
[INF] PROT: Registering p2p flows for peer <...> for protocol version 5
```

**Verification:**
- ✅ Node 2 connects to Node 1
- ✅ P2P flows registered
- ✅ Protocol version 5 negotiated
- ✅ No connection errors

**Time:** 2 minutes

---

### Test 5: Mining Test

**Objective:** Verify mining functionality and block creation

**Prerequisites:**
- Node 1 running (from Test 4)
- Wallet created with address (from Test 3)

**Terminal 3 - Start Miner:**
```bash
./stokesminer --simnet \
  --miningaddr=kaspasim:qpts2h78d9rpnsl8j0pvszrslh09ean3slzf89cvaumpw28hrmkdvv39jafyy \
  --rpcserver=127.0.0.1:17510 \
  --mine-when-not-synced
```

**Expected Output:**
```
[INF] KSMN: Version 0.12.22
[INF] KSMN: Connected to 127.0.0.1:17510
[INF] KSMN: Minimum average time per 10 blocks: 5ms
[INF] KSMN: Waiting for the initial template
[INF] KSMN: Found block f4df887cf1c6a54f... with parents [3d7f1715e6f7c2747...]
[INF] KSMN: Found block 122338ce850e50cf... with parents [3d7f1715e6f7c2747...]
[INF] KSMN: Found block 67f3c1702f4b9636... with parents [3d7f1715e6f7c2747...]
```

**Verification:**
- ✅ Miner connects to node
- ✅ Blocks are found
- ✅ Parent hash is STOKES simnet genesis: `3d7f1715e6f7c2747...`
- ✅ Multiple blocks mined successfully

**Note:** Blocks may be rejected with "node is in IBD" - this is expected in test environment

**Time:** 2-3 minutes

---

### Test 6: Genesis Verification Test

**Objective:** Verify all blocks use STOKES genesis, not Kaspa's

**Check Genesis Hash:**
```bash
# STOKES Simnet Genesis
3d7f1715e6f7c2744730462226a37c196d879f7391cdbcf8d28efe68e2655c779

# Kaspa Simnet Genesis (should NOT appear)
# Different hash - if you see this, something is wrong!
```

**Verification in Miner Output:**
```
Found block ... with parents [3d7f1715e6f7c2747...]
                              ^^^^^^^^^^^^^^^^^^^^
                              This should match STOKES genesis
```

**Verification:**
- ✅ All blocks parent to `3d7f1715e6f7c2747...`
- ✅ No Kaspa genesis references
- ✅ Network completely isolated

**Time:** 1 minute (visual inspection)

---

### Test 7: Emission Schedule Verification

**Objective:** Verify 50 STKS block reward

**Check in Code:**
```bash
# View emission logic
cat domain/consensus/processes/coinbasemanager/coinbasemanager.go | grep -A 10 "calcBlockSubsidy"
```

**Expected Values:**
- Initial reward: 50 STKS = 5,000,000,000 sompi
- Halving interval: 126,230,400 blocks
- After halving: 25 STKS = 2,500,000,000 sompi

**Verification:**
- ✅ Code shows 50 STKS initial reward
- ✅ Halving interval correct
- ✅ Halving logic implemented

**Time:** 2 minutes

---

## 🔧 Advanced Testing

### Test 8: RPC API Test

**Test RPC Endpoints:**
```bash
# Get block count
./stokesctl --simnet --rpcserver=127.0.0.1:17510 getblockcount

# Get block DAA score
./stokesctl --simnet --rpcserver=127.0.0.1:17510 getblockdaginfo

# Get peer info
./stokesctl --simnet --rpcserver=127.0.0.1:17510 getpeerinfo
```

**Expected Results:**
- Commands execute without errors
- Valid JSON responses
- Data matches network state

---

### Test 9: Stress Test

**Objective:** Test network under load

**Setup:**
- 3+ nodes running
- 2+ miners running
- Let run for 30+ minutes

**Monitor:**
- CPU usage
- Memory usage
- Disk I/O
- Network bandwidth
- Block propagation time

**Success Criteria:**
- No crashes
- Stable memory usage
- Blocks propagate quickly
- No data corruption

---

### Test 10: Network Partition Test

**Objective:** Test network resilience

**Steps:**
1. Start 4 nodes (A, B, C, D)
2. Connect: A-B, C-D (two separate networks)
3. Mine blocks on both networks
4. Reconnect: B-C
5. Observe chain reorganization

**Expected Behavior:**
- Both networks mine independently
- When reconnected, nodes sync
- Longest chain wins
- No data loss

---

## 📊 Test Results Template

### Test Report Format

```markdown
# Test Report: [Test Name]
**Date:** YYYY-MM-DD
**Tester:** [Name]
**Environment:** [OS, Hardware]

## Test Configuration
- STOKES Version: 0.12.22
- Network: Simnet/Testnet/Devnet
- Nodes: X
- Duration: X minutes

## Results
- [ ] Test passed
- [ ] Test failed
- [ ] Partial success

## Observations
[Detailed observations]

## Issues Found
[List any bugs or issues]

## Recommendations
[Suggestions for improvement]
```

---

## 🐛 Common Issues & Solutions

### Issue 1: "Kaspad is not synced"
**Symptom:** Miner shows "Kaspad is not synced. Skipping current block template"

**Solution:** Add `--mine-when-not-synced` flag to miner command

**Why:** Test networks with few nodes may not complete IBD

---

### Issue 2: "Connection refused"
**Symptom:** Miner or wallet can't connect to node

**Solution:** 
1. Verify node is running
2. Check correct port (17510 for simnet RPC)
3. Check firewall settings

---

### Issue 3: Wallet daemon timeout
**Symptom:** Wallet daemon times out during sync

**Solution:**
1. Ensure node is running with `--utxoindex`
2. Wait for node to fully start
3. Restart wallet daemon

---

### Issue 4: "Block rejected - node is in IBD"
**Symptom:** Blocks mined but rejected

**Solution:** This is expected in test environment with 2 nodes. In production with multiple nodes and activity, IBD completes normally.

**Workaround:** Continue mining - blocks are being created correctly

---

## ✅ Testing Checklist

### Pre-Launch Testing
- [ ] Compilation test passed
- [ ] Single node test passed
- [ ] Multi-node test passed
- [ ] Wallet creation test passed
- [ ] Mining test passed
- [ ] Genesis verification passed
- [ ] Emission schedule verified
- [ ] RPC API test passed
- [ ] Stress test passed (30+ min)
- [ ] Network partition test passed

### Security Testing
- [ ] No critical vulnerabilities found
- [ ] Peer review completed
- [ ] Code audit completed (if applicable)
- [ ] All tests passing
- [ ] No memory leaks
- [ ] No data corruption

### Performance Testing
- [ ] Block propagation < 1 second
- [ ] Transaction throughput acceptable
- [ ] Memory usage stable
- [ ] CPU usage reasonable
- [ ] Disk I/O acceptable

---

## 📈 Performance Benchmarks

### Target Metrics
- **Block propagation:** < 1 second
- **Transaction throughput:** 100+ TPS
- **Memory usage:** < 2GB per node
- **CPU usage:** < 50% average
- **Uptime:** 99.9%+

### Measurement Tools
```bash
# Monitor resources
htop

# Check memory
free -h

# Check disk I/O
iostat -x 1

# Network monitoring
iftop
```

---

## 🎓 Testing Best Practices

1. **Test early, test often** - Don't wait until the end
2. **Automate where possible** - Use scripts for repetitive tests
3. **Document everything** - Keep detailed test logs
4. **Test edge cases** - Don't just test happy path
5. **Test under load** - Stress test is critical
6. **Test failure scenarios** - What happens when things go wrong?
7. **Peer review** - Have others test independently
8. **Version control** - Tag tested versions

---

## 🚀 Continuous Testing

### Automated Testing
```bash
# Run all Go tests
go test ./...

# Run with race detector
go test -race ./...

# Run with coverage
go test -cover ./...

# Benchmark tests
go test -bench=. ./...
```

### CI/CD Integration
- Set up GitHub Actions
- Run tests on every commit
- Automated build verification
- Performance regression detection

---

## 📞 Support & Resources

### Getting Help
- Check documentation first
- Review test logs
- Search GitHub issues
- Ask in community Discord/Telegram

### Reporting Issues
When reporting bugs, include:
- STOKES version
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Relevant logs
- Test configuration

---

## 🎊 Conclusion

Thorough testing is critical for blockchain success. Follow this guide to ensure STOKES is stable, secure, and ready for production.

**Remember:** It's better to find bugs in testing than in production!

---

**Last Updated:** October 13, 2025  
**Version:** 1.0  
**Status:** Active
