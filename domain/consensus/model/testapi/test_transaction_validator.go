package testapi

import (
	"github.com/Sam-Stokes/stokes/domain/consensus/model"
	"github.com/Sam-Stokes/stokes/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
