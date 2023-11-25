package repositories

type MongoRepository struct{}

func NewMongoRepository() *MongoRepository {
	return &MongoRepository{}
}
