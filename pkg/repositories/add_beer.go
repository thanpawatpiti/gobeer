package repositories

import (
	"context"
	"fmt"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (r *Repository) GetBeerType(ctx context.Context, id int) error {
	result, err := r.mariaDB.Query("SELECT count(*) FROM beer_type WHERE id = ?", id)
	if err != nil {
		return err
	}

	var count int
	for result.Next() {
		err := result.Scan(&count)
		if err != nil {
			return err
		}
	}

	if count == 0 {
		return fmt.Errorf("beer type id %d not found", id)
	}

	return nil
}

func (r *Repository) AddBeer(ctx context.Context, input *entities.Beer) error {
	_, err := r.mariaDB.Exec("INSERT INTO beer (beer_name, beer_type_id, image) VALUES (?, ?, ?)", input.BeerName, input.BeerTypeID, input.Image)
	if err != nil {
		return err
	}

	return nil
}
