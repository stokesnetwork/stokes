package rpchandlers

import (
	"encoding/json"
	"github.com/stokesnetwork/stokes/app/appmessage"
	"github.com/stokesnetwork/stokes/app/protocol/protocolerrors"
	"github.com/stokesnetwork/stokes/app/rpc/rpccontext"
	"github.com/stokesnetwork/stokes/domain/consensus/ruleerrors"
	"github.com/stokesnetwork/stokes/domain/consensus/utils/consensushashing"
	"github.com/stokesnetwork/stokes/infrastructure/network/netadapter/router"
	"github.com/pkg/errors"
)

// HandleSubmitBlock handles the respectively named RPC command
func HandleSubmitBlock(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	submitBlockRequest := request.(*appmessage.SubmitBlockRequestMessage)

	var err error
	isSynced := false
	// STOKES: Allow solo mining - if no peers, consider synced for solo testnet
	hasPeers := context.ProtocolManager.Context().HasPeers()
	if hasPeers {
		// The node is considered synced if it has peers and consensus state is nearly synced
		isSynced, err = context.ProtocolManager.Context().IsNearlySynced()
		if err != nil {
			return nil, err
		}
	} else {
		// No peers = solo mode, always accept blocks
		isSynced = true
	}

	if !context.Config.AllowSubmitBlockWhenNotSynced && !isSynced {
		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Block not submitted - node is not synced"),
			RejectReason: appmessage.RejectReasonIsInIBD,
		}, nil
	}

	domainBlock, err := appmessage.RPCBlockToDomainBlock(submitBlockRequest.Block)
	if err != nil {
		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Could not parse block: %s", err),
			RejectReason: appmessage.RejectReasonBlockInvalid,
		}, nil
	}

	if !submitBlockRequest.AllowNonDAABlocks {
		virtualDAAScore, err := context.Domain.Consensus().GetVirtualDAAScore()
		if err != nil {
			return nil, err
		}
		// A simple heuristic check which signals that the mined block is out of date
		// and should not be accepted unless user explicitly requests
		daaWindowSize := uint64(context.Config.NetParams().DifficultyAdjustmentWindowSize)
		if virtualDAAScore > daaWindowSize && domainBlock.Header.DAAScore() < virtualDAAScore-daaWindowSize {
			return &appmessage.SubmitBlockResponseMessage{
				Error: appmessage.RPCErrorf("Block rejected. Reason: block DAA score %d is too far "+
					"behind virtual's DAA score %d", domainBlock.Header.DAAScore(), virtualDAAScore),
				RejectReason: appmessage.RejectReasonBlockInvalid,
			}, nil
		}
	}

	err = context.ProtocolManager.AddBlock(domainBlock)
	if err != nil {
		isProtocolOrRuleError := errors.As(err, &ruleerrors.RuleError{}) || errors.As(err, &protocolerrors.ProtocolError{})
		if !isProtocolOrRuleError {
			return nil, err
		}

		jsonBytes, _ := json.MarshalIndent(submitBlockRequest.Block.Header, "", "    ")
		if jsonBytes != nil {
			log.Warnf("The RPC submitted block triggered a rule/protocol error (%s), printing "+
				"the full header for debug purposes: \n%s", err, string(jsonBytes))
		}

		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Block rejected. Reason: %s", err),
			RejectReason: appmessage.RejectReasonBlockInvalid,
		}, nil
	}

	log.Infof("Accepted block %s via submitBlock", consensushashing.BlockHash(domainBlock))

	response := appmessage.NewSubmitBlockResponseMessage()
	return response, nil
}
