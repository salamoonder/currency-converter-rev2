package server

import (
	"context"
	"currency-converter-rev2/internal/currency/controller"
	"currency-converter-rev2/internal/currency/infra/repository"
	"currency-converter-rev2/internal/currency/service"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	host            = "localhost"
	port            = 5432
	user            = "olzhas"
	password        = "password"
	dbname          = "currency"
	serverPort      = "8080"
	maxOpenDbConn   = 10
	maxIdleDbConn   = 5
	maxDbLifeTime   = 5 * time.Minute
	shutdownTimeout = 5 * time.Second
)

var dbConn *DB

type DB struct {
	*sql.DB
}

func ConnectPostgres(dsn string) (*DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(maxOpenDbConn)
	conn.SetMaxIdleConns(maxIdleDbConn)
	conn.SetConnMaxLifetime(maxDbLifeTime)

	if err := testDB(conn); err != nil {
		return nil, err
	}

	return &DB{DB: conn}, nil
}

func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Println("Error pinging database:", err)
		return err
	}
	log.Println("*** Pinged database successfully! ***")
	return nil
}

type App struct {
	httpServer   *http.Server
	currencyServ service.IService
}

func NewApp() (*App, error) {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := ConnectPostgres(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	currencyRepo := repository.NewCurrencyRepo(db.DB)
	return &App{
		currencyServ: service.NewService(currencyRepo),
	}, nil

}

func (a *App) Run() error {
	routes := controller.Routes(a.currencyServ)
	a.httpServer = &http.Server{
		Addr:    ":" + serverPort,
		Handler: routes,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	waitForShutdownSignal()

	ctx, shutdown := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func waitForShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
}
