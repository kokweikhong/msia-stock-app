package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kokweikhong/msia-stock-app/api-service/internal/api/klsescreener"
	"github.com/kokweikhong/msia-stock-app/api-service/internal/utils"
)

type KLSEScreenerHandler interface {
	GetStockHistoricalData(w http.ResponseWriter, r *http.Request)
	GetKLSEIndexHistoricalData(w http.ResponseWriter, r *http.Request)
	GetQuoteResults(w http.ResponseWriter, r *http.Request)
}

type klseScreenerHandler struct {
	jsonHandler utils.JSONHandler
}

func NewKLSEScreenerHandler() KLSEScreenerHandler {
	jsonHandler := utils.NewJSONHandler()
	return &klseScreenerHandler{
		jsonHandler: jsonHandler,
	}
}

func (h *klseScreenerHandler) GetStockHistoricalData(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	data, err := klsescreener.GetBursaTickerHistoricalData(ticker)
	if err != nil {
		slog.Error("Error getting stock historical data", "error", err)
		h.jsonHandler.ErrorJSON(w, r, err, http.StatusInternalServerError)
		return
	}
	if data == nil {
		h.jsonHandler.WriteJSON(w, r, nil, http.StatusNotFound)
		return
	}
	h.jsonHandler.WriteJSON(w, r, data, http.StatusOK)
}

func (h *klseScreenerHandler) GetKLSEIndexHistoricalData(w http.ResponseWriter, r *http.Request) {
	// year := chi.URLParam(r, "year")
	data, err := klsescreener.GetBursaHistoricalData(klsescreener.OneYear)
	if err != nil {
		slog.Error("Error getting KLSE index historical data", "error", err)
		h.jsonHandler.ErrorJSON(w, r, err, http.StatusInternalServerError)
		return
	}

	if data == nil {
		h.jsonHandler.WriteJSON(w, r, nil, http.StatusNotFound)
		return
	}
	klsescreener.GetQuoteResults()
	h.jsonHandler.WriteJSON(w, r, data, http.StatusOK)
}

func (h *klseScreenerHandler) GetQuoteResults(w http.ResponseWriter, r *http.Request) {
	quotes, err := klsescreener.GetQuoteResults()
	if err != nil {
		slog.Error("Error getting quotes", "error", err)
		h.jsonHandler.ErrorJSON(w, r, err, http.StatusInternalServerError)
		return
	}

	if quotes == nil {
		h.jsonHandler.WriteJSON(w, r, nil, http.StatusNotFound)
		return
	}
	h.jsonHandler.WriteJSON(w, r, quotes, http.StatusOK)
}
