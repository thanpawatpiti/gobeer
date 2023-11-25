package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

func (c *Controller) UpdateBeer(ctx *fiber.Ctx) error {
	// Get form data
	id := ctx.Params("id")
	name := ctx.FormValue("name")
	beer_type_id := ctx.FormValue("beer_type_id")
	image, _ := ctx.FormFile("image")

	// Validate form data
	if len(id) < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "id is required",
		})
	}

	// Parse to int
	idInt := c.parseToInt(id)

	entity := &entities.Beer{}

	if len(name) > 0 {
		entity.BeerName = name
	}

	if len(beer_type_id) > 0 {
		entity.BeerTypeID = c.parseToInt(beer_type_id)
	}

	// Update beer
	err := c.service.UpdateBeer(ctx.Context(), entity, image, idInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success",
	})
}
