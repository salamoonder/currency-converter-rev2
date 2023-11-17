package service

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
	"log"
)

type service struct {
	currency ICurrencyRepository
}

func NewService(currencyRepository ICurrencyRepository) *service {
	return &service{
		currency: currencyRepository,
	}
}

func (s *service) CreateCurrency(ctx context.Context, curr entity.Currency) (*entity.Currency, error) {
	res, err := s.currency.CreateCurrency(ctx, curr)
	if err != nil {
		log.Println("cannot create currency")
	}
	return res, err
}
func (s *service) UpdateCurrency(ctx context.Context, curr entity.Currency) (*entity.Currency, error) {
	res, err := s.currency.UpdateCurrency(ctx, curr)
	if err != nil {
		log.Println("cannot update currency")
	}
	return res, err
}
func (s *service) CreateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
	return nil
}
func (s *service) UpdateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
	return nil
}
func (s *service) GetAllCurrencies(ctx context.Context) ([]*entity.Currency, error) {
	return s.currency.GetAllCurrencies(ctx)
}
func (s *service) GetCurrencyById(ctx context.Context, id string) (*entity.Currency, error) {
	return s.currency.GetCurrencyById(ctx, "")
}
