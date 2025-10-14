// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"github.com/kaspanet/go-muhash"
	"github.com/stokesnetwork/stokes/domain/consensus/model/externalapi"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/blockheader"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/subnetworks"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/transactionhelper"
	"math/big"
)

// ============================================================
// MAINNET GENESIS - STOKES
// ============================================================

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, // Subsidy (50 STKS)
	0x00, 0x00, 0x32, 0x00, 0x53, 0x54, 0x4f, 0x4b, // Script version
	0x45, 0x53, 0x20, 0x2d, 0x20, 0x46, 0x61, 0x69, // STOKES - Fair Launch 2025 - Bitcoin-style Halving
	0x72, 0x20, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x20, 0x32, 0x30, 0x32, 0x35, 0x20, 0x2d, 0x20,
	0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2d,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x20, 0x48, 0x61,
	0x6c, 0x76, 0x69, 0x6e, 0x67,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xdc, 0xa2, 0x5f, 0x14, 0xec, 0x25, 0xb3, 0x7e,
	0xfc, 0xb1, 0xba, 0x76, 0x71, 0x54, 0xef, 0x9f,
	0x24, 0x73, 0xe5, 0x96, 0xa5, 0x18, 0xd1, 0xf1,
	0xc0, 0xbe, 0x19, 0xb8, 0x7d, 0x78, 0x69, 0x49,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xcd, 0x59, 0xe0, 0x65, 0x22, 0x13, 0x03, 0xc3,
	0x40, 0x3f, 0xee, 0x8d, 0xae, 0x0a, 0xa0, 0x79,
	0xb9, 0x24, 0xde, 0x3b, 0x01, 0xb9, 0x76, 0xdd,
	0x61, 0x01, 0x1e, 0x81, 0x27, 0xed, 0xd8, 0x21,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1760309945665,
		211244,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

// ============================================================
// DEVNET GENESIS - STOKES
// ============================================================

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, // Subsidy (50 STKS)
	0x00, 0x00, 0x0e, 0x00, 0x73, 0x74, 0x6f, 0x6b, // Script version
	0x65, 0x73, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65, // stokes-devnet
	0x74,
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x90, 0xd1, 0x69, 0xc8, 0x33, 0x6c, 0xa6, 0x20,
	0x88, 0x27, 0x11, 0x41, 0xa5, 0xd2, 0xc6, 0x10,
	0xfd, 0x20, 0xac, 0x91, 0x3f, 0x67, 0x8e, 0x17,
	0x11, 0x54, 0x2b, 0xfd, 0x9a, 0x73, 0x20, 0x58,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x71, 0x70, 0x4e, 0x57, 0x9b, 0x26, 0x9d, 0xd6,
	0x8f, 0xa8, 0x42, 0xcd, 0xca, 0x04, 0x38, 0x33,
	0x19, 0xab, 0x11, 0x6c, 0xd4, 0xbf, 0x17, 0xda,
	0xd7, 0x4a, 0x9b, 0x5d, 0x13, 0x49, 0xa8, 0xd6,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1760309945665,
		298590,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

// ============================================================
// SIMNET GENESIS - STOKES
// ============================================================

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, // Subsidy (50 STKS)
	0x00, 0x00, 0x0e, 0x00, 0x73, 0x74, 0x6f, 0x6b, // Script version
	0x65, 0x73, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, // stokes-simnet
	0x74,
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x3d, 0x7f, 0x17, 0x15, 0xe6, 0xf7, 0xc2, 0x74,
	0x74, 0x30, 0x46, 0x22, 0x6a, 0x37, 0xc1, 0x96,
	0xd8, 0x79, 0xf7, 0x39, 0x1c, 0xdb, 0xcf, 0x8d,
	0x28, 0xef, 0xe6, 0x8e, 0x26, 0x55, 0xc7, 0x79,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xee, 0x34, 0x66, 0x98, 0x62, 0x5b, 0xe2, 0xcc,
	0x7a, 0xd3, 0x58, 0x4a, 0xae, 0x39, 0x13, 0xa6,
	0x42, 0x4c, 0xf6, 0x2b, 0xdd, 0xf0, 0xad, 0xeb,
	0x4e, 0x8c, 0x71, 0xf3, 0x5c, 0x91, 0x91, 0xb8,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1760309945665,
		545259519,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

// ============================================================
// TESTNET GENESIS - STOKES
// ============================================================

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, // Subsidy (50 STKS)
	0x00, 0x00, 0x0f, 0x00, 0x73, 0x74, 0x6f, 0x6b, // Script version
	0x65, 0x73, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, // stokes-testnet
	0x65, 0x74,
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x42, 0x40, 0xb4, 0xdb, 0xce, 0x9f, 0x3a, 0x5d,
	0x14, 0x83, 0xfd, 0x19, 0x14, 0x60, 0x51, 0x26,
	0x45, 0x71, 0x49, 0x47, 0x00, 0xed, 0x38, 0x18,
	0xb4, 0x6c, 0xd8, 0x4b, 0x08, 0x17, 0xef, 0x38,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x62, 0xe9, 0x9d, 0x63, 0xe4, 0xe4, 0xe2, 0x16,
	0xae, 0x84, 0x9b, 0x70, 0x87, 0xa7, 0x54, 0xba,
	0xf5, 0x27, 0x39, 0x92, 0xe2, 0x36, 0x21, 0x90,
	0xa1, 0x6f, 0x0a, 0x5b, 0x9c, 0x3a, 0x57, 0x97,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1760309945665,
		511705087,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
