package repositories

import (
	"context"
	"time"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (r *Repository) GetBeer(ctx context.Context, name *string, page int, per_page int) ([]*entities.Beer, error) {
	mockBeer := []*entities.Beer{
		{
			ID:         1,
			BeerName:   "Beer 1",
			BeerTypeID: 1,
			Image:      "Image 1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         2,
			BeerName:   "Beer 2",
			BeerTypeID: 1,
			Image:      "Image 2",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	return mockBeer, nil
}
