package controllers_real_estate_object

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *realEstateObjectController) GetAll(ctx fiber.Ctx) error {
	real_estate_objects, err := controller.realEstateObjectService.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(real_estate_objects)
}
