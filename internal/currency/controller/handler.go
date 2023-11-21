package controller

import (
	"currency-converter-rev2/internal/currency/service"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	currencyServ service.IService
}

func NewHandler(currencyServ service.IService) *Handler {
	return &Handler{
		currencyServ: currencyServ,
	}
}

func (c Handler) InitRoutes() *chi.Mux {
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
	router.Get("/currencies", c.GetAllCurrencies)
	router.Post("/currency", c.CreateCurrency)
	router.Put("/currency/{id}", c.UpdateCurrency)
	router.Delete("/currency/{id}", c.DeleteCurrency)

	//ExchangeRates
	router.Post("/exchange-rate", c.CreateExchageRate)
	router.Get("/exchange-rates", c.GetAllExchageRates)
	router.Get("/exchange-rate/{code}", c.GetExchangeRateByCode)
	//exchangeRouter.Get("/exchangeRates/exchangeRate/exchange{code}", controllers.ExchangeAmount)
	return router
}

func (c *Handler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var currencyCreateView CurrencyCreateView
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Bad request",
		})
		return
	}

	err = json.Unmarshal(body, &currencyCreateView)
	if err != nil {
		log.Printf("Error decoding JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Bad Request",
		})
		return
	}

	err = c.currencyServ.CreateCurrency(r.Context(), currencyCreateView.MapToEntity())
	if err != nil {
		log.Printf("Error creating currency: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Currency created successfully",
	})
}

func (c *Handler) GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currencies, err := c.currencyServ.GetAllCurrencies(r.Context())
	if err != nil {
		log.Printf("Error getting currencies: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}
	var currencyView CurrencyView
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"status":     true,
		"currencies": currencyView.MapToViewList(currencies),
	})
}

func (c *Handler) UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var currencyCreateView CurrencyCreateView
	id := chi.URLParam(r, "id")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Bad request",
		})
		return
	}

	err = json.Unmarshal(body, &currencyCreateView)
	if err != nil {
		log.Printf("Error decoding JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Bad Request",
		})
		return
	}
	err = c.currencyServ.UpdateCurrency(r.Context(), currencyCreateView.MapToEntity(), id)
	if err != nil {
		log.Printf("Error updating currency: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})

	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Currency updated successfully",
	})
}

func (c *Handler) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	err := c.currencyServ.DeleteCurrencyById(r.Context(), id)
	if err != nil {
		log.Printf("Error deleting currency: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})

	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Currency deleted successfully",
	})
}
func (c *Handler) CreateExchageRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var exchageRateCreateView CreateExchangeRateView
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Bad request",
		})
		return
	}

	err = json.Unmarshal(body, &exchageRateCreateView)
	if err != nil {
		log.Printf("Error decoding JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Bad Request",
		})
		return
	}

	err = c.currencyServ.CreateExchangeRate(r.Context(), exchageRateCreateView.MapToEntity())
	if err != nil {
		log.Printf("Error creating exchange rate: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Exchange rate created successfully",
	})
}
func (c *Handler) GetAllExchageRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eRates, err := c.currencyServ.GetAllExchangeRates(r.Context())
	if err != nil {
		log.Printf("Error getting exchange rates: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
		return
	}
	var eRview ExchangeRateView
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"status":        true,
		"exchangeRates": eRview.MapToViewList(eRates),
	})
}

func (c *Handler) GetExchangeRateByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	eRate, err := c.currencyServ.GetExchangeRateByCode(r.Context(), code)
	if err != nil {
		log.Printf("Error getting exchange rate: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
	}
	var eRview ExchangeRateView

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"status":       true,
		"exchangeRate": eRview.MapToView(eRate),
	})
}
