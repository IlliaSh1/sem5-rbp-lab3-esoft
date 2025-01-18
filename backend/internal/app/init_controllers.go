package app

import (
	controllers_client "github.com/IlliaSh1/backend/internal/controllers/client"
	controllers_deal "github.com/IlliaSh1/backend/internal/controllers/deal"
	controllers_needs "github.com/IlliaSh1/backend/internal/controllers/needs"
	controllers_offer "github.com/IlliaSh1/backend/internal/controllers/offer"
	controllers_real_estate_object "github.com/IlliaSh1/backend/internal/controllers/real_estate_object"
	controllers_realtor "github.com/IlliaSh1/backend/internal/controllers/realtor"
)

func (app *App) initControllers() error {
	api := app.fiberApp.Group("/api")

	controllers_client.AddClientControllerRoutes(api, app.serviceProvider.ClientRepo())

	controllers_realtor.AddRealtorControllerRoutes(api, app.serviceProvider.RealtorRepo())

	controllers_real_estate_object.AddRealEstateObjectController(
		api,
		app.serviceProvider.Transactor(),
		app.serviceProvider.RealEstateObjectRepo(),
	)

	controllers_offer.AddOfferControllerRoutes(api, app.serviceProvider.OfferRepo())

	controllers_needs.AddNeedController(api, app.serviceProvider.Transactor(), app.serviceProvider.NeedRepo())

	controllers_deal.AddDealControllerRoutes(api, app.serviceProvider.DealRepo())

	return nil
}
