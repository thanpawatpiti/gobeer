package repositories

import (
	"context"
	"strconv"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (r *Repository) UpdateBeer(ctx context.Context, entity *entities.Beer, id int) error {
	sqlRaw := "UPDATE beer SET "

	if len(entity.BeerName) > 0 {
		sqlRaw += "beer_name = '" + entity.BeerName + "',"
	}

	if entity.BeerTypeID > 0 {
		parseToString := strconv.Itoa(entity.BeerTypeID)
		sqlRaw += "beer_type_id = " + parseToString + ","
	}

	if len(entity.Image) > 0 {
		sqlRaw += "image = '" + entity.Image + "',"
	}

	sqlRaw = sqlRaw[:len(sqlRaw)-1]

	sqlRaw += " WHERE id = " + strconv.Itoa(id)

	_, err := r.mariaDB.Exec(sqlRaw)
	if err != nil {
		return err
	}

	return nil
}
