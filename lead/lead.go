package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vishal1297/crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	if lead.Name == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No lead found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Lead fetched successfully", "data": lead})
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(fiber.Map{"status": "success", "message": "All leads fetched successfully", "data": leads})
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": "Unable to parse lead", "data": nil})
	}
	db.Create(&lead)
	return c.JSON(fiber.Map{"status": "success", "message": "Lead created successfully", "data": lead})
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No lead found with ID", "data": nil})
	}
	db.Delete(&lead)
	return c.JSON(fiber.Map{"status": "success", "message": "Lead deleted successfully", "data": lead})
}
