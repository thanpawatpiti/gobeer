package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetBeer(ctx *fiber.Ctx) error {
	// Get query params
	name := ctx.Query("name")
	page := ctx.Query("page")
	per_page := ctx.Query("per_page")

	// Validate query params
	if name == "" {
		name = "%%"
	} else {
		name = "%" + name + "%"
	}

	if page == "" {
		page = "1"
	}

	if per_page == "" {
		per_page = "10"
	}

	// Parse to int
	pageInt := c.parseToInt(page)
	per_pageInt := c.parseToInt(per_page)

	// Get beers
	beers, err := c.service.GetBeer(ctx.Context(), &name, pageInt, per_pageInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success",
		"data":    beers,
	})
}

func (c *Controller) parseToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
