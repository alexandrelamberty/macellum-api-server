package routes

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service service.AuthService) {
	calendar := app.Group("/calendars")
	calendar.Post("/login", handlers.CreateCalendar(service))

}
