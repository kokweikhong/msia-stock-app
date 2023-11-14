package utils

import (
	"encoding/json"
	"net/http"

	"log/slog"
)

type JSONHandler interface {
	WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int)
	ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error
	ErrorJSON(w http.ResponseWriter, r *http.Request, err error, statusCode int)
}

type jsonHandler struct {
}

func NewJSONHandler() JSONHandler {
	return &jsonHandler{}
}

func (h *jsonHandler) WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Error writing JSON response", "error", err)
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func (h *jsonHandler) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		slog.Error("Error reading JSON request body", "error", err)
		http.Error(w, "Error reading JSON request body", http.StatusBadRequest)
		return err
	}
	return nil
}

func (h *jsonHandler) ErrorJSON(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	type jsonError struct {
		Error string `json:"error"`
	}
	h.WriteJSON(w, r, jsonError{Error: err.Error()}, statusCode)
}
