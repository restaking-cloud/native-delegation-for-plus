package config

import (
	"net/url"
	"math/big"
)

type SignatureSwapperConfig struct {
	Url *url.URL
	ChainID *big.Int
	Domain uint64
}