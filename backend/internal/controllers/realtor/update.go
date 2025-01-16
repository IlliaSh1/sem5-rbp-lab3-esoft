package controllers_realtor

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	repos_mysql_realtor "github.com/IlliaSh1/backend/internal/repos/mysql/realtor"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v3"
)

func (controller *realtorController) Update(ctx fiber.Ctx) error {
	id := fiber.Params[int](ctx, "id")
	if id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	req := &CreateRequest{}

	err := ctx.Bind().Body(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = req.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err = controller.realtorService.GetByID(id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_realtor.ErrRealtorNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		default:
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	realtor_to_update := models.Realtor{
		ID:                     id,
		Name:                   req.Name,
		Surname:                req.Surname,
		Patronymic:             req.Patronymic,
		PercentageOfCommission: req.PercentageOfCommission,
	}

	updated_realtor, err := controller.realtorService.Save(&realtor_to_update)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(updated_realtor)
}
