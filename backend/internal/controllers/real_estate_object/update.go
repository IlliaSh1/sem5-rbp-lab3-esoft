package controllers_real_estate_object

import (
	"context"
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	repos_mysql_real_estate_object "github.com/IlliaSh1/backend/internal/repos/mysql/real_estate_object"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *realEstateObjectController) Update(ctx fiber.Ctx) error {
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

	real_estate_object_to_update, err := controller.realEstateObjectService.GetById(id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, repos_mysql_real_estate_object.ErrRealEstateObjectNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if real_estate_object_to_update.RealEstateObjectTypeID != req.RealEstateObjectTypeID {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "real estate object type cannot be changed",
		})
	}

	real_estate_object_to_update.RealEstateObjectTypeID = req.RealEstateObjectTypeID
	real_estate_object_to_update.Latitude = req.Latitude
	real_estate_object_to_update.Longitude = req.Longitude
	real_estate_object_to_update.City = req.City
	real_estate_object_to_update.Street = req.Street
	real_estate_object_to_update.HouseNumber = req.HouseNumber
	real_estate_object_to_update.ApartmentNumber = req.ApartmentNumber

	var updated_real_estate_object *models.RealEstateObject = nil
	var updated_value interface{}

	err = controller.transactor.WithinTransaction(ctx.Context(), func(ctx context.Context) error {
		updated_real_estate_object, err = controller.realEstateObjectService.Save(ctx, real_estate_object_to_update)
		if err != nil {
			return err
		}

		switch req.RealEstateObjectTypeID {
		case int(models.KRealEstateObjectTypeApartmentID):
			apartment_to_update := &models.Apartment{
				RealEstateObjectID: updated_real_estate_object.ID,
				Area:               req.ApartmentData.Area,
				RoomsCount:         req.ApartmentData.RoomsCount,
				FloorNumber:        req.ApartmentData.FloorNumber,
			}
			created_apartment, err := controller.realEstateObjectService.SaveApartment(ctx, apartment_to_update)
			if err != nil {
				return err
			}
			updated_value = created_apartment
		case int(models.KRealEstateObjectTypeHouseID):
			house_to_update := &models.House{
				RealEstateObjectID: updated_real_estate_object.ID,
				FloorsCount:        req.HouseData.FloorsCount,
				RoomsCount:         req.HouseData.RoomsCount,
				Area:               req.HouseData.Area,
			}
			created_house, err := controller.realEstateObjectService.SaveHouse(ctx, house_to_update)
			if err != nil {
				return err
			}
			updated_value = created_house
		case int(models.KRealEstateObjectTypeLandID):
			land_to_update := &models.Land{
				RealEstateObjectID: updated_real_estate_object.ID,
				Area:               req.LandData.Area,
			}
			created_land, err := controller.realEstateObjectService.SaveLand(ctx, land_to_update)
			if err != nil {
				return err
			}
			updated_value = created_land
		default:
			return errors.New("invalid real estate object type id")
		}

		return nil
	})
	if err != nil {
		log.Error("within tx error: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"real_estate_object": updated_real_estate_object,
		"value":              updated_value,
	})
}
