package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"taskboard/auth-service/internal/config"
)

type Base struct {
	pool *pgxpool.Pool
}

func NewConnectionPool(config config.DBConfig) (*Base, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User, config.Password, config.Host, config.Port, config.Database,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open postgres connection: %w", err)
	}
	return &Base{pool: pool}, nil
}

func (s *Base) Stop() {
	s.pool.Close()
}
