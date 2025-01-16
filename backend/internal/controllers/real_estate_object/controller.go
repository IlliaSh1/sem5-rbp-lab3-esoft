package controllers_real_estate_object

import (
	repos_mysql_real_estate_object "github.com/IlliaSh1/backend/internal/repos/mysql/real_estate_object"
	"github.com/IlliaSh1/backend/internal/storage/transactor"
	"github.com/gofiber/fiber/v3"
)

type realEstateObjectController struct {
	transactor              *transactor.Transactor
	realEstateObjectService *repos_mysql_real_estate_object.RealEstateObjectRepo
}

func AddRealEstateObjectController(
	api fiber.Router,
	transactor *transactor.Transactor,
	realEstateObjectService *repos_mysql_real_estate_object.RealEstateObjectRepo,
) {
	controller := realEstateObjectController{
		transactor:              transactor,
		realEstateObjectService: realEstateObjectService,
	}

	// api.Get("/real-estate-objects", controller.GetAll)
	api.Post("/real-estate-objects", controller.Create)
}
