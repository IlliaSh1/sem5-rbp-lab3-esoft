package controllers_needs

import (
	"context"
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	repos_mysql_need "github.com/IlliaSh1/backend/internal/repos/mysql/need"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *needController) Update(ctx fiber.Ctx) error {
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

	need, err := controller.needService.GetByID(id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_need.ErrNeedNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		default:
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	if need.RealEstateObjectTypeID != req.RealEstateObjectTypeID {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "you can't change real estate object type id",
		})
	}

	need.ClientID = req.ClientID
	need.RealtorID = req.RealtorID
	need.MinPrice = req.MinPrice
	need.MaxPrice = req.MaxPrice

	var value interface{}

	err = controller.transactor.WithinTransaction(ctx.Context(), func(ctx context.Context) error {
		err = controller.needService.Save(need)
		if err != nil {
			return err
		}

		switch req.RealEstateObjectTypeID {
		case int(models.KRealEstateObjectTypeApartmentID):
			apartment_need := &models.ApartmentNeed{
				NeedID:         id,
				MinFloorNumber: req.ApartmentData.MinFloorNumber,
				MaxFloorNumber: req.ApartmentData.MaxFloorNumber,
				MinRoomsCount:  req.ApartmentData.MinRoomsCount,
				MaxRoomsCount:  req.ApartmentData.MaxRoomsCount,
				MinArea:        req.ApartmentData.MinArea,
				MaxArea:        req.ApartmentData.MaxArea,
			}
			err := controller.needService.SaveApartmentNeed(apartment_need)
			if err != nil {
				return err
			}
			value = apartment_need
		case int(models.KRealEstateObjectTypeHouseID):
			house_need := &models.HouseNeed{
				NeedID:         id,
				MinFloorsCount: req.HouseData.MinFloorsCount,
				MaxFloorsCount: req.HouseData.MaxFloorsCount,
				MinRoomsCount:  req.HouseData.MinRoomsCount,
				MaxRoomsCount:  req.HouseData.MaxRoomsCount,
				MinArea:        req.HouseData.MinArea,
				MaxArea:        req.HouseData.MaxArea,
			}
			err := controller.needService.SaveHouseNeed(house_need)
			if err != nil {
				return err
			}
			value = house_need
		case int(models.KRealEstateObjectTypeLandID):
			land_need := &models.LandNeed{
				NeedID:  id,
				MinArea: req.LandData.MinArea,
				MaxArea: req.LandData.MaxArea,
			}
			err := controller.needService.SaveLandNeed(land_need)
			if err != nil {
				return err
			}
			value = land_need
		}

		return nil
	})
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"need":  need,
		"value": value,
	})
	// err = controller.transactor.WithinTransaction(ctx.Context(), func(ctx context.Context) error {
	//     err = controller.needService.Save(id, req)
	//     if err != nil {
	//         return err
	//     }
}
