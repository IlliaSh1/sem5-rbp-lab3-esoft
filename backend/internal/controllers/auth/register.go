package controllers_auth

import (
	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller authController) Register(ctx fiber.Ctx) error {
	req := &RegisterRequest{}

	err := ctx.Bind().Body(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = req.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	password_hash, err := controller.authService.HashPassword(req.Password)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	user := models.User{
		Email:        req.Email,
		PasswordHash: password_hash,
	}

	err = controller.userService.Save(&user)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterRequest) Validate() (err error) {
	if req.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}
	if req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "password is required")
	}
	return nil
}
