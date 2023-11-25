package entities

import "time"

type Beer struct {
	ID         int64     `json:"id"`
	BeerName   string    `json:"beer_name"`
	BeerTypeID int64     `json:"beer_type_id"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeleteAt   time.Time `json:"delete_at"`
}
