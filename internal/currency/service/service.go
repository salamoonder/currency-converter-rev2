package service

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
)

type service struct {
	currency ICurrencyRepository
}

func NewService(currencyRepository ICurrencyRepository) *service {
	return &service{
		currency: currencyRepository,
	}
}

func (s *service) CreateCurrency(ctx context.Context, curr entity.Currency) error {
	return s.currency.CreateCurrency(ctx, curr)
}
func (s *service) UpdateCurrency(ctx context.Context, curr entity.Currency, id string) error {
	return s.currency.UpdateCurrency(ctx, curr, id)
}
func (s *service) CreateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
	return nil
}
func (s *service) UpdateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
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
