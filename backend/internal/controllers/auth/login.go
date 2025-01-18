package controllers_auth

import (
	"errors"

	repos_mysql_user "github.com/IlliaSh1/backend/internal/repos/mysql/user"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller authController) Login(ctx fiber.Ctx) error {
	req := &LoginRequest{}

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

	user, err := controller.userService.GetByEmail(req.Email)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_user.ErrUserNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		default:
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	err = controller.authService.ComparePasswordAndHash(req.Password, user.PasswordHash)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "email or password is incorrect",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginRequest) Validate() (err error) {
	if req.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}
	if req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "password is required")
	}
	return nil
}
