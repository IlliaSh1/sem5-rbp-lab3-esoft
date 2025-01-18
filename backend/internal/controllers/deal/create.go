package controllers_deal

import (
	"errors"
	"fmt"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *dealController) Create(ctx fiber.Ctx) error {

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

	deal := models.Deal{
		OfferID: req.OfferID,
		NeedID:  req.NeedID,
	}

	err = controller.dealService.Save(&deal)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(deal)
}

type CreateRequest struct {
	OfferID int `json:"offer_id"`
	NeedID  int `json:"need_id"`
}

func (req *CreateRequest) Validate() (err error) {
	if req.OfferID <= 0 {
		err = errors.Join(err, fmt.Errorf("offer id is required"))
	}
	if req.NeedID <= 0 {
		err = errors.Join(err, fmt.Errorf("need id is required"))
	}
	return err
}
