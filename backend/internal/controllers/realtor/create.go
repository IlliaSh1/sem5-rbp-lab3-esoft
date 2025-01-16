package controllers_realtor

import (
	"errors"
	"fmt"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v3"
)

func (controller *realtorController) Create(ctx fiber.Ctx) error {
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

	realtor_to_create := models.Realtor{
		Name:                   req.Name,
		Surname:                req.Surname,
		Patronymic:             req.Patronymic,
		PercentageOfCommission: req.PercentageOfCommission,
	}

	created_realtor, err := controller.realtorService.Create(&realtor_to_create)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(created_realtor)
}

type CreateRequest struct {
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Patronymic             string `json:"patronymic"`
	PercentageOfCommission *uint  `json:"percentage_of_commission"`
}

func (request *CreateRequest) Validate() (err error) {
	if request.Name == "" {
		err = errors.Join(err, fmt.Errorf("name is required"))
	}
	if request.Surname == "" {
		err = errors.Join(err, fmt.Errorf("surname is required"))
	}
	if request.Patronymic == "" {
		err = errors.Join(err, fmt.Errorf("patronymic is required"))
	}
	if request.PercentageOfCommission != nil && *request.PercentageOfCommission < 0 || *request.PercentageOfCommission > 100 {
		err = errors.Join(err, fmt.Errorf("percentage of commission must be between 0 and 100"))
	}

	return err
}
