// Package repos
package repos

import "github.com/jackc/pgx/v5/pgxpool"

type ShareRepo struct {
	Pool *pgxpool.Pool
}

func NewShareRepo(pool *pgxpool.Pool) *ShareRepo {
	return &ShareRepo{
		Pool: pool,
	}
}
