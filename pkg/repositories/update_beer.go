package repositories

import (
	"context"
	"fmt"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (r *Repository) UpdateBeer(ctx context.Context, entity *entities.Beer, id int) error {
	sqlRaw := ""
	if entity.Image != "" {
		sqlRaw = fmt.Sprintf("UPDATE beer SET beer_name = ?, beer_type_id = ?, image = '%s' WHERE id = ?", entity.Image)
	} else {
		sqlRaw = "UPDATE beer SET beer_name = ?, beer_type_id = ? WHERE id = ?"
	}

	_, err := r.mariaDB.Exec(sqlRaw, entity.BeerName, entity.BeerTypeID, id)
	if err != nil {
		return err
	}

	return nil
}
