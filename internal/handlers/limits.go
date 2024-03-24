package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllLimits(service service.LimitService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllLimits()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		return c.JSON(result)
	}
}

func CreateLimit(service service.LimitService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Limit
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Create the cart
		err = service.CreateLimit(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the created cart
		return c.JSON(&cart)
	}
}

func GetLimitByID(service service.LimitService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the cart from the service based on the parsed 'id'.Verify if the cart exists.
		category, err := service.GetLimitByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(category)
	}
}

func UpdateLimit(service service.LimitService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Limit
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Retrieve the cart from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetLimitByID(cart.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Update the cart
		err = service.UpdateLimit(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the updated cart
		return c.JSON(&cart)
	}
}

func DeleteLimit(service service.LimitService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the cart
		err = service.DeleteLimit(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(true)
	}
}
