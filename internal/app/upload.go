// Package app
package app

import "main/internal/adapter/postgres/repos"

type UplaodService struct {
	Repo *repos.UploadRepo
}

func NewUploadService(repo *repos.UploadRepo) *UplaodService {
	return &UplaodService{
		Repo: repo,
	}
}
