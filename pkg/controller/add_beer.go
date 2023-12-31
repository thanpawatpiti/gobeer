package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (c *Controller) AddBeer(ctx *fiber.Ctx) error {
	// Get form data
	name := ctx.FormValue("name")
	beer_type_id := ctx.FormValue("beer_type_id")
	image, _ := ctx.FormFile("image")

	// Validate form data
	if len(name) < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "name is required",
		})
	}

	if len(beer_type_id) < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "beer_type_id is required",
		})
	}

	// Parse to int
	beer_type_idInt := c.parseToInt(beer_type_id)

	entity := &entities.Beer{
		BeerName:   name,
		BeerTypeID: beer_type_idInt,
	}

	// Add beer
	err := c.service.AddBeer(ctx.Context(), entity, image)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "success",
	})
}
