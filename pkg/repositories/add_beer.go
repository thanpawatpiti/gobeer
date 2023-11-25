package repositories

import (
	"context"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (r *Repository) AddBeer(ctx context.Context, input *entities.Beer) error {
	_, err := r.mariaDB.Exec("INSERT INTO beer (beer_name, beer_type_id, image) VALUES (?, ?, ?)", input.BeerName, input.BeerTypeID, input.Image)
	if err != nil {
		return err
	}

	return nil
}
