package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/kaspanet/go-muhash"
	"github.com/stokesnetwork/stokes/domain/consensus/model/externalapi"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/blockheader"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/consensushashing"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/subnetworks"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/transactionhelper"
)

func main() {
	fmt.Println("Generating Devnet Genesis Block Hashes")
	fmt.Println("========================================")
	fmt.Println()

	// Devnet genesis transaction outputs (empty)
	devnetGenesisTxOuts := []*externalapi.DomainTransactionOutput{}

	// Devnet genesis transaction payload
	// Subsidy: 0.007922022 STKS = 792,202 sompi = 0xC168A in hex (little-endian: 0x8A, 0x16, 0x0C)
	devnetGenesisTxPayload := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
		0x8A, 0x16, 0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, // Subsidy (0.007922022 STKS = 792,202 sompi)
		0x00, 0x00, 0x0e, 0x00, 0x73, 0x74, 0x6f, 0x6b, // Script version
		0x65, 0x73, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65, // stokes-devnet
		0x74,
	}

	// Create the coinbase transaction
	devnetGenesisCoinbaseTx := transactionhelper.NewSubnetworkTransaction(0,
		[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
		&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

	// Calculate the merkle root (hash of the coinbase transaction)
	devnetGenesisMerkleRoot := consensushashing.TransactionHash(devnetGenesisCoinbaseTx)

	// Create the genesis block header
	// STOKES: Using same bits as testnet (511705087) for easier local mining
	header := blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1760309945665,
		511705087,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	)

	// Create the genesis block
	devnetGenesisBlock := externalapi.DomainBlock{
		Header:       header,
		Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
	}

	// Calculate the genesis block hash
	devnetGenesisHash := consensushashing.BlockHash(&devnetGenesisBlock)

	// Print the results
	fmt.Println("Devnet Genesis Merkle Root:")
	fmt.Printf("var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{\n")
	printHashAsGoArray(devnetGenesisMerkleRoot)
	fmt.Println("})")
	fmt.Println()

	fmt.Println("Devnet Genesis Hash:")
	fmt.Printf("var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{\n")
	printHashAsGoArray(devnetGenesisHash)
	fmt.Println("})")
	fmt.Println()

	fmt.Println("Verification:")
	fmt.Printf("Merkle Root: %s\n", devnetGenesisMerkleRoot)
	fmt.Printf("Genesis Hash: %s\n", devnetGenesisHash)
}

func printHashAsGoArray(hash *externalapi.DomainHash) {
	bytes := hash.ByteSlice()
	for i := 0; i < len(bytes); i += 8 {
		fmt.Print("\t")
		for j := 0; j < 8 && i+j < len(bytes); j++ {
			fmt.Printf("0x%02x, ", bytes[i+j])
		}
		fmt.Println()
	}
}

func printHashAsHex(hash *externalapi.DomainHash) {
	fmt.Println(hex.EncodeToString(hash.ByteSlice()))
}
