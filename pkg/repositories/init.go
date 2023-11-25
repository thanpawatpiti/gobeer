package repositories

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mariaDB *sql.DB
	mongoDB *mongo.Client
}

func NewRepository(mariaDB *sql.DB, mongoDB *mongo.Client) *Repository {
	return &Repository{
		mariaDB: mariaDB,
		mongoDB: mongoDB,
	}
}
