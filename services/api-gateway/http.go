package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	// TODO: call trip service

	response := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusOK, response)
}
