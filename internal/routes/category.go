package routes

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(app fiber.Router, service service.CategoryService) {
	app.Get("/categories", handlers.GetAllCategories(service))
	app.Post("/categories", handlers.CreateCategory(service))
	app.Get("/categories/:id", handlers.GetCategoryByID(service))
	app.Patch("/categories/:id", handlers.UpdateCategory(service))
	app.Delete("/categories/:id", handlers.DeleteCategory(service))
}
