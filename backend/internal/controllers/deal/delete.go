package controllers_deal

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *dealController) Delete(ctx fiber.Ctx) error {
	id := fiber.Params[int](ctx, "id")
	if id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	err := controller.dealService.Delete(id)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
