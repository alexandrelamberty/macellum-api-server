package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllIncomes(service service.IncomeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllIncomes()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		return c.JSON(result)
	}
}

func CreateIncome(service service.IncomeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Income
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Create the cart
		err = service.CreateIncome(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the created cart
		return c.JSON(&cart)
	}
}

func GetIncomeByID(service service.IncomeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the cart from the service based on the parsed 'id'.Verify if the cart exists.
		income, err := service.GetIncomeByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(income)
	}
}

func UpdateIncome(service service.IncomeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Income
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Retrieve the cart from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetIncomeByID(cart.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Update the cart
		err = service.UpdateIncome(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the updated cart
		return c.JSON(&cart)
	}
}

func DeleteIncome(service service.IncomeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the cart
		err = service.DeleteIncome(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the deleted cart ID
		// FIXME: Return the deleted cart id
		return c.JSON(true)
	}
}
