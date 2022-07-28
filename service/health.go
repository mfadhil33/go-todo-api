package service

import "github.com/edwintantawi/go-todo-api/entity/health"

type healthService struct {
	repository health.Repository
}

func NewHealthService(r health.Repository) *healthService {
	return &healthService{repository: r}
}

func (s *healthService) Check() error {
	return s.repository.Ping()
}
