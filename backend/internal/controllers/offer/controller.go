package controllers_offer

import (
	repos_mysql_offer "github.com/IlliaSh1/backend/internal/repos/mysql/offer"
	"github.com/gofiber/fiber/v3"
)

type offerController struct {
	offerService *repos_mysql_offer.OfferRepo
}

func AddOfferControllerRoutes(api fiber.Router, offerService *repos_mysql_offer.OfferRepo) {
	controller := offerController{
		offerService: offerService,
	}

	api.Get("/offers", controller.GetAll)
	api.Post("/offers", controller.Create)
	api.Put("/offers/:id", controller.Update)
	api.Delete("/offers/:id", controller.Delete)
}
