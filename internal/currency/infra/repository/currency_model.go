package repository

import (
	"currency-converter-rev2/internal/currency/entity"
	"time"
)

type CurrencyData struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Sign      string    `json:"sign"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func (c *CurrencyData) MapToEntity() entity.Currency {
	var curr entity.Currency
	curr.ID = c.ID
	curr.Code = c.Code
	curr.Name = c.Name
	curr.Sign = c.Sign
	curr.CreatedAt = c.CreatedAt
	curr.UpdatedAt = c.UpdatedAt
	curr.DeletedAt = c.DeletedAt
	return curr
}

func (c *CurrencyData) MapToRepo(curr entity.Currency) error {
	c.ID = curr.ID
	c.Code = curr.Code
	c.Name = curr.Name
	c.Sign = curr.Sign
	c.CreatedAt = curr.CreatedAt
	c.UpdatedAt = curr.UpdatedAt
	c.DeletedAt = curr.DeletedAt
	return nil
}
