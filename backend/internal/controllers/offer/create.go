package controllers_offer

import (
	"errors"
	"fmt"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *offerController) Create(ctx fiber.Ctx) error {
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

	offer_to_create := models.Offer{
		RealEstateObjectID: req.RealEstateObjectID,
		ClientID:           req.ClientID,
		RealtorID:          req.RealtorID,
		Price:              req.OfferPrice,
	}

	created_offer, err := controller.offerService.Save(&offer_to_create)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(created_offer)
}

type CreateRequest struct {
	RealEstateObjectID int `json:"real_estate_object_id"`
	ClientID           int `json:"client_id"`
	RealtorID          int `json:"realtor_id"`
	OfferPrice         int `json:"offer_price"`
}

func (request *CreateRequest) Validate() (err error) {
	if request.RealEstateObjectID <= 0 {
		err = errors.Join(err, fmt.Errorf("real estate object id is required"))
	}
	if request.ClientID <= 0 {
		err = errors.Join(err, fmt.Errorf("client id is required"))
	}
	if request.RealtorID <= 0 {
		err = errors.Join(err, fmt.Errorf("realtor id is required"))
	}
	if request.OfferPrice <= 0 {
		err = errors.Join(err, fmt.Errorf("offer price is required and must be greater than 0"))
	}

	return err
}
