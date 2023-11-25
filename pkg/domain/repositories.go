package domain

import (
	"context"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

type BeerRepository interface {
	GetBeer(ctx context.Context, name string, page int, per_page int) ([]interface{}, error)
	GetTotalBeer(ctx context.Context, name string) (int, error)
	GetBeerType(ctx context.Context, id int) error
	AddBeer(ctx context.Context, input *entities.Beer) error
	CreateLog(ctx context.Context, action string, input interface{}) error
	UpdateBeer(ctx context.Context, input *entities.Beer, id int) error
	DeleteBeer(ctx context.Context, id int) error
	GetMenu(ctx context.Context) (interface{}, error)
}
