package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func OrderRouter(app fiber.Router, service service.OrderService) {
	app.Get("/orders", handlers.GetAllOrders(service))
	app.Post("/orders", handlers.CreateOrder(service))
	app.Get("/orders/:id", handlers.GetOrderByID(service))
	app.Patch("/orders/:id", handlers.UpdateOrder(service))
	app.Delete("/orders/:id", handlers.DeleteOrder(service))
}
