package dagconfig

import (
	"github.com/stokesnetwork/stokes/domain/consensus/utils/constants"
	"time"
)

// The documentation refers to the following constants which aren't explicated in the code:
//	d - an upper bound on the round trip time of a block
//	delta - the expected fraction of time the width of the network exceeds defaultGHOSTDAGK
//
// For more information about defaultGHOSTDAGK, and its dependency on delta and defaultTargetTimePerBlock
// please refer to the PHANTOM paper: https://eprint.iacr.org/2018/104.pdf
//
// For more information about the DAA constants defaultDifficultyAdjustmentWindowSize, defaultTimestampDeviationTolerance,
// and their relation to defaultGHOSTDAGK and defaultTargetTimePerBlock see:
// https://research.kas.pa/t/handling-timestamp-manipulations/97
//
// For more information about defaultMergeSetSizeLimit, defaultFinalityDuration and their relation to pruning, see:
// https://research.kas.pa/t/a-proposal-for-finality-in-ghostdag/66/17
// https://research.kas.pa/t/some-of-the-intuition-behind-the-design-of-the-invalidation-rules-for-pruning/95
//

const (
	defaultMaxCoinbasePayloadLength = 204
	// defaultMaxBlockMass is a bound on the mass of a block, larger values increase the bound d
	// on the round trip time of a block, which affects the other parameters as described below
	defaultMaxBlockMass = 500_000
	// defaultMassPerTxByte, defaultMassPerScriptPubKeyByte and defaultMassPerSigOp define the number of grams per
	// transaction byte, script pub key byte and sig op respectively.
	// These values are used when calculating a transactions mass.
	defaultMassPerTxByte           = 1
	defaultMassPerScriptPubKeyByte = 10
	defaultMassPerSigOp            = 1000
	// defaultMaxBlockParents is the number of blocks any block can point to.
	// Should be about d/defaultTargetTimePerBlock where d is a bound on the round trip time of a block.
	defaultMaxBlockParents = 10
	// defaultGHOSTDAGK is a bound on the number of blue blocks in the anticone of a blue block. Approximates the maximal
	// width of the network.
	// Formula (1) in section 4.2 of the PHANTOM paper shows how to calculate defaultGHOSTDAGK. The delta term represents a bound
	// on the expected fraction of the network life in which the width was higher than defaultGHOSTDAGK. The current value of K
	// was calculated for d = 5 seconds and delta = 0.05.
	defaultGHOSTDAGK = 18
	// defaultMergeSetSizeLimit is a bound on the size of the past of a block and the size of the past
	// of its selected parent. Any block which violates this bound is invalid.
	// Should be at least an order of magnitude smaller than defaultFinalityDuration/defaultTargetTimePerBlock.
	// (Higher values make pruning attacks easier by a constant, lower values make merging after a split or a spike
	// in block take longer)
	defaultMergeSetSizeLimit                       = defaultGHOSTDAGK * 10
	
	// STOKES: Bitcoin-style halving emission parameters (TESTNET/MAINNET)
	// Genesis block reward (first block only)
	defaultSubsidyGenesisReward                    = 50 * constants.SompiPerStokes
	// Initial block reward: 50 STKS (matches Bitcoin's initial reward)
	defaultPreDeflationaryPhaseBaseSubsidy         = 50 * constants.SompiPerStokes
	// This is unused in halving model but kept for compatibility
	defaultDeflationaryPhaseBaseSubsidy            = 50 * constants.SompiPerStokes
	
	// STOKES: DEVNET-specific emission parameters (100M max supply)
	// Devnet uses a much lower reward to test the new economic model
	// Calculation for 100M max supply with 4-year halving:
	//   - Halving interval: 6,311,520,000 blocks (4 years at 50 BPS)
	//   - Geometric series: max_supply = first_period × 2
	//   - First period: 100M / 2 = 50M STKS
	//   - Reward per block: 50M / 6,311,520,000 = 0.007922022 STKS/block
	//   - In sompi: 0.007922022 × 100,000,000 = 792,202.2 sompi ≈ 792,202 sompi
	// Blocks per second: 50 (20ms per block)
	// Daily emission: ~34,239 STKS (4,320,000 blocks/day × 0.007922022)
	devnetSubsidyGenesisReward                     = 792202 // 0.007922022 STKS in sompi
	devnetPreDeflationaryPhaseBaseSubsidy          = 792202 // 0.007922022 STKS in sompi
	devnetDeflationaryPhaseBaseSubsidy             = 792202 // 0.007922022 STKS in sompi
	
	defaultCoinbasePayloadScriptPublicKeyMaxLength = 150
	// defaultDifficultyAdjustmentWindowSize is the number of blocks in a block's past used to calculate its difficulty
	// target.
	// The DAA should take the median of 2640 blocks, so in order to do that we need 2641 window size.
	defaultDifficultyAdjustmentWindowSize = 2641
	// defaultTimestampDeviationTolerance is the allowed deviance of an inconming block's timestamp, measured in block delays.
	// A new block can't hold a timestamp lower than the median timestamp of the (defaultTimestampDeviationTolerance*2-1) blocks
	// with highest accumulated blue work in its past, such blocks are considered invalid.
	// A new block can't hold a timestamp higher than the local system time + defaultTimestampDeviationTolerance/defaultTargetTimePerBlock,
	// such blocks are not marked as invalid but are rejected.
	defaultTimestampDeviationTolerance = 132
	// defaultFinalityDuration is an approximate lower bound of how old the finality block is. The finality block is chosen to
	// be the newest block in the selected chain whose blue score difference from the selected tip is at least
	// defaultFinalityDuration/defaultTargetTimePerBlock.
	// The pruning block is selected similarly, with the following duration:
	//	pruning block duration =
	//		2*defaultFinalityDuration/defaultTargetTimePerBlock + 4*defaultMergeSetSizeLimit*defaultGHOSTDAGK + 2*defaultGHOSTDAGK + 2
	defaultFinalityDuration = 24 * time.Hour
	// defaultTargetTimePerBlock represents how much time should pass on average between two consecutive block creations.
	// Should be parametrized such that the average width of the DAG is about defaultMaxBlockParents and such that most of the
	// time the width of the DAG is at most defaultGHOSTDAGK.
	defaultTargetTimePerBlock = 1 * time.Second
	
	// STOKES: DEVNET uses 50 blocks/second (20ms per block) for high-throughput testing
	devnetTargetTimePerBlock = 20 * time.Millisecond

	defaultPruningProofM = 1000

	// STOKES: Halving interval configuration (TESTNET/MAINNET)
	// Bitcoin halves every 210,000 blocks (~4 years at 10 min/block)
	// STOKES targets 1 block/second, so 4 years = 4 * 365.25 * 24 * 60 * 60 = 126,230,400 seconds
	// We use DAA score (which approximates seconds) for halving intervals
	defaultHalvingIntervalDaaScore = uint64(4 * 365.25 * 24 * 60 * 60) // ~126.23M blocks (4 years)
	
	// STOKES: DEVNET halving interval (50 blocks/second)
	// 4 years at 50 blocks/second = 4 * 365.25 * 24 * 60 * 60 * 50 = 6,311,520,000 blocks
	devnetHalvingIntervalDaaScore = uint64(4 * 365.25 * 24 * 60 * 60 * 50) // ~6.31B blocks (4 years at 50 BPS)
	
	// defaultDeflationaryPhaseDaaScore is kept for compatibility but unused in halving model
	// In STOKES, halving starts immediately after genesis (no pre-deflationary phase)
	defaultDeflationaryPhaseDaaScore = 0

	defaultMergeDepth = 3600
)
