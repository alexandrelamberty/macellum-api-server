package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllProviders(service service.ProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllProviders()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListProvidersSuccessResponse(result))
	}
}

func CreateProvider(service service.ProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var provider domain.Provider
		// Parse the request body
		err := c.BodyParser(&provider)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		// Create the provider
		err = service.CreateProvider(&provider)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		// Return the created provider
		return c.JSON(responses.GetProviderSuccessResponse(&provider))
	}
}

func GetProviderByID(service service.ProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the provider from the service based on the parsed 'id'.Verify if the provider exists.
		provider, err := service.GetProviderByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		return c.JSON(responses.GetProviderSuccessResponse(provider))
	}
}

func UpdateProvider(service service.ProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var provider domain.Provider
		// Parse the request body
		err := c.BodyParser(&provider)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		// Retrieve the provider from the service based on the parsed 'id'. Verify if the provider exists.
		_, err = service.GetProviderByID(provider.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		// Update the provider
		err = service.UpdateProvider(&provider)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		// Return the updated provider
		return c.JSON(responses.UpdateProviderSuccessResponse(&provider))
	}
}

func DeleteProvider(service service.ProviderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the provider
		err = service.DeleteProvider(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProviderErrorResponse(err.Error()))
		}

		// Return the deleted provider ID
		// FIXME: Return the deleted provider id
		return c.JSON(responses.DeleteProviderSuccessResponse())
	}
}
