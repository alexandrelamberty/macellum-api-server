package routes

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func ProviderRouter(app fiber.Router, service service.ProviderService) {
	app.Get("/providers", handlers.GetAllProviders(service))
	app.Post("/providers", handlers.CreateProvider(service))
	app.Get("/providers/:id", handlers.GetProviderByID(service))
	app.Patch("/providers/:id", handlers.UpdateProvider(service))
	app.Delete("/providers/:id", handlers.DeleteProvider(service))
}
