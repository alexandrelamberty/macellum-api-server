package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, service service.ProductService) {
	app.Get("/products", handlers.GetAllProducts(service))
	app.Post("/products", handlers.CreateProduct(service))
	app.Get("/products/:id", handlers.GetProductByID(service))
	app.Patch("/products/:id", handlers.UpdateProduct(service))
	app.Delete("/products/:id", handlers.DeleteProduct(service))
}
