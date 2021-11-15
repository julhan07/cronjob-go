package service

import (
	"github.com/julhan07/fiber-service/repo"
)

type service struct {
	repo repo.RepoIndex
}

type ServiceIndex interface {
	CronJob()
}

func NewServiceIndex(repo repo.RepoIndex) ServiceIndex {
	return &service{
		repo: repo,
	}
}
