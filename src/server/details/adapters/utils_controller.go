package adapters

import (
	"encoding/json"
	"net/http"
)

func replyWithError(w http.ResponseWriter, statusCode int, err error) {
	errResponse := &ErrorResponse{
		Message: err.Error(),
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResponse)
}

func replyJSONResponse(w http.ResponseWriter, output interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
