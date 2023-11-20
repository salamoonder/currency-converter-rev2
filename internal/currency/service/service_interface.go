package service

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
)

type IService interface {
	GetAllCurrencies(ctx context.Context) ([]entity.Currency, error)
	GetCurrencyById(ctx context.Context, id string) (entity.Currency, error)
	CreateCurrency(ctx context.Context, curr entity.Currency) error
	UpdateCurrency(ctx context.Context, curr entity.Currency) error
	//DeleteCurrency(ctx context.Context, id string) error
	//CreateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error
	//UpdateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error
}
type ICurrencyRepository interface {
	CreateCurrency(ctx context.Context, curr entity.Currency) error
	UpdateCurrency(ctx context.Context, curr entity.Currency) error
	GetAllCurrencies(ctx context.Context) ([]entity.Currency, error)
	GetCurrencyById(ctx context.Context, id string) (entity.Currency, error)
}

type IExchangeRateRepository interface {
	CreateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error
	UpdateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error
	//GetAllExchangeRates(ctx context.Context) ([]*entity.ExchangeRate, error)
	//GetExchangeRateById(ctx context.Context, id string) (entity.ExchangeRate, error)
}
