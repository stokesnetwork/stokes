package dagconfig

import (
	"testing"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/constants"
)

// TestDevnetRewardCalculations verifies the devnet economic model
// Option 2: 0.007922022 STKS/block with 4-year halving for 100M max supply
func TestDevnetRewardCalculations(t *testing.T) {
	// Verify devnet reward is 0.007922022 STKS = 792,202 sompi
	expectedReward := uint64(792202)
	if devnetPreDeflationaryPhaseBaseSubsidy != expectedReward {
		t.Errorf("Devnet reward should be %d sompi (0.007922022 STKS), got %d", 
			expectedReward, devnetPreDeflationaryPhaseBaseSubsidy)
	}

	// Verify blocks per second calculation
	// 50 blocks/second = 20ms per block
	expectedBlockTime := uint64(20) // milliseconds
	actualBlockTime := uint64(devnetTargetTimePerBlock.Milliseconds())
	if actualBlockTime != expectedBlockTime {
		t.Errorf("Devnet block time should be %dms (50 BPS), got %dms", 
			expectedBlockTime, actualBlockTime)
	}

	// Verify blocks per day
	// 50 blocks/sec × 60 sec/min × 60 min/hr × 24 hr/day = 4,320,000 blocks/day
	blocksPerDay := uint64(50 * 60 * 60 * 24)
	expectedBlocksPerDay := uint64(4_320_000)
	if blocksPerDay != expectedBlocksPerDay {
		t.Errorf("Expected %d blocks per day, calculated %d", expectedBlocksPerDay, blocksPerDay)
	}

	// Verify daily emission
	// 4,320,000 blocks/day × 0.007922022 STKS/block = ~34,239 STKS/day
	dailyEmissionSompi := blocksPerDay * devnetPreDeflationaryPhaseBaseSubsidy
	dailyEmissionSTKS := float64(dailyEmissionSompi) / float64(constants.SompiPerStokes)
	expectedDailySTKS := 34239.0
	
	// Allow 100 STKS tolerance for rounding
	if dailyEmissionSTKS < expectedDailySTKS-100 || dailyEmissionSTKS > expectedDailySTKS+100 {
		t.Errorf("Expected ~%.0f STKS/day, got %.1f STKS/day", 
			expectedDailySTKS, dailyEmissionSTKS)
	} else {
		t.Logf("✓ Daily emission: %.1f STKS/day", dailyEmissionSTKS)
	}

	// Verify halving interval
	// 4 years at 50 blocks/second = 4 × 365.25 × 24 × 60 × 60 × 50 = 6,311,520,000 blocks
	expectedHalvingInterval := uint64(4 * 365.25 * 24 * 60 * 60 * 50)
	if devnetHalvingIntervalDaaScore != expectedHalvingInterval {
		t.Errorf("Devnet halving interval should be %d blocks, got %d", 
			expectedHalvingInterval, devnetHalvingIntervalDaaScore)
	}

	// Verify max supply calculation
	// First halving period: 6,311,520,000 blocks × 0.007922022 STKS = 50,000,000 STKS
	// Geometric series sum: 50,000,000 × 2 = 100,000,000 STKS
	firstHalvingEmissionSompi := devnetHalvingIntervalDaaScore * devnetPreDeflationaryPhaseBaseSubsidy
	firstHalvingEmissionSTKS := float64(firstHalvingEmissionSompi) / float64(constants.SompiPerStokes)
	
	// The geometric series sum for infinite halvings is: firstPeriod × 2
	totalSupplySTKS := firstHalvingEmissionSTKS * 2
	expectedMaxSupply := 100_000_000.0 // 100M STKS
	
	// Allow 1% tolerance
	tolerance := expectedMaxSupply * 0.01
	if totalSupplySTKS < expectedMaxSupply-tolerance || totalSupplySTKS > expectedMaxSupply+tolerance {
		t.Errorf("Total supply is %.0f STKS, expected ~%.0f STKS", 
			totalSupplySTKS, expectedMaxSupply)
		t.Logf("First halving period emits: %.0f STKS", firstHalvingEmissionSTKS)
	} else {
		t.Logf("✓ Max supply: %.0f STKS (first period: %.0f STKS)", totalSupplySTKS, firstHalvingEmissionSTKS)
	}
	
	// Log the reward per block for verification
	rewardSTKS := float64(devnetPreDeflationaryPhaseBaseSubsidy) / float64(constants.SompiPerStokes)
	t.Logf("✓ Reward per block: %.9f STKS", rewardSTKS)
}

// TestTestnetRewardsUnchanged verifies testnet still uses 50 STKS
func TestTestnetRewardsUnchanged(t *testing.T) {
	expectedReward := uint64(50 * constants.SompiPerStokes)
	if defaultPreDeflationaryPhaseBaseSubsidy != expectedReward {
		t.Errorf("Testnet reward should remain 50 STKS (%d sompi), got %d", 
			expectedReward, defaultPreDeflationaryPhaseBaseSubsidy)
	}

	// Verify testnet uses 1 block/second
	expectedBlockTime := uint64(1000) // milliseconds
	actualBlockTime := uint64(defaultTargetTimePerBlock.Milliseconds())
	if actualBlockTime != expectedBlockTime {
		t.Errorf("Testnet block time should be %dms (1 BPS), got %dms", 
			expectedBlockTime, actualBlockTime)
	}
}

// TestDevnetParamsConfiguration verifies DevnetParams uses correct values
func TestDevnetParamsConfiguration(t *testing.T) {
	if DevnetParams.SubsidyGenesisReward != devnetSubsidyGenesisReward {
		t.Errorf("DevnetParams should use devnetSubsidyGenesisReward")
	}
	
	if DevnetParams.PreDeflationaryPhaseBaseSubsidy != devnetPreDeflationaryPhaseBaseSubsidy {
		t.Errorf("DevnetParams should use devnetPreDeflationaryPhaseBaseSubsidy")
	}
	
	if DevnetParams.TargetTimePerBlock != devnetTargetTimePerBlock {
		t.Errorf("DevnetParams should use devnetTargetTimePerBlock")
	}
	
	if DevnetParams.HalvingIntervalDaaScore != devnetHalvingIntervalDaaScore {
		t.Errorf("DevnetParams should use devnetHalvingIntervalDaaScore")
	}
}

// TestTestnetParamsUnchanged verifies TestnetParams still uses default values
func TestTestnetParamsUnchanged(t *testing.T) {
	if TestnetParams.SubsidyGenesisReward != defaultSubsidyGenesisReward {
		t.Errorf("TestnetParams should use defaultSubsidyGenesisReward")
	}
	
	if TestnetParams.PreDeflationaryPhaseBaseSubsidy != defaultPreDeflationaryPhaseBaseSubsidy {
		t.Errorf("TestnetParams should use defaultPreDeflationaryPhaseBaseSubsidy")
	}
	
	if TestnetParams.TargetTimePerBlock != defaultTargetTimePerBlock {
		t.Errorf("TestnetParams should use defaultTargetTimePerBlock")
	}
	
	if TestnetParams.HalvingIntervalDaaScore != defaultHalvingIntervalDaaScore {
		t.Errorf("TestnetParams should use defaultHalvingIntervalDaaScore")
	}
}
