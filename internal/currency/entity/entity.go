package entity

type ExchangeRate struct {
	ID              string
	BaseCurrency    Currency
	TargetCurrency  Currency
	Rate            float64
	Amount          float64
	ConvertedAmount float64
}

type Currency struct {
	ID   string
	Code string
	Name string
	Sign string
}
