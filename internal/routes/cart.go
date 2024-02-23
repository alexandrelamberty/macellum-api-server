package routes

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func CartRouter(app fiber.Router, service service.CartService) {
	app.Get("/carts", handlers.GetAllCarts(service))
	app.Post("/carts", handlers.CreateCart(service))
	app.Get("/carts/:id", handlers.GetCartByID(service))
	app.Patch("/carts/:id", handlers.UpdateCart(service))
	app.Delete("/carts/:id", handlers.DeleteCart(service))
}
