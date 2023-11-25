package services

import (
	"context"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (s *Service) GetBeer(ctx context.Context, name *string, page int, per_page int) ([]*entities.Beer, error) {
	return s.repo.GetBeer(ctx, name, page, per_page)
}
