package routes

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/handlers"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func CalendarRouter(app fiber.Router, service service.CalendarService) {
	app.Get("/calendars", handlers.GetAllCalendars(service))
	app.Post("/calendars", handlers.CreateCalendar(service))
	app.Get("/calendars/:id", handlers.GetCalendarByID(service))
	app.Post("/calendars/:id", handlers.UpdateCalendar(service))
	app.Delete("/calendars/:id", handlers.DeleteCalendar(service))
}
