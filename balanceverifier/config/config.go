package config

import (
	"net/url"
	"math/big"
)

type BalanceVerifierConfig struct {
	Url *url.URL
	ChainID *big.Int
}