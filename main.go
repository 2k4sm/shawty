package main

import (
	"log"
	"os"
	"time"

	"github.com/2k4sm/shawty/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New(
		fiber.Config{
			AppName: "Shawty",
		},
	)

	setUpRoutes(app)

	app.Use(logger.New())

	port := ":" + os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "server started on Port :" + port,
			"time":    time.Now(),
		})
	})

	log.Fatal(app.Listen(port))
}
