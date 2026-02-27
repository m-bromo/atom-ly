package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/m-bromo/atom-ly/config"
)

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		cfg.Env.PostgresDB.Name,
		cfg.Env.PostgresDB.User,
		cfg.Env.PostgresDB.Password,
		cfg.Env.PostgresDB.Host,
		cfg.Env.PostgresDB.Port,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
