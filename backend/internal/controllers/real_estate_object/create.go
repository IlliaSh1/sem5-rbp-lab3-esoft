package controllers_real_estate_object

import (
	"context"
	"errors"
	"fmt"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *realEstateObjectController) Create(ctx fiber.Ctx) error {
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

	real_estate_object_to_create := models.RealEstateObject{
		RealEstateObjectTypeID: req.RealEstateObjectTypeID,
		Coordinates:            req.Coordinates,
		City:                   req.City,
		Street:                 req.Street,
		HouseNumber:            req.HouseNumber,
		ApartmentNumber:        req.ApartmentNumber,
	}

	var created_real_estate_object *models.RealEstateObject = nil
	var created_value interface{}

	err = controller.transactor.WithinTransaction(ctx.Context(), func(ctx context.Context) error {
		created_real_estate_object, err = controller.realEstateObjectService.Save(ctx, &real_estate_object_to_create)
		if err != nil {
			return err
		}

		switch req.RealEstateObjectTypeID {
		case int(models.KRealEstateObjectTypeApartmentID):
			apartment_to_create := &models.Apartment{
				RealEstateObjectID: created_real_estate_object.ID,
				Area:               req.ApartmentData.Area,
				RoomsCount:         req.ApartmentData.RoomsCount,
				FloorNumber:        req.ApartmentData.FloorNumber,
			}
			created_apartment, err := controller.realEstateObjectService.SaveApartment(ctx, apartment_to_create)
			if err != nil {
				return err
			}
			created_value = created_apartment
		case int(models.KRealEstateObjectTypeHouseID):
			house_to_create := &models.House{
				RealEstateObjectID: created_real_estate_object.ID,
				FloorsCount:        req.HouseData.FloorsCount,
				RoomsCount:         req.HouseData.RoomsCount,
				Area:               req.HouseData.Area,
			}
			created_house, err := controller.realEstateObjectService.SaveHouse(ctx, house_to_create)
			if err != nil {
				return err
			}
			created_value = created_house
		case int(models.KRealEstateObjectTypeLandID):
			land_to_create := &models.Land{
				RealEstateObjectID: created_real_estate_object.ID,
				Area:               req.LandData.Area,
			}
			created_land, err := controller.realEstateObjectService.SaveLand(ctx, land_to_create)
			if err != nil {
				return err
			}
			created_value = created_land
		default:
			return errors.New("invalid real estate object type id")
		}

		return nil
	})
	if err != nil {
		log.Error("within tx error: ", err)
		return nil
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"real_estate_object": created_real_estate_object,
		"value":              created_value,
	})
}

type CreateRequest struct {
	RealEstateObjectTypeID int `json:"real_estate_object_type_id"`

	ApartmentData *CreateApartmentData `json:"apartment_data"`
	HouseData     *CreateHouseData     `json:"house_data"`
	LandData      *CreateLandData      `json:"land_data"`

	Coordinates *models.Point `json:"coordinates"`

	City            *string `json:"city" gorm:"column:city"`
	Street          *string `json:"street" gorm:"column:street"`
	HouseNumber     *string `json:"house_number" gorm:"column:house_number"`
	ApartmentNumber *string `json:"apartment_number" gorm:"column:apartment_number"`
}

func (req *CreateRequest) Validate() (err error) {
	switch req.RealEstateObjectTypeID {
	case int(models.KRealEstateObjectTypeApartmentID):
		if req.ApartmentData == nil {
			err = errors.Join(err, fmt.Errorf("apartment_data is required"))
		}
	case int(models.KRealEstateObjectTypeHouseID):
		if req.HouseData == nil {
			err = errors.Join(err, fmt.Errorf("house_data is required"))
		}
	case int(models.KRealEstateObjectTypeLandID):
		if req.LandData == nil {
			err = errors.Join(err, fmt.Errorf("land_data is required"))
		}
	default:
		err = errors.Join(err, fmt.Errorf("real_estate_object_type_id is invalid, it must be 1, 2 or 3"))
	}

	return err
}

type CreateApartmentData struct {
	FloorNumber *int  `json:"floor_number"`
	RoomsCount  *uint `json:"rooms_count"`
	Area        *uint `json:"area"`
}

type CreateHouseData struct {
	FloorsCount *uint `json:"floors_count"`
	RoomsCount  *uint `json:"rooms_count"`
	Area        *uint `json:"area"`
}

type CreateLandData struct {
	Area *uint `json:"area"`
}
