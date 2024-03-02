package config

import (
	"net/url"
	"math/big"
)

type SubgraphConfig struct {
	Url *url.URL
	ChainID *big.Int
}