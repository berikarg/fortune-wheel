package service

import (
	"time"

	"github.com/berikarg/fortune-wheel/api"
	"github.com/berikarg/fortune-wheel/pkg/repository"
)

type SpinResultService struct {
	repo repository.SpinResult
}

func NewSpinResultService(repo repository.SpinResult) *SpinResultService {
	return &SpinResultService{repo: repo}
}

func (s *SpinResultService) GetAll() ([]api.SpinResult, error) {
	return s.repo.GetAll()
}
func (s *SpinResultService) Create(result api.SpinResult) error {
	result.Timestamp = time.Now()
	return s.repo.Create(result)
}
