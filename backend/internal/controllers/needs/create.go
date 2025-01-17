package controllers_needs

import (
	"context"
	"errors"
	"fmt"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *needController) Create(ctx fiber.Ctx) error {
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

	need := &models.Need{
		RealEstateObjectTypeID: req.RealEstateObjectTypeID,
		ClientID:               req.ClientID,
		RealtorID:              req.RealtorID,
		MinPrice:               req.MinPrice,
		MaxPrice:               req.MaxPrice,
	}
	var need_value interface{}

	err = controller.transactor.WithinTransaction(ctx.Context(), func(ctx context.Context) error {
		err = controller.needService.Save(need)
		if err != nil {
			return err
		}

		switch req.RealEstateObjectTypeID {
		case int(models.KRealEstateObjectTypeApartmentID):
			apartment_need := &models.ApartmentNeed{
				NeedID:         need.ID,
				MinFloorNumber: req.ApartmentData.MinFloorNumber,
				MaxFloorNumber: req.ApartmentData.MaxFloorNumber,
				MinRoomsCount:  req.ApartmentData.MinRoomsCount,
				MaxRoomsCount:  req.ApartmentData.MaxRoomsCount,
				MinArea:        req.ApartmentData.MinArea,
				MaxArea:        req.ApartmentData.MaxArea,
			}
			err = controller.needService.SaveApartmentNeed(apartment_need)
			if err != nil {
				return err
			}
			need_value = apartment_need
		case int(models.KRealEstateObjectTypeHouseID):
			house_need := &models.HouseNeed{
				NeedID:         need.ID,
				MinFloorsCount: req.HouseData.MinFloorsCount,
				MaxFloorsCount: req.HouseData.MaxFloorsCount,
				MinRoomsCount:  req.HouseData.MinRoomsCount,
				MaxRoomsCount:  req.HouseData.MaxRoomsCount,
				MinArea:        req.HouseData.MinArea,
				MaxArea:        req.HouseData.MaxArea,
			}
			err = controller.needService.SaveHouseNeed(house_need)
			if err != nil {
				return err
			}
			need_value = house_need
		case int(models.KRealEstateObjectTypeLandID):
			land_need := &models.LandNeed{
				NeedID:  need.ID,
				MinArea: req.LandData.MinArea,
				MaxArea: req.LandData.MaxArea,
			}
			err = controller.needService.SaveLandNeed(land_need)
			if err != nil {
				return err
			}
			need_value = land_need
		default:
			return errors.New("real_estate_object_type_id is invalid, it must be 1, 2 or 3")
		}

		return nil
	})
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"need":  need,
		"value": need_value,
	})
}

type CreateRequest struct {
	RealEstateObjectTypeID int `json:"real_estate_object_type_id"`

	ApartmentData *CreateApartmenNeedtData `json:"apartment_data"`
	HouseData     *CreateHouseNeedData     `json:"house_data"`
	LandData      *CreateLandNeedData      `json:"land_data"`

	ClientID  int `json:"client_id"`
	RealtorID int `json:"realtor_id"`

	MinPrice *uint `json:"min_price"`
	MaxPrice *uint `json:"max_price"`
}

func (req *CreateRequest) Validate() (err error) {
	switch req.RealEstateObjectTypeID {
	case int(models.KRealEstateObjectTypeApartmentID):
		if req.ApartmentData == nil {
			err = errors.Join(err, fmt.Errorf("apartment data is required"))
		}
	case int(models.KRealEstateObjectTypeHouseID):
		if req.HouseData == nil {
			err = errors.Join(err, fmt.Errorf("house data is required"))
		}
	case int(models.KRealEstateObjectTypeLandID):
		if req.LandData == nil {
			err = errors.Join(err, fmt.Errorf("land data is required"))
		}
	default:
		err = errors.Join(err, fmt.Errorf("real estate object type id is invalid, it must be 1, 2 or 3"))
	}

	if req.ClientID <= 0 {
		err = errors.Join(err, fmt.Errorf("client id is required"))
	}
	if req.RealtorID <= 0 {
		err = errors.Join(err, fmt.Errorf("realtor id is required"))
	}

	return err
}

type CreateApartmenNeedtData struct {
	MinFloorNumber *int  `json:"min_floor_number"`
	MaxFloorNumber *int  `json:"max_floor_number"`
	MinRoomsCount  *uint `json:"min_rooms_count"`
	MaxRoomsCount  *uint `json:"max_rooms_count"`
	MinArea        *uint `json:"min_area"`
	MaxArea        *uint `json:"max_area"`
}

type CreateHouseNeedData struct {
	MinFloorsCount *uint `json:"min_floors_count"`
	MaxFloorsCount *uint `json:"max_floors_count"`
	MinRoomsCount  *uint `json:"min_rooms_count"`
	MaxRoomsCount  *uint `json:"max_rooms_count"`
	MinArea        *uint `json:"min_area"`
	MaxArea        *uint `json:"max_area"`
}

type CreateLandNeedData struct {
	MinArea *uint `json:"min_area"`
	MaxArea *uint `json:"max_area"`
}
