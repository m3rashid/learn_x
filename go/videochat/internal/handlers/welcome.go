package handlers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) {
	return c.Render("welcome", nil, "layouts/main")
}
