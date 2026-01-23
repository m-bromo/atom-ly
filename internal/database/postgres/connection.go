package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/m-bromo/atom-ly/config"
)

func NewPostgresConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		config.Env.PostgresDB.Name,
		config.Env.PostgresDB.User,
		config.Env.PostgresDB.Password,
		config.Env.PostgresDB.Host,
		config.Env.PostgresDB.Port,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("Failed to connect to postgres database", "error", err)
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		slog.Error("Failed to ping postgres database", "error", err)
		return nil, err
	}

	slog.Info("Connected to postgres database")

	return conn, nil
}
