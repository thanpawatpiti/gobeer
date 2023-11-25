package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/gobeer/pkg/utilities"
)

func (c *Controller) GetMenu(ctx *fiber.Ctx) error {
	menu, err := c.service.GetMenu(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(utilities.Response(true, 200, menu))
}
