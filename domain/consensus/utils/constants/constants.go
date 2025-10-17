package constants

import "math"

const (
	// BlockVersion represents the current block version
	BlockVersion uint16 = 1

	// MaxTransactionVersion is the current latest supported transaction version.
	MaxTransactionVersion uint16 = 0

	// MaxScriptPublicKeyVersion is the current latest supported public key script version.
	MaxScriptPublicKeyVersion uint16 = 0

	// SompiPerStokes is the number of sompi in one stokes (1 STKS).
	// STOKES: Renamed from SompiPerKaspa to reflect STOKES branding
	SompiPerStokes = 100_000_000
	
	// SompiPerKaspa is kept for backward compatibility with existing code
	SompiPerKaspa = SompiPerStokes

	// MaxSompi is the maximum transaction amount allowed in sompi.
	// STOKES: Network-specific max supplies:
	//   - TESTNET/MAINNET: ~12.6B STKS (50 STKS/block, 1 BPS, 4-year halving)
	//     Formula: 126,230,400 blocks/halving × 50 STKS × 2 (geometric series) = 12,623,040,000 STKS
	//   - DEVNET: 100M STKS (0.001585 STKS/block, 50 BPS, 4-year halving)
	//     Formula: 6,311,520,000 blocks/halving × 0.001585 STKS × 2 (geometric series) ≈ 100,000,000 STKS
	// This constant is set to the larger value to accommodate all networks
	MaxSompi = uint64(12_623_040_000 * SompiPerStokes)

	// MaxTxInSequenceNum is the maximum sequence number the sequence field
	// of a transaction input can be.
	MaxTxInSequenceNum uint64 = math.MaxUint64

	// SequenceLockTimeDisabled is a flag that if set on a transaction
	// input's sequence number, the sequence number will not be interpreted
	// as a relative locktime.
	SequenceLockTimeDisabled uint64 = 1 << 63

	// SequenceLockTimeMask is a mask that extracts the relative locktime
	// when masked against the transaction input sequence number.
	SequenceLockTimeMask uint64 = 0x00000000ffffffff

	// LockTimeThreshold is the number below which a lock time is
	// interpreted to be a DAA score.
	LockTimeThreshold = 5e11 // Tue Nov 5 00:53:20 1985 UTC

	// UnacceptedDAAScore is used to for UTXOEntries that were created by transactions in the mempool, or otherwise
	// not-yet-accepted transactions.
	UnacceptedDAAScore = math.MaxUint64
)
