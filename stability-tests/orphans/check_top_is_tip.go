package main

import (
	"github.com/Sam-Stokes/stokes/domain/consensus/model/externalapi"
	"github.com/Sam-Stokes/stokes/domain/consensus/utils/consensushashing"
	"github.com/Sam-Stokes/stokes/stability-tests/common/rpc"
	"github.com/pkg/errors"
)

func checkTopBlockIsTip(rpcClient *rpc.Client, topBlock *externalapi.DomainBlock) error {
	selectedTipHashResponse, err := rpcClient.GetSelectedTipHash()
	if err != nil {
		return err
	}

	topBlockHash := consensushashing.BlockHash(topBlock)
	if selectedTipHashResponse.SelectedTipHash != topBlockHash.String() {
		return errors.Errorf("selectedTipHash is '%s' while expected to be topBlock's hash `%s`",
			selectedTipHashResponse.SelectedTipHash, topBlockHash)
	}

	return nil
}
