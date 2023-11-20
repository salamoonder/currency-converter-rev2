package entity

import (
	"time"
)

type ExchangeRate struct {
	ID              string
	BaseCurrency    Currency
	TargetCurrency  Currency
	Rate            float64
	Amount          float64
	ConvertedAmount float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Currency struct {
	ID        string
	Code      string
	Name      string
	Sign      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
