package controllers_deal

import (
	repos_mysql_deal "github.com/IlliaSh1/backend/internal/repos/mysql/deal"
	"github.com/gofiber/fiber/v3"
)

type dealController struct {
	dealService *repos_mysql_deal.DealRepo
}

func AddDealControllerRoutes(api fiber.Router, dealService *repos_mysql_deal.DealRepo) {
	controller := dealController{
		dealService: dealService,
	}

	api.Get("/deals", controller.GetAll)
	api.Post("/deals", controller.Create)
	api.Put("/deals/:id", controller.Update)
	api.Delete("/deals/:id", controller.Delete)
}
