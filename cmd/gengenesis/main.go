package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/stokesnetwork/stokes/domain/consensus/model/externalapi"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/blockheader"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/consensushashing"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/subnetworks"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/transactionhelper"
	"github.com/kaspanet/go-muhash"
)

// STOKES Genesis Block Generator
// This tool generates new genesis blocks for all STOKES networks

func main() {
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                STOKES Genesis Block Generator                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Generate genesis blocks for all networks
	generateMainnetGenesis()
	fmt.Println()
	generateTestnetGenesis()
	fmt.Println()
	generateSimnetGenesis()
	fmt.Println()
	generateDevnetGenesis()
	fmt.Println()

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║  Copy the output above into domain/dagconfig/genesis.go      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
}

func generateMainnetGenesis() {
	fmt.Println("// ============================================================")
	fmt.Println("// MAINNET GENESIS")
	fmt.Println("// ============================================================")
	fmt.Println()

	// Mainnet genesis message: "STOKES - Fair Launch 2025 - Bitcoin-style Halving"
	genesisMessage := "STOKES - Fair Launch 2025 - Bitcoin-style Halving"
	
	// Current timestamp (you can customize this)
	timestamp := time.Now().UnixMilli()
	
	// Initial difficulty (same as Kaspa mainnet for now)
	bits := uint32(0x3392c)
	
	// Nonce (will need to be mined for real mainnet)
	nonce := uint64(0)

	genesis := createGenesisBlock(genesisMessage, timestamp, bits, nonce, 50*100_000_000)
	printGenesisBlock("mainnet", genesis, genesisMessage)
}

func generateTestnetGenesis() {
	fmt.Println("// ============================================================")
	fmt.Println("// TESTNET GENESIS")
	fmt.Println("// ============================================================")
	fmt.Println()

	genesisMessage := "stokes-testnet"
	timestamp := time.Now().UnixMilli()
	bits := uint32(0x1e7fffff) // Easier difficulty for testnet
	nonce := uint64(0)

	genesis := createGenesisBlock(genesisMessage, timestamp, bits, nonce, 50*100_000_000)
	printGenesisBlock("testnet", genesis, genesisMessage)
}

func generateSimnetGenesis() {
	fmt.Println("// ============================================================")
	fmt.Println("// SIMNET GENESIS")
	fmt.Println("// ============================================================")
	fmt.Println()

	genesisMessage := "stokes-simnet"
	timestamp := time.Now().UnixMilli()
	bits := uint32(0x207fffff) // Very easy difficulty for simnet
	nonce := uint64(0)

	genesis := createGenesisBlock(genesisMessage, timestamp, bits, nonce, 50*100_000_000)
	printGenesisBlock("simnet", genesis, genesisMessage)
}

func generateDevnetGenesis() {
	fmt.Println("// ============================================================")
	fmt.Println("// DEVNET GENESIS")
	fmt.Println("// ============================================================")
	fmt.Println()

	genesisMessage := "stokes-devnet"
	timestamp := time.Now().UnixMilli()
	bits := uint32(0x48e5e) // Medium difficulty for devnet
	nonce := uint64(0)

	genesis := createGenesisBlock(genesisMessage, timestamp, bits, nonce, 50*100_000_000)
	printGenesisBlock("devnet", genesis, genesisMessage)
}

func createGenesisBlock(message string, timestamp int64, bits uint32, nonce uint64, subsidy uint64) *externalapi.DomainBlock {
	// Create coinbase payload with genesis message
	payload := createCoinbasePayload(message, subsidy)

	// Create coinbase transaction
	coinbaseTx := transactionhelper.NewSubnetworkTransaction(
		0,
		[]*externalapi.DomainTransactionInput{},
		[]*externalapi.DomainTransactionOutput{}, // No outputs in genesis
		&subnetworks.SubnetworkIDCoinbase,
		0,
		payload,
	)

	// Calculate merkle root
	merkleRoot := calculateMerkleRoot([]*externalapi.DomainTransaction{coinbaseTx})

	// Create genesis block header
	header := blockheader.NewImmutableBlockHeader(
		0,                                  // Version
		[]externalapi.BlockLevelParents{},  // No parents
		merkleRoot,                         // Merkle root
		&externalapi.DomainHash{},          // Accepted ID merkle root (empty)
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()), // UTXO commitment
		timestamp,                          // Timestamp
		bits,                               // Difficulty bits
		nonce,                              // Nonce
		0,                                  // DAA score
		0,                                  // Blue score
		big.NewInt(0),                      // Blue work
		&externalapi.DomainHash{},          // Pruning point
	)

	// Create genesis block
	genesisBlock := &externalapi.DomainBlock{
		Header:       header,
		Transactions: []*externalapi.DomainTransaction{coinbaseTx},
	}

	return genesisBlock
}

