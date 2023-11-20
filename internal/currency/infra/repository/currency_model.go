package repository

import (
	"currency-converter-rev2/internal/currency/entity"
)

type CurrencyData struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Sign string `json:"sign"`
}


func (c *CurrencyData) MapToEntity() entity.Currency {
	var curr entity.Currency
	curr.ID = c.ID
	curr.Code = c.Code
	curr.Name = c.Name
	curr.Sign = c.Sign
	return curr
}

func (c *CurrencyData) MapToRepo(curr entity.Currency) error {
	c.ID = curr.ID
	c.Code = curr.Code
	c.Name = curr.Name
	c.Sign = curr.Sign
	return nil
}
