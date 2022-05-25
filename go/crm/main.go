package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/m3rashid/learn_x/go/crm/database"
	"github.com/m3rashid/learn_x/go/crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/lead", lead.GetLeads)
	app.Get("/api/lead/:id", lead.GetOneLead)
	app.Post("/api/lead", lead.NewLead)
	app.Delete("/api/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DbConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Connection opened to database")
	database.DbConn.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(":5000")
	defer database.DbConn.Close()
}
