package app

import controllers_client "github.com/IlliaSh1/backend/internal/controllers/client"

func (app *App) initControllers() error {
	api := app.fiberApp.Group("/api")

	controllers_client.AddClientControllerRoutes(api, app.serviceProvider.ClientRepo())

	return nil
}
