package models

import "time"

type Beer struct {
	ID         int       `gorm:"primary_key" json:"id"`
	BeerName   string    `gorm:"type:varchar(250)" json:"beer_name"`
	BeerTypeID int       `gorm:"type:int(10)" json:"beer_type_id"`
	Image      string    `gorm:"type:varchar(250)" json:"image"`
	CreatedAt  time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime" json:"updated_at"`
	DeleteAt   time.Time `gorm:"type:datetime" json:"delete_at"`
}
