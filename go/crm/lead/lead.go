package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/m3rashid/learn_x/go/crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DbConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetOneLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
	return nil
}

func NewLead(c *fiber.Ctx) error {
	db := database.DbConn
	var lead Lead
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(&lead)
	c.JSON(lead)
	return nil
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with ID"))
	}
	db.Delete(&lead)
	c.Send([]byte("Lead deleted successfully"))
	return nil
}
