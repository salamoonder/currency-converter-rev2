package service

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
)

type service struct {
	currency     ICurrencyRepository
	exchangeRate IExchangeRateRepository
}

func NewService(currencyRepository ICurrencyRepository, exchangeRateRepo IExchangeRateRepository) *service {
	return &service{
		currency:     currencyRepository,
		exchangeRate: exchangeRateRepo,
	}
}

func (s *service) CreateCurrency(ctx context.Context, curr entity.Currency) error {
	return s.currency.CreateCurrency(ctx, curr)
}
func (s *service) UpdateCurrency(ctx context.Context, curr entity.Currency, id string) error {
	return s.currency.UpdateCurrency(ctx, curr, id)
}
func (s *service) CreateExchangeRate(ctx context.Context, er entity.ShortExchangeRate) error {
	return s.exchangeRate.CreateExchangeRate(ctx, er)
}
func (s *service) UpdateExchangeRate(ctx context.Context, er entity.ExchangeRate) error {
	return nil
}
func (s *service) GetAllCurrencies(ctx context.Context) ([]entity.Currency, error) {
	return s.currency.GetAllCurrencies(ctx)
}
func (s *service) GetCurrencyById(ctx context.Context, id string) (entity.Currency, error) {
	return s.currency.GetCurrencyById(ctx, id)
}
func (s *service) DeleteCurrencyById(ctx context.Context, id string) error {
	return s.currency.DeleteCurrencyById(ctx, id)
}
func (s *service) GetAllExchangeRates(ctx context.Context) ([]entity.ExchangeRate, error) {
	return s.exchangeRate.GetAllExchangeRates(ctx)
}
func (s *service) GetExchangeRateByCode(ctx context.Context, code string) (entity.ExchangeRate, error) {
	return s.exchangeRate.GetExchangeRateByCode(ctx, code)
}
