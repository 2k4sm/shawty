package routes

import (
	"github.com/2k4sm/shawty/handlers"
	"github.com/2k4sm/shawty/repositories"
	"github.com/2k4sm/shawty/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(router fiber.Router, db *gorm.DB) {
	userRepos := repositories.NewUserRepo(db)
	userServs := services.NewUserServ(userRepos)
	userHandlers := handlers.NewUserHandler(userServs)

	router.Post("/signup", userHandlers.SignUpHandler)
	router.Post("/login", userHandlers.LoginHandler)
}

func SetupShawtyRoutes(router fiber.Router) {
	router.Get("/:url", ResolveURL)
	router.Post("/", ShortenURL)
}
