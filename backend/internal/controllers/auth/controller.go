package controllers_auth

import (
	repos_mysql_user "github.com/IlliaSh1/backend/internal/repos/mysql/user"
	services_auth "github.com/IlliaSh1/backend/internal/services/auth"
	"github.com/gofiber/fiber/v3"
)

type authController struct {
	authService *services_auth.AuthService
	userService *repos_mysql_user.UserRepo
}

func AddAuthControllerRoutes(api fiber.Router, authService *services_auth.AuthService, userService *repos_mysql_user.UserRepo) {
	controller := authController{
		authService: authService,
		userService: userService,
	}

	api.Post("/register", controller.Register)
}
