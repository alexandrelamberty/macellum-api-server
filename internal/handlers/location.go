package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllLocations(service service.LocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllLocations()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		return c.JSON(result)
	}
}

func CreateLocation(service service.LocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var location domain.Location
		// Parse the request body
		err := c.BodyParser(&location)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Create the location
		err = service.CreateLocation(&location)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the created location
		return c.JSON(&location)
	}
}

func GetLocationByID(service service.LocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the location from the service based on the parsed 'id'.Verify if the cart exists.
		category, err := service.GetLocationByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(category)
	}
}

func UpdateLocation(service service.LocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var location domain.Location
		// Parse the request body
		err := c.BodyParser(&location)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Retrieve the location from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetLocationByID(location.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Update the location
		err = service.UpdateLocation(&location)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the updated location
		return c.JSON(&location)
	}
}

func DeleteLocation(service service.LocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the location
		err = service.DeleteLocation(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the deleted location ID
		// FIXME: Return the deleted location id
		return c.JSON(true)
	}
}
