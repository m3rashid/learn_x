package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	user  = "User"
	admin = "Admin"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := getUserFromDb()
		if user.Role != "Admin" {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		return fn(ctx)
	}
}

func handleBaseRoute(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

func getUserFromDb() User {
	return User{
		Username: "Rashid",
		Role:     admin,
	}
}

func main() {
	app := fiber.New()

	app.Get("/", handleBaseRoute)
	app.Get("/admin", onlyAdmin(handleBaseRoute))

	log.Fatal(app.Listen(":4000"))
}
