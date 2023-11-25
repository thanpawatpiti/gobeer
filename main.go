package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/gobeer/conf/database"
	"github.com/thanpawatpiti/gobeer/pkg/controller"
	"github.com/thanpawatpiti/gobeer/pkg/repositories"
	"github.com/thanpawatpiti/gobeer/pkg/services"
)

func main() {
	database.InitMariaDB()
	database.InitMongoDB()

	// Init repositories
	repo := repositories.NewRepository(database.MariaDB, database.MongoDB)

	// Init services
	service := services.NewService(repo)

	// Init controller
	controller := controller.NewController(service)

	// Init Fiber
	app := fiber.New()

	// Routes group
	api := app.Group("/api/v1")

	// Routes
	api.Get("/beer", controller.GetBeer)
	api.Post("/beer", controller.AddBeer)
	api.Put("/beer/:id", controller.UpdateBeer)
	api.Delete("/beer/:id", controller.DeleteBeer)

	api.Get("/menu", controller.GetMenu)

	// Start server
	app.Listen(":3000")
}
