package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func CustomerRouter(app fiber.Router, service service.CustomerService) {
	app.Get("/customers", handlers.GetAllCustomers(service))
	app.Post("/customers", handlers.CreateCustomer(service))
	app.Get("/customers/:id", handlers.GetCustomerByID(service))
	app.Patch("/customers/:id", handlers.UpdateCustomer(service))
	app.Delete("/customers/:id", handlers.DeleteCustomer(service))
}
