package controllers_needs

import (
	repos_mysql_need "github.com/IlliaSh1/backend/internal/repos/mysql/need"
	"github.com/IlliaSh1/backend/internal/storage/transactor"
	"github.com/gofiber/fiber/v3"
)

type needController struct {
	transactor *transactor.Transactor

	needService *repos_mysql_need.NeedRepo
}

func AddNeedController(api fiber.Router, transactor *transactor.Transactor, needService *repos_mysql_need.NeedRepo) {
	controller := needController{
		transactor:  transactor,
		needService: needService,
	}

	api.Get("/needs", controller.GetAll)
	api.Post("/needs", controller.Create)
	api.Put("/needs/:id", controller.Update)
	api.Delete("/needs/:id", controller.Delete)
}
