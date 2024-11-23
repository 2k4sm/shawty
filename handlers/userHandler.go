package handlers

import (
	"net/http"

	"github.com/2k4sm/shawty/dto"
	"github.com/2k4sm/shawty/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	LoginHandler(ctx *fiber.Ctx) error
	SignUpHandler(ctx *fiber.Ctx) error
	// UpdatePassHandler(ctx *fiber.Ctx) error
	// DeleteUserHandler(ctx *fiber.Ctx) error
}

type UserHandler struct {
	service services.UserServInterface
}

func NewUserHandler(serv services.UserServInterface) UserHandlerInterface {
	return &UserHandler{
		service: serv,
	}
}

func (uh *UserHandler) LoginHandler(ctx *fiber.Ctx) error {
	var existingUser dto.UserAuth

	if err := ctx.BodyParser(&existingUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(existingUser)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := uh.service.Login(&existingUser)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp := fiber.Map{
		"token" : token,
	}

	return ctx.Status(http.StatusCreated).JSON(resp)
}

func (uh UserHandler) SignUpHandler(ctx *fiber.Ctx) error {
	var newUser dto.UserAuth

	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(newUser)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := uh.service.SignUp(&newUser)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resp := fiber.Map{
		"token" : token,
	}

	return ctx.Status(http.StatusCreated).JSON(resp)
}

// func (uh *UserHandler) UpdatePassHandler(ctx *fiber.Ctx) error {

// }

// func (uh *UserHandler) DeleteUserHandler(ctx *fiber.Ctx) error {

// }
