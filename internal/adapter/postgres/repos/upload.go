// Package repos
package repos

import "github.com/jackc/pgx/v5/pgxpool"

type UploadRepo struct {
	Pool *pgxpool.Pool
}

func NewUploadRepo(pool *pgxpool.Pool) *UploadRepo {
	return &UploadRepo{
		Pool: pool,
	}
}



