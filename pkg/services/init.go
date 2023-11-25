package services

import "github.com/thanpawatpiti/gobeer/pkg/repositories"

type Service struct {
	mariaDB repositories.MariaRepository
	mongoDB repositories.MongoRepository
}

func NewService(mariaDB repositories.MariaRepository, mongoDB repositories.MongoRepository) *Service {
	return &Service{
		mariaDB: mariaDB,
		mongoDB: mongoDB,
	}
}
