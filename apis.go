package k2

import (
	"encoding/json"
	"net/http"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

const (
	// Router paths
	pathRoot     = "/"
	pathExit     = "/eth/v1/exit"
	pathClaim    = "/eth/v1/claim"
	pathRegister = "/eth/v1/register"
)

func (k2 *K2Service) handleRoot(w http.ResponseWriter, _ *http.Request) {
	k2.respondOK(w, "K2 module is running")
}

func (k2 *K2Service) handleExit(w http.ResponseWriter, r *http.Request) {
	// Post call.
	// Handles the removal of the validators delegated balance from the K2 contract,
	// and therefore exits the K2 protocol.

	payload := phase0.BLSPubKey{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		k2.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := k2.processExit(payload)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	k2.respondOK(w, result)

}

func (k2 *K2Service) handleClaim(w http.ResponseWriter, r *http.Request) {
	// Post call.
	// Handles the claim of rewards for the validators in the K2 contract.

	payload := []phase0.BLSPubKey{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		k2.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := k2.batchProcessClaims(payload)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	k2.respondOK(w, result)

}

func (k2 *K2Service) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Post call.
	// Handles the native delegation of validators in the K2 contract if configured to do so.
	// This performs the registration of validators if required, in the Proposer Registry,
	// and then natively delegates them to the K2 contract.
	// If the module is not configured to perform native delegation, it will only register
	// the validators in the Proposer Registry.

	payload := []apiv1.SignedValidatorRegistration{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		k2.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := k2.batchProcessValidatorRegistrations(payload)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	k2.respondOK(w, result)
}
