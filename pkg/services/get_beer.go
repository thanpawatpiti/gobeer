package services

import (
	"context"
)

func (s *Service) GetBeer(ctx context.Context, name string, page int, per_page int) ([]interface{}, error) {
	return s.repo.GetBeer(ctx, name, page, per_page)
}
