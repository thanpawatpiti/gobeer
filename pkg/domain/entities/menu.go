package entities

type Menu struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Route      string         `json:"route"`
	Icon       string         `json:"icon"`
	IsChildren bool           `json:"is_children"`
	Children   []MenuChildren `json:"children"`
}

type MenuChildren struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Route    string `json:"route"`
}
