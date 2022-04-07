package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main () {
	todos := []Todo{}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	app.Get("/", func (c *fiber.Ctx) error {
  	return c.SendString("Hello, World!")
  })

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if error := c.BodyParser(todo); error != nil {
			return error
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, error := c.ParamsInt("id")
		if error!=nil {
			return c.Status(400).SendString("Invalid ID")
		}

		for i, t :=range todos {
			if t.ID == id {
				todos[i].Done = !todos[i].Done
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":5000"))
}