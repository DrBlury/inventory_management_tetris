package repo

import (
	"context"
	"fmt"

	generated "linuxcode/inventory_manager/pkg/repo/generated"

	"github.com/jackc/pgx/v5"
)

func Connect(cfg *Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)
	return pgx.Connect(context.Background(), connString)
}

func CreateDB(cfg *Config) (*generated.Queries, error) {
	connection, err := Connect(cfg)
	if err != nil {
		return nil, err
	}
	return generated.New(connection), nil
}
