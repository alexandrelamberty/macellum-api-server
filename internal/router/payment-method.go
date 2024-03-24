package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func PaymentMethodRouter(app fiber.Router, service service.PaymentMethodService) {
	app.Get("/payment-methods", handlers.GetAllPaymentMethods(service))
	app.Post("/payment-methods", handlers.CreatePaymentMethod(service))
	app.Get("/payment-methods/:id", handlers.GetPaymentMethodByID(service))
	app.Patch("/payment-methods/:id", handlers.UpdatePaymentMethod(service))
	app.Delete("/payment-methods/:id", handlers.DeletePaymentMethod(service))
}
