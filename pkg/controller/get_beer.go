package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) GetBeer(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, GetBeer 👋!")
}
