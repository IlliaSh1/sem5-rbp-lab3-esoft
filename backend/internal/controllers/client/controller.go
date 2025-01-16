package controllers_client

import (
	repos_mysql_client "github.com/IlliaSh1/backend/internal/repos/mysql/client"
	"github.com/gofiber/fiber/v3"
)

type ClientController struct {
	clientService *repos_mysql_client.ClientRepo
}

func AddClientControllerRoutes(
	api fiber.Router,
	clientService *repos_mysql_client.ClientRepo,
) {
	controller := ClientController{
		clientService: clientService,
	}

	api.Get("/clients", controller.GetAll)
	api.Post("/clients", controller.Create)
	api.Put("/clients/:id", controller.Update)
}
