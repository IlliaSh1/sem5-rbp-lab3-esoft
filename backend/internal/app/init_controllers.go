package app

import (
	controllers_client "github.com/IlliaSh1/backend/internal/controllers/client"
	controllers_realtor "github.com/IlliaSh1/backend/internal/controllers/realtor"
)

func (app *App) initControllers() error {
	api := app.fiberApp.Group("/api")

	controllers_client.AddClientControllerRoutes(api, app.serviceProvider.ClientRepo())

	controllers_realtor.AddRealtorController(api, app.serviceProvider.RealtorRepo())

	return nil
}
