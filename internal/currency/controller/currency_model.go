package controller

import (
	"currency-converter-rev2/internal/currency/entity"
	"time"
)

type CurrencyCreateView struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Sign string `json:"sign"`
}

func (c *CurrencyCreateView) MapToView(curr entity.Currency) {
	c.Code = curr.Code
	c.Name = curr.Name
	c.Sign = curr.Sign
}

func (c *CurrencyCreateView) MapToEntity() entity.Currency {
	var curr entity.Currency
	curr.Code = c.Code
	curr.Name = c.Name
	curr.Sign = c.Sign
	return curr
}

func (c *CurrencyCreateView) MapToViewList(currIn []entity.Currency) []CurrencyCreateView {
	var currView []CurrencyCreateView
	for _, v := range currIn {
		var curr CurrencyCreateView
		curr.MapToView(v)
		currView = append(currView, curr)
	}
	return currView
}

type CurrencyView struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Sign      string    `json:"sign"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *CurrencyView) MapToView(curr entity.Currency) {
	c.ID = curr.ID
	c.Code = curr.Code
	c.Name = curr.Name
	c.Sign = curr.Sign
	c.CreatedAt = curr.CreatedAt
}
func (c *CurrencyView) MapToViewList(currIn []entity.Currency) []CurrencyView {
	var currView []CurrencyView
	for _, v := range currIn {
		var curr CurrencyView
		curr.MapToView(v)
		currView = append(currView, curr)
	}
	return currView
}
func (c *CurrencyView) MapToEntity() entity.Currency {
	var curr entity.Currency
	curr.ID = c.ID
	curr.Code = c.Code
	curr.Name = c.Name
	curr.Sign = c.Sign
	curr.CreatedAt = c.CreatedAt
	return curr
}
