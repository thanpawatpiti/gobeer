package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

// InitMongoDB is a function that connect to MongoDB
func InitMongoDB() error {
	// Set client options
	option := options.Client().ApplyURI("mongodb+srv://thanpawatpiti:h55Qjmjb6HoTx18H@cluster0.jxlfooc.mongodb.net/?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	MongoDB = client

	fmt.Println("Connected to MongoDB!")

	return nil
}
