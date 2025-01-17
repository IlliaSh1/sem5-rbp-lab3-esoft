package controllers_needs

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *needController) GetAll(ctx fiber.Ctx) error {
	needs, err := controller.needService.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(needs)
}
