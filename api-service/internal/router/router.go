package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kokweikhong/msia-stock-app/api-service/internal/handler"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	klseScreenerHandler := handler.NewKLSEScreenerHandler()

	r.Route("/klsescreener", func(r chi.Router) {
		r.Get("/stock/{ticker}", klseScreenerHandler.GetStockHistoricalData)
		r.Get("/index", klseScreenerHandler.GetKLSEIndexHistoricalData)
		r.Get("/quotes", klseScreenerHandler.GetQuoteResults)
	})

	return r
}