func createCoinbasePayload(message string, subsidy uint64) []byte {
	payload := []byte{}

	// Blue score (8 bytes) - 0 for genesis
	payload = append(payload, 0, 0, 0, 0, 0, 0, 0, 0)

	// Subsidy (8 bytes) - 50 STKS in sompi (little endian)
	subsidyBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		subsidyBytes[i] = byte(subsidy >> (i * 8))
	}
	payload = append(payload, subsidyBytes...)

	// Script version (2 bytes)
	payload = append(payload, 0, 0)

	// Varint for script length
	payload = append(payload, byte(len(message)+1))

	// OP_FALSE
	payload = append(payload, 0x00)

	// Genesis message
	payload = append(payload, []byte(message)...)

	return payload
}

func calculateMerkleRoot(transactions []*externalapi.DomainTransaction) *externalapi.DomainHash {
	if len(transactions) == 0 {
		return &externalapi.DomainHash{}
	}

	// For a single transaction (genesis), the merkle root is just the transaction hash
	txHash := consensushashing.TransactionHash(transactions[0])
	return txHash
}

func printGenesisBlock(network string, genesis *externalapi.DomainBlock, message string) {
	// Calculate genesis hash
	genesisHash := consensushashing.HeaderHash(genesis.Header)
	merkleRoot := genesis.Header.HashMerkleRoot()

	fmt.Printf("var %sGenesisTxOuts = []*externalapi.DomainTransactionOutput{}\n\n", network)

	// Print payload
	payload := genesis.Transactions[0].Payload
	fmt.Printf("var %sGenesisTxPayload = []byte{\n", network)
	for i := 0; i < len(payload); i += 8 {
		fmt.Print("\t")
		end := i + 8
		if end > len(payload) {
			end = len(payload)
		}
		for j := i; j < end; j++ {
			fmt.Printf("0x%02x, ", payload[j])
		}
		
		// Add comment for first few lines
		if i == 0 {
			fmt.Print("// Blue score")
		} else if i == 8 {
			fmt.Print("// Subsidy (50 STKS)")
		} else if i == 16 {
			fmt.Print("// Script version")
		} else if i == 24 {
			fmt.Printf("// %s", message)
		}
		fmt.Println()
	}
	fmt.Println("}")
	fmt.Println()

	// Print coinbase transaction
	fmt.Printf("var %sGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,\n", network)
	fmt.Printf("\t[]*externalapi.DomainTransactionInput{}, %sGenesisTxOuts,\n", network)
	fmt.Printf("\t&subnetworks.SubnetworkIDCoinbase, 0, %sGenesisTxPayload)\n", network)
	fmt.Println()

	// Print genesis hash
	fmt.Printf("var %sGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{\n", network)
	printHashBytes(genesisHash)
	fmt.Println("})")
	fmt.Println()

	// Print merkle root
	fmt.Printf("var %sGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{\n", network)
	printHashBytes(merkleRoot)
	fmt.Println("})")
	fmt.Println()

	// Print genesis block
	fmt.Printf("var %sGenesisBlock = externalapi.DomainBlock{\n", network)
	fmt.Println("\tHeader: blockheader.NewImmutableBlockHeader(")
	fmt.Println("\t\t0,")
	fmt.Println("\t\t[]externalapi.BlockLevelParents{},")
	fmt.Printf("\t\t%sGenesisMerkleRoot,\n", network)
	fmt.Println("\t\t&externalapi.DomainHash{},")
	fmt.Println("\t\texternalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),")
	fmt.Printf("\t\t%d,\n", genesis.Header.TimeInMilliseconds())
	fmt.Printf("\t\t%d,\n", genesis.Header.Bits())
	fmt.Printf("\t\t%d,\n", genesis.Header.Nonce())
	fmt.Println("\t\t0,")
	fmt.Println("\t\t0,")
	fmt.Println("\t\tbig.NewInt(0),")
	fmt.Println("\t\t&externalapi.DomainHash{},")
	fmt.Println("\t),")
	fmt.Printf("\tTransactions: []*externalapi.DomainTransaction{%sGenesisCoinbaseTx},\n", network)
	fmt.Println("}")
	fmt.Println()

	// Print summary
	fmt.Printf("// Genesis Hash: %s\n", genesisHash)
	fmt.Printf("// Merkle Root:  %s\n", merkleRoot)
	fmt.Printf("// Timestamp:    %d (%s)\n", genesis.Header.TimeInMilliseconds(), time.UnixMilli(genesis.Header.TimeInMilliseconds()).Format(time.RFC3339))
	fmt.Printf("// Message:      %s\n", message)
}

func printHashBytes(hash *externalapi.DomainHash) {
	hashBytes := hash.ByteSlice()
	for i := 0; i < len(hashBytes); i += 8 {
		fmt.Print("\t")
		end := i + 8
		if end > len(hashBytes) {
			end = len(hashBytes)
		}
		for j := i; j < end; j++ {
			fmt.Printf("0x%02x, ", hashBytes[j])
		}
		fmt.Println()
	}
}

func init() {
	// Print hex encoding helper
	_ = hex.EncodeToString
}
