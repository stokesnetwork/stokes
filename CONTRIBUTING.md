# Contributing to Stokes

Thank you for your interest in contributing to Stokes! We're building a fair-launch cryptocurrency that combines Bitcoin's proven halving economics with Kaspa's innovative blockDAG technology.

Every contribution matters - whether it's code, documentation, testing, or community support. Let's build something great together! 🚀

---

## 🎯 Our Vision

Stokes is built on these core principles:

- **Fair Launch** - No premine, no ICO, no insider advantages
- **Transparency** - Open development and honest communication
- **Innovation** - Best of Bitcoin + Kaspa + Ethereum
- **Community-Driven** - Your voice matters

---

## 🚀 Getting Started

### First Time Contributors

Welcome! Here's how to get started:

1. **Read the documentation**
   - [README.md](README.md) - Project overview
   - [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) - Community guidelines
   - [Website](https://sam-stokes.github.io/stokes) - Quick start guide

2. **Set up your development environment**
   ```bash
   # Clone the repository
   git clone https://github.com/Sam-Stokes/stokes
   cd stokes
   
   # Build the project
   go build -o stokesd .
   go build -o stokesctl ./cmd/stokesctl
   go build -o stokesminer ./cmd/stokesminer
   go build -o stokeswallet ./cmd/stokeswallet
   
   # Run tests
   go test ./...
   ```

3. **Find something to work on**
   - Browse [open issues](https://github.com/Sam-Stokes/stokes/issues)
   - Look for issues tagged `good first issue` or `help wanted`
   - Ask in discussions if you need guidance

### Experienced Contributors

If you're familiar with blockchain development:

- Review the codebase and suggest improvements
- Help with code reviews
- Improve test coverage
- Optimize performance
- Enhance documentation

---

## 💡 Ways to Contribute

### 1. Code Contributions

**What we need:**
- Bug fixes
- Performance improvements
- Test coverage
- Documentation
- New features (discuss first!)

**Before you code:**
- Check if an issue exists, or create one
- Discuss major changes before implementing
- Follow existing code patterns
- Write tests for your changes

### 2. Testing & Bug Reports

**Help us by:**
- Running testnet nodes
- Mining and reporting issues
- Testing wallet functionality
- Documenting edge cases
- Providing detailed bug reports

**Good bug reports include:**
- Clear description of the issue
- Steps to reproduce
- Expected vs actual behavior
- System information (OS, Go version)
- Relevant logs or screenshots

### 3. Documentation

**We need help with:**
- Improving README and guides
- Writing tutorials
- Translating documentation
- Creating video guides
- Updating website content

### 4. Community Support

**You can help by:**
- Answering questions in discussions
- Helping newcomers get started
- Sharing your Stokes experience
- Creating community content
- Spreading the word

---

## 🔧 Development Workflow

### 1. Fork & Clone

```bash
# Fork the repository on GitHub, then:
git clone https://github.com/YOUR_USERNAME/stokes
cd stokes
git remote add upstream https://github.com/Sam-Stokes/stokes
```

### 2. Create a Branch

```bash
# Create a feature branch
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/bug-description
```

### 3. Make Your Changes

- Write clean, readable code
- Follow Go best practices
- Add comments for complex logic
- Update documentation if needed
- Write or update tests

### 4. Test Your Changes

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run specific package tests
go test ./domain/consensus/...

# Format your code
go fmt ./...
```

### 5. Commit Your Changes

```bash
# Stage your changes
git add .

# Commit with a clear message
git commit -m "feat: add halving verification test

- Add test for Bitcoin-style halving calculation
- Verify reward decreases correctly
- Test edge cases for halving boundaries"
```

**Commit message format:**
- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `test:` Adding or updating tests
- `refactor:` Code refactoring
- `perf:` Performance improvements
- `chore:` Maintenance tasks

### 6. Push & Create Pull Request

```bash
# Push to your fork
git push origin feature/your-feature-name

# Then create a Pull Request on GitHub
```

---

## 📝 Pull Request Guidelines

### Before Submitting

- [ ] Code compiles without errors
- [ ] All tests pass
- [ ] New code has tests
- [ ] Documentation is updated
- [ ] Code follows existing patterns
- [ ] Commit messages are clear

### PR Description Should Include

- **What:** Brief description of changes
- **Why:** Reason for the changes
- **How:** Approach taken
- **Testing:** How you tested it
- **Screenshots:** If UI changes

### Example PR Description

```markdown
## What
Implement Bitcoin-style halving verification tests

## Why
Need comprehensive tests to ensure halving logic works correctly across all eras

## How
- Added test cases for each halving era
- Verified reward calculation at boundaries
- Tested edge cases (64+ halvings)

## Testing
- All new tests pass
- Existing tests still pass
- Tested manually on testnet
```

### Review Process

1. **Automated checks** - CI must pass
2. **Code review** - Maintainers will review
3. **Feedback** - Address any comments
4. **Approval** - Once approved, we'll merge
5. **Thank you!** - Your contribution is live! 🎉

---

## 🧪 Testing Guidelines

### Running Tests

```bash
# All tests
go test ./...

# Specific package
go test ./domain/consensus/processes/coinbasemanager

# With coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

### Writing Tests

- Test files end with `_test.go`
- Use table-driven tests when appropriate
- Test both success and failure cases
- Include edge cases
- Keep tests focused and clear

### Example Test

```go
func TestHalvingCalculation(t *testing.T) {
    tests := []struct {
        name     string
        block    uint64
        expected uint64
    }{
        {"era 0", 0, 50 * constants.SompiPerKaspa},
        {"era 1", 126230400, 25 * constants.SompiPerKaspa},
        {"era 2", 252460800, 12.5 * constants.SompiPerKaspa},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := calcHalvingBlockSubsidy(tt.block)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

---

## 📚 Code Style

### Go Guidelines

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Keep functions focused and small
- Write clear variable names
- Add comments for exported functions
- Handle errors explicitly

### Naming Conventions

- **Files:** `lowercase_with_underscores.go`
- **Packages:** `lowercase` (single word)
- **Functions:** `CamelCase` (exported) or `camelCase` (internal)
- **Constants:** `CamelCase` or `UPPER_CASE`
- **Variables:** `camelCase`

### Comments

```go
// CalcBlockSubsidy calculates the block subsidy based on the current
// block DAA score, implementing Bitcoin-style halving every ~4 years.
//
// The subsidy starts at 50 STKS and halves every 126,230,400 blocks.
// After 64 halvings, the subsidy becomes 0.
func CalcBlockSubsidy(blockDaaScore uint64) uint64 {
    // Implementation...
}
```

---

## 🤝 Communication

### Where to Ask Questions

- **GitHub Discussions** - General questions and ideas
- **GitHub Issues** - Bug reports and feature requests
- **Discord** - Real-time chat (coming soon)
- **Telegram** - Community chat (coming soon)

### Getting Help

Don't hesitate to ask for help! We're here to support you:

- Not sure where to start? Ask in discussions
- Stuck on implementation? Open a draft PR and ask for guidance
- Need clarification? Comment on the issue

---

## 🎯 Project Priorities

### Current Focus (Testnet Phase)

1. **Stability** - Fix bugs, improve reliability
2. **Testing** - Increase test coverage
3. **Documentation** - Make it easier to contribute
4. **Performance** - Optimize critical paths
5. **Community** - Build a strong foundation

### Future Goals

- Mainnet launch preparation
- Block explorer
- Mining pools
- Exchange listings
- Ecosystem tools

---

## 🏆 Recognition

We value every contribution! Contributors will be:

- Listed in release notes
- Mentioned in the README
- Recognized in the community
- Part of Stokes history

---

## ❓ FAQ

**Q: Do I need to be an expert to contribute?**  
A: No! We welcome contributors of all skill levels.

**Q: How long does PR review take?**  
A: Usually within a few days, but it depends on complexity.

**Q: Can I work on multiple issues?**  
A: Yes, but focus on one at a time for best results.

**Q: What if my PR isn't accepted?**  
A: We'll explain why and suggest improvements. Don't be discouraged!

**Q: Can I contribute if I don't code?**  
A: Absolutely! Testing, documentation, and community support are crucial.

---

## 📜 License

By contributing to Stokes, you agree that your contributions will be licensed under the [ISC License](LICENSE).

---

## 🙏 Thank You

Every contribution, no matter how small, helps make Stokes better. Thank you for being part of this journey!

**Let's build the future of fair-launch cryptocurrencies together.** 🚀

---

**Questions?** Open a discussion or reach out to the maintainers.

**Last Updated:** October 14, 2025