package controller

import (
	"currency-converter-rev2/internal/currency/entity"
	"time"
)

type CreateExchangeRateView struct {
	BaseCurrency   string  `json:"baseCurrency"`
	TargetCurrency string  `json:"targetCurrency"`
	Rate           float64 `json:"rate"`
}

func (c *CreateExchangeRateView) MapToEntity() entity.ShortExchangeRate {
	var er entity.ShortExchangeRate
	er.BaseCurrency = c.BaseCurrency
	er.TargetCurrency = c.TargetCurrency
	er.Rate = c.Rate
	return er
}

type ExchangeRateView struct {
	ID             string       `json:"id"`
	BaseCurrency   CurrencyView `json:"baseCurrency"`
	TargetCurrency CurrencyView `json:"targetCurrency"`
	Rate           float64      `json:"rate"`
	CreatedAt      time.Time    `json:"createdAt"`
}

func (c *ExchangeRateView) MapToView(er entity.ExchangeRate) ExchangeRateView {
	c.ID = er.ID
	c.BaseCurrency.MapToView(er.BaseCurrency)
	c.TargetCurrency.MapToView(er.TargetCurrency)
	c.Rate = er.Rate
	c.CreatedAt = er.CreatedAt
	return *c
}

func (c *ExchangeRateView) MapToViewList(erIn []entity.ExchangeRate) []ExchangeRateView {
	var erView []ExchangeRateView
	for _, v := range erIn {
		var er ExchangeRateView
		er.MapToView(v)
		erView = append(erView, er)
	}
	return erView
}
