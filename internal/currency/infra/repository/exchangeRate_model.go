package repository

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExchangeRateData struct {
	ID              string       `json:"id"`
	BaseCurrency    CurrencyData `json:"baseCurrency"`
	TargetCurrency  CurrencyData `json:"targetCurrency"`
	Rate            float64      `json:"rate"`
	Amount          float64      `json:"amount,omitempty"`
	ConvertedAmount float64      `json:"convertedAmount,omitempty"`
}

type exchangeRateRepo struct {
	db *pgxpool.Pool
}

func NewExchangeRateRepo(db *pgxpool.Pool) *exchangeRateRepo {
	return &exchangeRateRepo{
		db: db,
	}
}

func (e *ExchangeRateData) MapToRepo(exchangeRate entity.ExchangeRate) {
	e.ID = exchangeRate.ID
	e.BaseCurrency.MapToRepo(exchangeRate.BaseCurrency)
	e.TargetCurrency.MapToRepo(exchangeRate.TargetCurrency)
	e.Rate = exchangeRate.Rate
	e.Amount = exchangeRate.Amount
	e.ConvertedAmount = exchangeRate.ConvertedAmount
}

func (e *ExchangeRateData) MapToEntity() entity.ExchangeRate {
	var rate entity.ExchangeRate
	rate.ID = e.ID
	rate.BaseCurrency = e.BaseCurrency.MapToEntity()
	rate.TargetCurrency = e.TargetCurrency.MapToEntity()
	rate.Rate = e.Rate
	rate.Amount = e.Amount
	rate.ConvertedAmount = e.ConvertedAmount
	return rate
}

func (r exchangeRateRepo) CreateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
	exchRepo := ExchangeRateData{}
	exchRepo.MapToRepo(exch)
	return nil
}

func (r exchangeRateRepo) UpdateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
	exchRepo := ExchangeRateData{}
	exchRepo.MapToRepo(exch)
	return nil
}
