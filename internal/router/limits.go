package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func LimitRouter(app fiber.Router, service service.LimitService) {
	app.Get("/limits", handlers.GetAllLimits(service))
	app.Post("/limits", handlers.CreateLimit(service))
	app.Get("/limits/:id", handlers.GetLimitByID(service))
	app.Patch("/limits/:id", handlers.UpdateLimit(service))
	app.Delete("/limits/:id", handlers.DeleteLimit(service))
}
