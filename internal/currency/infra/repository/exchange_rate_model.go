package repository

import (
	"currency-converter-rev2/internal/currency/entity"
	"time"
)

type ExchangeRateData struct {
	ID              string       `json:"id"`
	BaseCurrency    CurrencyData `json:"baseCurrency"`
	TargetCurrency  CurrencyData `json:"targetCurrency"`
	Rate            float64      `json:"rate"`
	Amount          float64      `json:"amount,omitempty"`
	ConvertedAmount float64      `json:"convertedAmount,omitempty"`
	CreatedAt       time.Time    `json:"createdAt"`
	UpdatedAt       time.Time    `json:"updatedAt,omitempty"`
	DeletedAt       time.Time    `json:"deletedAt,omitempty"`
}
type ShortExchangeRateData struct {
	BaseCurrency   string  `json:"baseCurrency"`
	TargetCurrency string  `json:"targetCurrency"`
	Rate           float64 `json:"rate"`
}

func (e *ShortExchangeRateData) MapToRepo(rate entity.ShortExchangeRate) {
	e.BaseCurrency = rate.BaseCurrency
	e.TargetCurrency = rate.TargetCurrency
	e.Rate = rate.Rate
}
func (e *ShortExchangeRateData) MapToEntity() entity.ShortExchangeRate {
	var rate entity.ShortExchangeRate
	rate.BaseCurrency = e.BaseCurrency
	rate.TargetCurrency = e.TargetCurrency
	rate.Rate = e.Rate
	return rate
}

func (e *ExchangeRateData) MapToRepo(exchangeRate entity.ExchangeRate) {
	e.ID = exchangeRate.ID
	e.BaseCurrency.MapToRepo(exchangeRate.BaseCurrency)
	e.TargetCurrency.MapToRepo(exchangeRate.TargetCurrency)
	e.Rate = exchangeRate.Rate
	e.Amount = exchangeRate.Amount
	e.ConvertedAmount = exchangeRate.ConvertedAmount
	e.CreatedAt = exchangeRate.CreatedAt
	e.UpdatedAt = exchangeRate.UpdatedAt
	e.DeletedAt = exchangeRate.DeletedAt
}

func (e *ExchangeRateData) MapToEntity() entity.ExchangeRate {
	var rate entity.ExchangeRate
	rate.ID = e.ID
	rate.BaseCurrency = e.BaseCurrency.MapToEntity()
	rate.TargetCurrency = e.TargetCurrency.MapToEntity()
	rate.Rate = e.Rate
	rate.Amount = e.Amount
	rate.ConvertedAmount = e.ConvertedAmount
	rate.CreatedAt = e.CreatedAt
	rate.UpdatedAt = e.UpdatedAt
	rate.DeletedAt = e.DeletedAt
	return rate
}
