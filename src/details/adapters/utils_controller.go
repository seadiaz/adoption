package adapters

import (
	"encoding/json"
	"net/http"
)

func replyWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	errResponse := &ErrorResponse{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(errResponse)
}
