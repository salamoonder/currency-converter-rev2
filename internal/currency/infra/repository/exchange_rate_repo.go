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

	query := `select exchange_rates.id ,
       		  c.id, c.code,c.full_name,c.sign,
    	      c2.id, c2.code, c2.full_name, c2.sign,
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
		err := rows.Scan(
			&eRate.ID,
			&eRate.BaseCurrency.ID,
			&eRate.BaseCurrency.Code,
			&eRate.BaseCurrency.Name,
			&eRate.BaseCurrency.Sign,
			&eRate.TargetCurrency.ID,
			&eRate.TargetCurrency.Code,
			&eRate.TargetCurrency.Name,
			&eRate.TargetCurrency.Sign,
			&eRate.Rate,
		)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		eRates = append(eRates, eRate)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v\n", err)
		return nil, err
	}

	if len(eRates) == 0 {
		return nil, errors.New("no currencies found")
	}

	return eRates, nil
}
