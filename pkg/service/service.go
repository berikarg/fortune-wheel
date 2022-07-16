package service

import (
	"github.com/berikarg/fortune-wheel/api"
)

type SpinResult interface {
	GetAll() ([]api.SpinResult, error)
	Create(result api.SpinResult) error
}
