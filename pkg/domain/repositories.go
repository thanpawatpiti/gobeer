package domain

import (
	"context"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

type BeerRepository interface {
	GetBeer(ctx context.Context, name *string, page int, per_page int) ([]*entities.Beer, error)
}
