package repositories

import (
	"context"
	"encoding/json"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (r *Repository) GetMenu(ctx context.Context) (interface{}, error) {
	result, err := r.mariaDB.Query("SELECT m.id,m.name,m.route,m.icon,m.is_children,JSON_ARRAYAGG(JSON_OBJECT('id', mc.id,'parent_id', mc.parent_id,'name', mc.name,'route', mc.route)) AS children FROM menus m LEFT JOIN menu_children mc ON m.id = mc.parent_id GROUP BY m.id ORDER BY m.id ASC")

	if err != nil {
		return nil, err
	}

	var menu []entities.Menu
	for result.Next() {
		var id int
		var name string
		var route string
		var icon string
		var is_children bool
		var children interface{}

		err := result.Scan(&id, &name, &route, &icon, &is_children, &children)
		if err != nil {
			return nil, err
		}

		var childrens []entities.MenuChildren
		json.Unmarshal(children.([]byte), &childrens)

		data := entities.Menu{
			ID:         id,
			Name:       name,
			Route:      route,
			Icon:       icon,
			IsChildren: is_children,
			Children:   childrens,
		}

		menu = append(menu, data)
	}

	return menu, nil
}
