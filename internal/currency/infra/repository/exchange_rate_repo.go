package repository

import (
	"context"
	"currency-converter-rev2/internal/currency/entity"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type exchangeRateRepo struct {
	db *pgxpool.Pool
}

func NewExchangeRateRepo(db *pgxpool.Pool) *exchangeRateRepo {
	return &exchangeRateRepo{
		db: db,
	}
}

func (r exchangeRateRepo) CreateExchangeRate(ctx context.Context, exch entity.ShortExchangeRate) error {
	exchRepo := ShortExchangeRateData{}
	exchRepo.MapToRepo(exch)
	query := `INSERT INTO exchange_rates (base_currency_id, target_currency_id, rate, created_at)
			VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(
		ctx,
		query,
		exchRepo.BaseCurrency,
		exchRepo.TargetCurrency,
		exchRepo.Rate,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r exchangeRateRepo) UpdateExchangeRate(ctx context.Context, exch entity.ExchangeRate) error {
	exchRepo := ExchangeRateData{}
	exchRepo.MapToRepo(exch)
	return nil
}
func (r exchangeRateRepo) GetAllExchangeRates(ctx context.Context) ([]entity.ExchangeRate, error) {

	query := `select exchange_rates.id , exchange_rates.created_at,
       		  c.id, c.code,c.full_name,c.sign,c.created_at,
    	      c2.id, c2.code, c2.full_name, c2.sign,c2.created_at,
   			  exchange_rates.rate from exchange_rates
  			  join currencies c on c.id= exchange_rates.base_currency_id
 			  join currencies c2 on c2.id = exchange_rates.target_currency_id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var eRates []entity.ExchangeRate
	for rows.Next() {
		var eRate entity.ExchangeRate
		err = rows.Scan(
			&eRate.ID,
			&eRate.CreatedAt,
			&eRate.BaseCurrency.ID,
			&eRate.BaseCurrency.Code,
			&eRate.BaseCurrency.Name,
			&eRate.BaseCurrency.Sign,
			&eRate.BaseCurrency.CreatedAt,
			&eRate.TargetCurrency.ID,
			&eRate.TargetCurrency.Code,
			&eRate.TargetCurrency.Name,
			&eRate.TargetCurrency.Sign,
			&eRate.TargetCurrency.CreatedAt,
			&eRate.Rate,
		)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		eRates = append(eRates, eRate)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v\n", err)
		return nil, err
	}

	if len(eRates) == 0 {
		return nil, errors.New("no currencies found")
	}

	return eRates, nil
}
func (r exchangeRateRepo) GetExchangeRateByCode(ctx context.Context, code string) (entity.ExchangeRate, error) {

	query := `  SELECT exchange_rates.id , exchange_rates.created_at,
				c.id, c.code,c.full_name,c.sign,c.created_at,
				c2.id, c2.code, c2.full_name, c2.sign,c2.created_at,
				exchange_rates.rate FROM exchange_rates
				JOIN currencies c ON c.id= exchange_rates.base_currency_id
				JOIN currencies c2 ON c2.id = exchange_rates.target_currency_id
				WHERE c.code = $1 AND c2.code = $2
			`
	var exchangeRate ExchangeRateData
	var codeFrom = code[:3]
	var codeTo = code[3:]
	row := r.db.QueryRow(ctx, query, codeFrom, codeTo)
	err := row.Scan(
		&exchangeRate.ID,
		&exchangeRate.CreatedAt,
		&exchangeRate.BaseCurrency.ID,
		&exchangeRate.BaseCurrency.Code,
		&exchangeRate.BaseCurrency.Name,
		&exchangeRate.BaseCurrency.Sign,
		&exchangeRate.BaseCurrency.CreatedAt,
		&exchangeRate.TargetCurrency.ID,
		&exchangeRate.TargetCurrency.Code,
		&exchangeRate.TargetCurrency.Name,
		&exchangeRate.TargetCurrency.Sign,
		&exchangeRate.TargetCurrency.CreatedAt,
		&exchangeRate.Rate,
	)
	if err != nil {
		log.Printf("Error scanning row: %v\n", err)
		return exchangeRate.MapToEntity(), err
	}
	return exchangeRate.MapToEntity(), nil
}
