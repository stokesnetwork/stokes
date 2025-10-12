# Complete Phase 1 Manually - Step by Step Guide

## What We've Already Done ✅

1. ✅ Changed network magic bytes (prevents Kaspa connection)
2. ✅ Updated all ports (17xxx range)
3. ✅ Removed Kaspa DNS seeds
4. ✅ Updated network names (stokes-*)
5. ✅ Set MaxSompi to 12.6B STKS
6. ✅ Implemented Bitcoin-style halving

## What You Need to Do Now

### STEP 1: Create GitHub Repository

1. Go to https://github.com/new
2. Repository name: `stokes`
3. Description: `STOKES - Kaspa fork with Bitcoin-style halving emission`
4. Choose **Private** (recommended for now)
5. Do **NOT** check "Initialize with README"
6. Click "Create repository"

**Write down your GitHub username**: ___________________

### STEP 2: Update go.mod

```bash
cd /Users/sam/srv/playground/stokes

# Open go.mod in editor
nano go.mod

# Change line 1 from:
module github.com/kaspanet/kaspad

# To (replace YOUR_USERNAME):
module github.com/YOUR_USERNAME/stokes

# Save and exit (Ctrl+X, Y, Enter)
```

### STEP 3: Update All Import Statements

```bash
# Replace YOUR_USERNAME with your actual GitHub username
find . -name "*.go" -type f -exec sed -i '' 's|github.com/kaspanet/kaspad|github.com/YOUR_USERNAME/stokes|g' {} +

# This updates all .go files at once
```

### STEP 4: Rename Binary Directories

```bash
cd /Users/sam/srv/playground/stokes/cmd

# Rename the directories
mv kaspactl stokesctl
mv kaspaminer stokesminer
mv kaspawallet stokeswallet

cd ..
```

### STEP 5: Update Git Remote

```bash
# Remove old remote (points to Kaspa)
git remote remove origin

# Add your new remote (replace YOUR_USERNAME)
git remote add origin https://github.com/YOUR_USERNAME/stokes.git

# Verify
git remote -v
```

### STEP 6: Commit Your Changes

```bash
# Stage all changes
git add .

# Commit
git commit -m "Phase 1: Network isolation and rebranding complete

- Changed network magic bytes to prevent Kaspa conflicts
- Updated ports to 17xxx range
- Removed Kaspa DNS seeds
- Updated network names to stokes-*
- Set MaxSompi to 12.6B STKS
- Implemented Bitcoin-style halving
- Updated module name and imports
- Renamed binaries (kaspa* → stokes*)"
```

### STEP 7: Push to GitHub

```bash
# Push to your new repository
git push -u origin master

# If that fails, try:
git push -u origin main

# You may need to authenticate with GitHub
```

### STEP 8: Test Compilation (If Go is Installed)

```bash
# Update dependencies
go mod tidy

# Try to build
go build ./...

# If successful, you'll see no output
# If errors, note them down - we'll fix them
```

---

## Troubleshooting

### Error: "go: module github.com/YOUR_USERNAME/stokes: git ls-remote"

**Solution**: This is normal! Your imports reference your new repo, but Go can't find it yet because you just created it. This will resolve after you push.

### Error: "package github.com/YOUR_USERNAME/stokes/... is not in GOROOT"

**Solution**: Run `go mod tidy` after pushing to GitHub. Go needs to download dependencies.

### Error: "fatal: remote origin already exists"

**Solution**: 
```bash
git remote remove origin
git remote add origin https://github.com/YOUR_USERNAME/stokes.git
```

### Error: "failed to push some refs"

**Solution**: Make sure you created the GitHub repository first, and it's empty (no README).

---

## Verification Checklist

After completing all steps, verify:

```bash
# 1. Check go.mod has your username
head -1 go.mod
# Should show: module github.com/YOUR_USERNAME/stokes

# 2. Check git remote
git remote -v
# Should show your GitHub repo

# 3. Check binary directories
ls cmd/
# Should show: genkeypair  stokesctl  stokesminer  stokeswallet

# 4. Check a sample import
grep "github.com" domain/consensus/factory.go | head -5
# Should show your GitHub username, not kaspanet

# 5. Check network names
grep "Name:" domain/dagconfig/params.go | head -4
# Should show stokes-mainnet, stokes-testnet, etc.
```

---

## What's Still Missing (Phase 2)

### 🔴 CRITICAL - Genesis Block

**Current Status**: Still using Kaspa's genesis block

**Why Critical**: Your chain needs its own genesis to be truly independent

**What to Do**:
1. Don't launch publicly yet
2. Genesis generation is complex - needs separate tool
3. For now, you can test privately with existing genesis
4. Before mainnet launch, MUST generate new genesis

**Files to Update Later**:
- `domain/dagconfig/genesis.go`
- All network params (GenesisBlock, GenesisHash)

### 🟡 IMPORTANT - Additional Rebranding

**Remaining "kaspa" references**: ~1000+

**Critical ones**:
- Config directory paths (`~/.kaspad` → `~/.stokesd`)
- Log messages
- Error messages
- User-facing strings

**Can wait**:
- Code comments
- Internal variable names
- Documentation

---

## Quick Command Reference

```bash
# See what changed
git status

# See current remote
git remote -v

# Check module name
head -1 go.mod

# Count remaining "kaspa" references
grep -r "kaspa" --include="*.go" . | wc -l

# Test build (if Go installed)
go build ./...

# Run tests (if Go installed)
go test ./...
```

---

## Next Steps After Phase 1

1. **Test Compilation**
   - Install Go if needed
   - Fix any compilation errors
   - Run unit tests

2. **Genesis Block**
   - Research genesis generation
   - Create generation tool
   - Generate new genesis for all networks

3. **Private Testing**
   - Deploy devnet
   - Mine test blocks
   - Verify emission schedule

4. **Complete Rebranding**
   - Update remaining "kaspa" references
   - Update documentation
   - Create new README

5. **Security Audit**
   - Review all changes
   - Test edge cases
   - Consider professional audit

---

## Getting Help

If you get stuck:

1. **Check PHASE1_PROGRESS.md** - Detailed status
2. **Check error messages** - Usually self-explanatory
3. **Search GitHub Issues** - Others may have similar issues
4. **Ask in Discord/Community** - Crypto dev communities are helpful

---

## Important Reminders

⚠️ **DO NOT launch publicly yet** - Genesis block issue must be resolved

⚠️ **Keep repo private** - Until you're ready for public testing

⚠️ **Test thoroughly** - Use devnet/simnet for testing

⚠️ **Backup everything** - Keep copies of important changes

✅ **You're making progress** - Phase 1 is 50% done!

---

## Summary of Commands

```bash
# Complete workflow
cd /Users/sam/srv/playground/stokes

# 1. Edit go.mod (change module name)
nano go.mod

# 2. Update imports (replace YOUR_USERNAME)
find . -name "*.go" -type f -exec sed -i '' 's|github.com/kaspanet/kaspad|github.com/YOUR_USERNAME/stokes|g' {} +

# 3. Rename directories
cd cmd && mv kaspactl stokesctl && mv kaspaminer stokesminer && mv kaspawallet stokeswallet && cd ..

# 4. Update git remote (replace YOUR_USERNAME)
git remote remove origin
git remote add origin https://github.com/YOUR_USERNAME/stokes.git

# 5. Commit and push
git add .
git commit -m "Phase 1: Network isolation complete"
git push -u origin master
```

---

**You're doing great! Keep going! 🚀**
