package main

import (
	"os"
	"time"

	"github.com/2k4sm/shawty/database"
	"github.com/2k4sm/shawty/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	dbConf := database.DBConfig{
		DB_HOST:   os.Getenv("DB_HOST"),
		DB_USER:   os.Getenv("DB_USER"),
		DB_PG_PWD: os.Getenv("DB_PG_PWD"),
		DB_NAME:   os.Getenv("DB_NAME"),
		DB_PORT:   os.Getenv("DB_PORT"),
		SSLMODE:   os.Getenv("SSLMODE"),
	}

	db := database.InitPGdb(&dbConf)

	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New(
		fiber.Config{
			AppName: "Shawty",
		},
	)
	app.Use(logger.New())

	api := app.Group("/api")
	routes.SetupShawtyRoutes(api)
	routes.SetupUserRoutes(api, db)

	port := ":" + os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "muah😘..." + port,
			"time":    time.Now(),
		})
	})

	log.Fatal(app.Listen(port))
}
