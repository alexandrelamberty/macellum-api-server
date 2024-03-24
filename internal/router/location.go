package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func LocationRouter(app fiber.Router, service service.LocationService) {
	app.Get("/locations", handlers.GetAllLocations(service))
	app.Post("/locations", handlers.CreateLocation(service))
	app.Get("/locations/:id", handlers.GetLocationByID(service))
	app.Patch("/locations/:id", handlers.UpdateLocation(service))
	app.Delete("/locations/:id", handlers.DeleteLocation(service))
}
