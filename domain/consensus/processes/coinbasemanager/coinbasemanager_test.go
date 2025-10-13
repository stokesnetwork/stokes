package coinbasemanager

import (
	"testing"

	"github.com/Sam-Stokes/stokes/domain/consensus/model/externalapi"
	"github.com/Sam-Stokes/stokes/domain/consensus/utils/constants"
)

// STOKES: Updated test for Bitcoin-style halving
func TestCalcHalvingBlockSubsidy(t *testing.T) {
	// STOKES halving parameters
	const halvingInterval = 126230400 // ~4 years at 1 block/second
	const baseSubsidy = 50 * constants.SompiPerKaspa // 50 STKS
	
	coinbaseManagerInterface := New(
		nil, // databaseContext
		0,   // subsidyGenesisReward
		baseSubsidy, // preDeflationaryPhaseBaseSubsidy
		0,   // coinbasePayloadScriptPublicKeyMaxLength
		&externalapi.DomainHash{}, // genesisHash
		0,   // deflationaryPhaseDaaScore
		0,   // deflationaryPhaseBaseSubsidy
		halvingInterval, // halvingIntervalDaaScore
		nil, // dagTraversalManager
		nil, // ghostdagDataStore
		nil, // acceptanceDataStore
		nil, // daaBlocksStore
		nil, // blockStore
		nil, // pruningStore
		nil) // blockHeaderStore
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	tests := []struct {
		name                 string
		blockDaaScore        uint64
		expectedBlockSubsidy uint64
	}{
		{
			name:                 "genesis/first block",
			blockDaaScore:        0,
			expectedBlockSubsidy: baseSubsidy, // 50 STKS
		},
		{
			name:                 "just before first halving",
			blockDaaScore:        halvingInterval - 1,
			expectedBlockSubsidy: baseSubsidy, // 50 STKS
		},
		{
			name:                 "first halving",
			blockDaaScore:        halvingInterval,
			expectedBlockSubsidy: baseSubsidy / 2, // 25 STKS
		},
		{
			name:                 "second halving",
			blockDaaScore:        halvingInterval * 2,
			expectedBlockSubsidy: baseSubsidy / 4, // 12.5 STKS
		},
		{
			name:                 "third halving",
			blockDaaScore:        halvingInterval * 3,
			expectedBlockSubsidy: baseSubsidy / 8, // 6.25 STKS
		},
		{
			name:                 "after 10 halvings",
			blockDaaScore:        halvingInterval * 10,
			expectedBlockSubsidy: baseSubsidy / 1024,
		},
		{
			name:                 "after 64 halvings (subsidy depleted)",
			blockDaaScore:        halvingInterval * 64,
			expectedBlockSubsidy: 0,
		},
		{
			name:                 "after 100 halvings (subsidy depleted)",
			blockDaaScore:        halvingInterval * 100,
			expectedBlockSubsidy: 0,
		},
	}

	for _, test := range tests {
		blockSubsidy := coinbaseManagerInstance.calcHalvingBlockSubsidy(test.blockDaaScore)
		if blockSubsidy != test.expectedBlockSubsidy {
			t.Errorf("TestCalcHalvingBlockSubsidy: test '%s' failed. Want: %d, got: %d",
				test.name, test.expectedBlockSubsidy, blockSubsidy)
		}
	}
}

// STOKES: Test to verify total supply cap
func TestTotalSupplyCap(t *testing.T) {
	const halvingInterval = 126230400 // ~4 years at 1 block/second
	const baseSubsidy = 50 * constants.SompiPerKaspa // 50 STKS
	
	coinbaseManagerInterface := New(
		nil, // databaseContext
		0,   // subsidyGenesisReward
		baseSubsidy, // preDeflationaryPhaseBaseSubsidy
		0,   // coinbasePayloadScriptPublicKeyMaxLength
		&externalapi.DomainHash{}, // genesisHash
		0,   // deflationaryPhaseDaaScore
		0,   // deflationaryPhaseBaseSubsidy
		halvingInterval, // halvingIntervalDaaScore
		nil, // dagTraversalManager
		nil, // ghostdagDataStore
		nil, // acceptanceDataStore
		nil, // daaBlocksStore
		nil, // blockStore
		nil, // pruningStore
		nil) // blockHeaderStore
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	// Calculate total supply by summing all halvings
	totalSupply := uint64(0)
	for halvingEra := uint64(0); halvingEra < 64; halvingEra++ {
		blockDaaScore := halvingEra * halvingInterval
		subsidy := coinbaseManagerInstance.calcHalvingBlockSubsidy(blockDaaScore)
		
		if subsidy == 0 {
			break
		}
		
		// Add this era's total emission
		eraSupply := subsidy * halvingInterval
		totalSupply += eraSupply
	}
	
	// Expected: ~12.6 billion STKS (actual is 12,623,039,986 due to halving math)
	// This is the exact result of: 50 * 126230400 * (1 + 1/2 + 1/4 + 1/8 + ...)
	expectedSupplySTKS := uint64(12623039986)
	actualSupplySTKS := totalSupply / constants.SompiPerKaspa
	
	if actualSupplySTKS != expectedSupplySTKS {
		t.Errorf("TestTotalSupplyCap: total supply mismatch. Want: %d STKS, got: %d STKS",
			expectedSupplySTKS, actualSupplySTKS)
	}
	
	t.Logf("Total supply verified: %d STKS (~12.6 billion)", actualSupplySTKS)
}
