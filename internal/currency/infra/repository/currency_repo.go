package repository

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type currencyRepo struct {
	db *pgxpool.Pool
}

func NewCurrencyRepo(db *pgxpool.Pool) *currencyRepo {
	return &currencyRepo{
		db: db,
	}
}

func (r currencyRepo) CreateCurrency(ctx context.Context, curr entity.Currency) error {
	currRepo := CurrencyData{}
	currRepo.MapToRepo(curr)
	query := `INSERT INTO currencies (code,full_name,sign)
			VALUES ($1, $2, $3) returning *`
	_, err := r.db.Exec(
		ctx,
		query,
		currRepo.Code,
		currRepo.Name,
		currRepo.Sign,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r currencyRepo) UpdateCurrency(ctx context.Context, curr entity.Currency) error {
	currRepo := CurrencyData{}
	currRepo.MapToRepo(curr)
	query := `
		UPDATE currencies
		SET full_name = $2, sign = $3
		WHERE code = $1 returning *
	`
	_, err := r.db.Exec(ctx, query, currRepo.Code, currRepo.Name, currRepo.Sign)
	if err != nil {
		return err
	}
	return nil
}

func (r currencyRepo) GetAllCurrencies(ctx context.Context) ([]entity.Currency, error) {
	query := "SELECT id, code,full_name,sign FROM currencies"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var currencies []entity.Currency
	for rows.Next() {
		var currency entity.Currency
		err := rows.Scan(&currency.ID, &currency.Code, &currency.Name, &currency.Sign)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v\n", err)
		return nil, err
	}

	if len(currencies) == 0 {
		return nil, errors.New("no currencies found")
	}

	return currencies, nil
}

func (r currencyRepo) GetCurrencyById(ctx context.Context, id string) (entity.Currency, error) {
	return entity.Currency{}, nil
}
