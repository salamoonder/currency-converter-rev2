package dbconnection

import (
	"context"
	"currency-converter-rev2/internal/config"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	errors2 "github.com/pkg/errors"
	"io/fs"
	"log"
	"log/slog"
)

var dbConn *DB

type DB struct {
	*sql.DB
}

func ConnectPostgres(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%s", cfg.UserDB, cfg.PasswordDB, cfg.HostDB, cfg.PortDB, cfg.NameDB, cfg.MaxOpenDBConn)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errors2.Wrapf(err, "postgresclient.NewPostgres.pgxpool.ParseConfig, failed to parse postgres url")
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	log.Println("connected to DB")

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, errors2.Wrapf(err, "postgresclient.NewPostgres-db.Acquire, error while acquiring connection")
	}

	err = migrateDatabase(ctx, conn.Conn(), cfg.SchemeDB, cfg.MigrationPath)
	if err != nil {
		return nil, errors2.Wrapf(err, "postgresclient.NewPostgres.migrateDatabase, error while migrating database")
	}
	conn.Release()

	return pool, nil
}

func migrateDatabase(ctx context.Context, conn *pgx.Conn, schemeTable string, migrationPath fs.FS) error {
	migrator, err := migrate.NewMigrator(ctx, conn, schemeTable)
	if err != nil {
		return errors2.Wrapf(err, "postgresclient.migrateDatabase.migrate.NewMigrator, failed to create migrator")
	}

	err = migrator.LoadMigrations(migrationPath)
	if err != nil {
		return errors2.Wrapf(err, "postgresclient.migrateDatabase.migrator.LoadMigrations, error while loading migrations")
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		fmt.Println(err, "===")
		dc := errors.New(pgerrcode.DuplicateTable)
		if errors.As(err, &dc) {
			log.Printf("postgresclient.migrateDatabase.migrator.Migrate, failed to migrate, schemeTable: %s", schemeTable)
		}
		if !errors.As(err, &dc) {
			return errors2.Wrapf(err, "postgresclient.migrateDatabase.migrator.Migrate, error while migrating")
		}
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return errors2.Wrapf(err, "postgresclient.migrateDatabase.migrator.GetCurrentVersion, error while getting current version")
	}

	slog.Info("End migrateDatabase, Current", slog.Any("version", ver))
	return nil
}
