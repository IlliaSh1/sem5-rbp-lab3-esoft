package controllers_offer

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	repos_mysql_offer "github.com/IlliaSh1/backend/internal/repos/mysql/offer"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *offerController) Update(ctx fiber.Ctx) error {
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

	_, err = controller.offerService.GetByID(id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_offer.ErrOfferNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		default:
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	offer_to_update := models.Offer{
		ID:                 id,
		RealEstateObjectID: req.RealEstateObjectID,
		ClientID:           req.ClientID,
		RealtorID:          req.RealtorID,
		Price:              req.OfferPrice,
	}

	updated_offer, err := controller.offerService.Save(&offer_to_update)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(updated_offer)
}
