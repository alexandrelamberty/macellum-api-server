package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllPaymentMethods(service service.PaymentMethodService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllPaymentMethods()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		return c.JSON(result)
	}
}

func CreatePaymentMethod(service service.PaymentMethodService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payementMethod domain.PaymentMethod
		// Parse the request body
		err := c.BodyParser(&payementMethod)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Create the payementMethod
		err = service.CreatePaymentMethod(&payementMethod)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the created payementMethod
		return c.JSON(&payementMethod)
	}
}

func GetPaymentMethodByID(service service.PaymentMethodService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the payementMethod from the service based on the parsed 'id'.Verify if the cart exists.
		category, err := service.GetPaymentMethodByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		return c.JSON(category)
	}
}

func UpdatePaymentMethod(service service.PaymentMethodService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payementMethod domain.PaymentMethod
		// Parse the request body
		err := c.BodyParser(&payementMethod)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Retrieve the payementMethod from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetPaymentMethodByID(payementMethod.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Update the payementMethod
		err = service.UpdatePaymentMethod(&payementMethod)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the updated payementMethod
		return c.JSON(&payementMethod)
	}
}

func DeletePaymentMethod(service service.PaymentMethodService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the payementMethod
		err = service.DeletePaymentMethod(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ErrorResponse(err.Error()))
		}

		// Return the deleted payementMethod ID
		// FIXME: Return the deleted payementMethod id
		return c.JSON(true)
	}
}
