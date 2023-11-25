package repositories

import (
	"context"
	"fmt"
)

func (r *Repository) GetBeer(ctx context.Context, name string, page int, per_page int) ([]interface{}, error) {
	var beers []interface{}

	sqlRaw := ""
	if name == "" {
		sqlRaw = fmt.Sprintf("SELECT beer.id, beer.beer_name, beer_type.type_name, beer.image FROM beer INNER JOIN beer_type ON beer.beer_type_id = beer_type.id WHERE beer.deleted_at IS NULL ORDER BY beer.id ASC OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", (page-1)*per_page, per_page)
	} else {
		sqlRaw = fmt.Sprintf("SELECT beer.id, beer.beer_name, beer_type.type_name, beer.image FROM beer INNER JOIN beer_type ON beer.beer_type_id = beer_type.id WHERE beer.deleted_at IS NULL AND beer.beer_name LIKE \"%s\" ORDER BY beer.id ASC OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", name, (page-1)*per_page, per_page)
	}

	rows, err := r.mariaDB.Query(sqlRaw)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var beer_name string
		var type_name string
		var image string

		err := rows.Scan(&id, &beer_name, &type_name, &image)
		if err != nil {
			return nil, err
		}

		data := map[string]interface{}{
			"id":    id,
			"name":  beer_name,
			"type":  type_name,
			"image": image,
		}

		beers = append(beers, data)
	}

	return beers, nil
}

func (r *Repository) GetTotalBeer(ctx context.Context, name string) (int, error) {
	sqlRaw := ""
	if name == "" {
		sqlRaw = "SELECT COUNT(*) FROM beer WHERE deleted_at IS NULL"
	} else {
		sqlRaw = fmt.Sprintf("SELECT COUNT(*) FROM beer WHERE deleted_at IS NULL AND beer_name LIKE \"%s\"", name)
	}

	var total int
	err := r.mariaDB.QueryRow(sqlRaw).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
