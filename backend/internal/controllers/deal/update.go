package controllers_deal

import (
	"errors"

	repos_mysql_deal "github.com/IlliaSh1/backend/internal/repos/mysql/deal"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *dealController) Update(ctx fiber.Ctx) error {
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

	deal, err := controller.dealService.GetByID(id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_deal.ErrDealNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	deal.OfferID = req.OfferID
	deal.NeedID = req.NeedID

	err = controller.dealService.Save(deal)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(deal)
}
