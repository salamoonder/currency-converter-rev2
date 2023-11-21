package main

import (
	"context"
	"currency-converter-rev2/internal/config"
	"currency-converter-rev2/internal/currency/controller"
	"currency-converter-rev2/internal/currency/infra/repository"
	"currency-converter-rev2/internal/currency/service"
	"currency-converter-rev2/internal/dbconnection"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

var shutDownTimeOut = 10 * time.Second

func main() {
	ctx, shutDown := context.WithTimeout(context.Background(), shutDownTimeOut)
	defer shutDown()

	cfg := config.NewConfig()
	pool, err := dbconnection.ConnectPostgres(ctx, cfg)
	if err != nil {
		log.Println(err)
		return
	}
	currencyRepo := repository.NewCurrencyRepo(pool)
	exchangeRateRepo := repository.NewExchangeRateRepo(pool)
	service := service.NewService(currencyRepo, exchangeRateRepo)

	contr := controller.NewHandler(service)
	mux := contr.InitRoutes()
	startHttp(mux, cfg.HTTPPort)
}

func startHttp(r *chi.Mux, port string) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
