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
	router.Put("/currency/{id}", c.UpdateCurrencyById)
	router.Delete("/currencies/currency/{id}", c.DeleteCurrency)

	////ExchangeRates
	//exchangeRouter := chi.NewRouter()
	//exchangeRouter.Get("/exchangeRates", controllers.GetAllExchangeRates)
	//exchangeRouter.Get("/exchangeRates/exchangeRate/{code}", controllers.GetExchangeRateByCode)
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

func (c *Handler) UpdateCurrencyById(w http.ResponseWriter, r *http.Request) {
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
