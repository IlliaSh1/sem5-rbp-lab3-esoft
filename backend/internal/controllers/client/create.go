package controllers_client

import (
	"errors"
	"fmt"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *ClientController) Create(ctx fiber.Ctx) error {
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

	client_to_create := models.Client{
		Email:      req.Email,
		Phone:      req.Phone,
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
	}

	created_client, err := controller.clientService.Save(&client_to_create)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(created_client)
}

type CreateRequest struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`

	Phone *string `json:"phone"`
	Email *string `json:"email"`
}

func (request *CreateRequest) Validate() (err error) {
	if request.Phone == nil && request.Email == nil {
		err = errors.Join(err, fmt.Errorf("phone or email is required"))
	}

	return err
}
