package controller

import (
	"currency-converter-rev2/internal/currency/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func Routes(serv service.IService) http.Handler {
	h := NewHandler(serv)
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//Currencies
	router.Get("/currencies", h.GetAllCurrencies)
	router.Post("/currency", h.CreateCurrency)
	router.Put("/currency/{code}", h.UpdateCurrencyByCode)
	//currencyroute.Delete("/currencies/currency/{id}", controllers.DeleteCurrency)
	//
	////ExchangeRates
	//exchangeRouter := chi.NewRouter()
	//exchangeRouter.Get("/exchangeRates", controllers.GetAllExchangeRates)
	//exchangeRouter.Get("/exchangeRates/exchangeRate/{code}", controllers.GetExchangeRateByCode)
	//exchangeRouter.Get("/exchangeRates/exchangeRate/exchange{code}", controllers.ExchangeAmount)
	return router
}
