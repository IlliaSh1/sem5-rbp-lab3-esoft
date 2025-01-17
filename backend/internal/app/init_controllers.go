package app

import (
	controllers_client "github.com/IlliaSh1/backend/internal/controllers/client"
	controllers_real_estate_object "github.com/IlliaSh1/backend/internal/controllers/real_estate_object"
	controllers_realtor "github.com/IlliaSh1/backend/internal/controllers/realtor"
)

func (app *App) initControllers() error {
	api := app.fiberApp.Group("/api")

	controllers_client.AddClientControllerRoutes(api, app.serviceProvider.ClientRepo())

	controllers_realtor.AddRealtorController(api, app.serviceProvider.RealtorRepo())

	controllers_real_estate_object.AddRealEstateObjectController(
		api,
		app.serviceProvider.Transactor(),
		app.serviceProvider.RealEstateObjectRepo(),
	)
	return nil
}
