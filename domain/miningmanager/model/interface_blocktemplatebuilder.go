package model

import (
	consensusexternalapi "github.com/Sam-Stokes/stokes/domain/consensus/model/externalapi"
)

// BlockTemplateBuilder builds block templates for miners to consume
type BlockTemplateBuilder interface {
	BuildBlockTemplate(coinbaseData *consensusexternalapi.DomainCoinbaseData) (*consensusexternalapi.DomainBlockTemplate, error)
	ModifyBlockTemplate(newCoinbaseData *consensusexternalapi.DomainCoinbaseData,
		blockTemplateToModify *consensusexternalapi.DomainBlockTemplate) (*consensusexternalapi.DomainBlockTemplate, error)
}
