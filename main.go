package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vishal1297/crm-basic/database"
	"github.com/vishal1297/crm-basic/lead"
)

const HOST_PORT = "localhost:3000"

func main() {
	app := fiber.New()

	// connect to DB
	initDB()

	setupRoutes(app)

	log.Print("- Application will start at ", HOST_PORT)

	// start server
	log.Fatal(app.Listen(HOST_PORT))

	// at last close DB connection
	defer database.DBConn.Close()
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
	log.Print("- Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	log.Print("- Database Migrated")
}
