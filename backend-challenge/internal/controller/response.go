package controller

import (
	"encoding/json"
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/model"
)

// Error Types
const (
	ErrorTypeValidation = "validation"
	ErrorTypeNotFound   = "not_found"
	ErrorTypeUnknown    = "error"
)

// WriteJSON writes any Go value as JSON with the specified HTTP status code.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "failed to encode json response", http.StatusInternalServerError)
	}
}

// WriteJSONError returns an ApiResponse per OpenAPI spec.
func WriteJSONError(w http.ResponseWriter, status int, errType string, msg string) {
	resp := model.ApiResponse{
		Code:    status,
		Type:    errType,
		Message: msg,
	}

	WriteJSON(w, status, resp)
}
