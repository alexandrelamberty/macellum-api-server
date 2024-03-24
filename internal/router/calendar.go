package router

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func CalendarRouter(app fiber.Router, service service.CalendarService) {
	calendar := app.Group("/calendars")
	calendar.Get("/", handlers.GetAllCalendars(service))
	calendar.Post("/", handlers.CreateCalendar(service))
	calendar.Get("/:id", handlers.GetCalendarByID(service))
	calendar.Patch("/:id", handlers.UpdateCalendar(service))
	calendar.Delete("/:id", handlers.DeleteCalendar(service))
}
