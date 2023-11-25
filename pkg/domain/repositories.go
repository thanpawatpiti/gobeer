package domain

import (
	"context"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

type BeerRepository interface {
	GetBeer(ctx context.Context, name string, page int, per_page int) ([]interface{}, error)
	GetTotalBeer(ctx context.Context, name string) (int, error)
	AddBeer(ctx context.Context, input *entities.Beer) error
}
