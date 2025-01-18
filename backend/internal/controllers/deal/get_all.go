package controllers_deal

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *dealController) GetAll(ctx fiber.Ctx) error {
	deals, err := controller.dealService.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(deals)
}
