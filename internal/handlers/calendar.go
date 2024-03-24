package handlers

import (
	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllCalendars(service service.CalendarService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllCalendars()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListCalendarsSuccessResponse(result))
	}
}

func CreateCalendar(service service.CalendarService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var calendar domain.Calendar
		// Parse the request body
		err := c.BodyParser(&calendar)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		// Create the calendar
		err = service.CreateCalendar(&calendar)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		// Return the created calendar
		return c.JSON(responses.GetCalendarSuccessResponse(&calendar))
	}
}

func GetCalendarByID(service service.CalendarService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the calendar from the service based on the parsed 'id'.Verify if the calendar exists.
		calendar, err := service.GetCalendarByID(uintID)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		return c.JSON(responses.GetCalendarSuccessResponse(calendar))
	}
}

func UpdateCalendar(service service.CalendarService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var calendar domain.Calendar
		// Parse the request body
		err := c.BodyParser(&calendar)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		// Retrieve the calendar from the service based on the parsed 'id'. Verify if the calendar exists.
		_, err = service.GetCalendarByID(calendar.ID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		// Update the calendar
		err = service.UpdateCalendar(&calendar)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		// Return the updated calendar
		return c.JSON(responses.UpdateCalendarSuccessResponse(&calendar))
	}
}

func DeleteCalendar(service service.CalendarService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the calendar
		err = service.DeleteCalendar(uintID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responses.CalendarErrorResponse(err.Error()))
		}

		// Return the deleted calendar ID
		// FIXME: Return the deleted calendar id
		return c.JSON(responses.DeleteCalendarSuccessResponse())
	}
}
