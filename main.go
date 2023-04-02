package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vishal1297/crm-basic/database"
	"github.com/vishal1297/crm-basic/lead"
)

const PORT = 3000

func main() {
	app := fiber.New()

	// connect to DB
	initDB()

	setupRoutes(app)
	app.Listen(PORT)

	fmt.Println("Application started at PORT", PORT)

	// at last close DB connection
	defer database.DBConn.Close()

	fmt.Println("Application is running...")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Database connection failed")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
