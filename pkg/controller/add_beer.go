package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) AddBeer(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, AddBeer 👋!")
}
