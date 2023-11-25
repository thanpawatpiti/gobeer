package services

import (
	"github.com/thanpawatpiti/gobeer/pkg/domain"
)

type Service struct {
	repo domain.BeerRepository
}

func NewService(repo domain.BeerRepository) domain.BeerSerives {
	return &Service{
		repo: repo,
	}
}
