package subgraph

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type MetaQuery struct {
	Meta struct {
		HasIndexingErrors bool
		Block             struct {
			Number int
			Hash   string
		}
	} `graphql:"_meta"`
}

type NodeRunnerByIdsQuery struct {
	NodeRunners []struct {
		Id            common.Address
		BlsPublicKeys []struct {
			Id phase0.BLSPubKey
		} `graphql:"blsPublicKeys (skip: $skip_validators, first: $limit_validators)"`
	} `graphql:"nodeRunners(where: {id_in: $ids}, first: 1000)"`
}
