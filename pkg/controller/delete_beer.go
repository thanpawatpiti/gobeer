package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) DeleteBeer(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, DeleteBeer 👋!")
}
