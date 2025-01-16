package controllers_client

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	repos_mysql_client "github.com/IlliaSh1/backend/internal/repos/mysql/client"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *ClientController) Update(ctx fiber.Ctx) error {
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

	_, err = controller.clientService.GetById(id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_client.ErrClientNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	client_to_update := models.Client{
		ID:         id,
		Email:      req.Email,
		Phone:      req.Phone,
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
	}

	updated_client, err := controller.clientService.Save(&client_to_update)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(updated_client)
}
