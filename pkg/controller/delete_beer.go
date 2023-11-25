package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) DeleteBeer(ctx *fiber.Ctx) error {
	// Get form data
	id := ctx.Params("id")

	// Validate form data
	if len(id) < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "id is required",
		})
	}

	// Parse to int
	idInt := c.parseToInt(id)

	// Delete beer
	err := c.service.DeleteBeer(ctx.Context(), idInt)
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
