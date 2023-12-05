package common

import (
	"net/url"
	"strings"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

func CreateUrl(urlString string) (*url.URL, error) {
	if urlString == "" {
		return nil, nil
	}
	if !strings.HasPrefix(urlString, "http") {
		urlString = "http://" + urlString
	}

	return url.ParseRequestURI(urlString)
}

func GetListOfBLSKeysFromSignedValidatorRegistration(payload []apiv1.SignedValidatorRegistration) (pubkeys []phase0.BLSPubKey, payloadMap map[string]apiv1.SignedValidatorRegistration) {
	payloadMap = make(map[string]apiv1.SignedValidatorRegistration)
	for _, reg := range payload {
		pubkeys = append(pubkeys, reg.Message.Pubkey)
		payloadMap[reg.Message.Pubkey.String()] = reg
	}
	return pubkeys, payloadMap
}