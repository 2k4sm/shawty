package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/2k4sm/shawty/dto"
	"github.com/2k4sm/shawty/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	// LoginHandler(ctx *fiber.Ctx) error
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

// func (uh *UserHandler) LoginHandler(ctx *fiber.Ctx) error {

// }

func (uh *UserHandler) SignUpHandler(ctx *fiber.Ctx) error {
	var newUser dto.UserSignup

	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := uh.service.SignUp(&newUser)

	if err != nil && errors.Is(err, fmt.Errorf("user already exists")) {
		return ctx.Status(http.StatusConflict).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userResp := dto.UserSignup{}

	userResp.Email = user.Email
	userResp.Name = user.Name
	userResp.Password = user.Password

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(&userResp)
}

// func (uh *UserHandler) UpdatePassHandler(ctx *fiber.Ctx) error {

// }

// func (uh *UserHandler) DeleteUserHandler(ctx *fiber.Ctx) error {

// }
