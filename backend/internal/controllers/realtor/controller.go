package controllers_realtor

import (
	repos_mysql_realtor "github.com/IlliaSh1/backend/internal/repos/mysql/realtor"
	"github.com/gofiber/fiber/v3"
)

type realtorController struct {
	realtorService *repos_mysql_realtor.RealtorRepo
}

func AddRealtorController(
	api fiber.Router,
	realtorService *repos_mysql_realtor.RealtorRepo,
) error {
	controller := realtorController{
		realtorService: realtorService,
	}

	api.Post("/realtor", controller.Create)

	return nil
}
