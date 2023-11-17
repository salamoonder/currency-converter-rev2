package controller

import (
	"currency-converter-rev2/internal/currency/entity"
	"currency-converter-rev2/internal/currency/service"
	"currency-converter-rev2/pkg/helpers"
	"encoding/json"
	"log"
	"net/http"
)

type CurrencyHandler struct {
	currencyServ service.IService
}

func NewHandler(currencyServ service.IService) *CurrencyHandler {
	return &CurrencyHandler{
		currencyServ: currencyServ,
	}
}
func (c *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	var currencyData entity.Currency
	err := json.NewDecoder(r.Body).Decode(&currencyData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error decoding JSON: %v\n", err)
		return
	}

	res, err := c.currencyServ.CreateCurrency(r.Context(), currencyData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error creating currency: %v\n", err)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, &res)
}

func (c *CurrencyHandler) GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := c.currencyServ.GetAllCurrencies(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting currencies: %v\n", err)
		return
	}

	jsonResponse, err := json.Marshal(currencies)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (c *CurrencyHandler) UpdateCurrencyByCode(w http.ResponseWriter, r *http.Request) {
	var currencyData entity.Currency
	err := json.NewDecoder(r.Body).Decode(&currencyData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error decoding JSON: %v\n", err)
		return
	}

	res, err := c.currencyServ.UpdateCurrency(r.Context(), currencyData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error updating currency: %v\n", err)
		return
	}
	helpers.WriteJSON(w, http.StatusCreated, &res)
}
