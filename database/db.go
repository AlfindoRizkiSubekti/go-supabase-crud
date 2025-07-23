package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// DB adalah pool koneksi database yang akan digunakan di seluruh aplikasi
var DB *pgxpool.Pool

// InitDB menginisialisasi koneksi database
func InitDB() error {
	// 1. Load environment variables
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// 2. Dapatkan connection string dari environment
	connStr := os.Getenv("DB_CONNECTION")
	if connStr == "" {
		return fmt.Errorf("DB_CONNECTION environment variable not set")
	}

	// 3. Parse configuration
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("unable to parse database config: %v", err)
	}

	// 4. Konfigurasi koneksi
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	// 5. Buat connection pool
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	// 6. Test koneksi
	conn, err := DB.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %v", err)
	}
	defer conn.Release()

	return nil
}
