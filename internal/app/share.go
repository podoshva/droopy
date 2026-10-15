// Package app
package app

import "main/internal/adapter/postgres/repos"

type ShareService struct {
	Repo *repos.ShareRepo
}

func NewShareService(repo *repos.ShareRepo) *ShareService {
	return &ShareService{
		Repo: repo,
	}
}
