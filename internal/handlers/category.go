package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(service service.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllCategories()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListCategoriesSuccessResponse(result))
	}
}

func CreateCategory(service service.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Category
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		// Create the cart
		err = service.CreateCategory(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		// Return the created cart
		return c.JSON(responses.GetCategorySuccessResponse(&cart))
	}
}

func GetCategoryByID(service service.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the cart from the service based on the parsed 'id'.Verify if the cart exists.
		category, err := service.GetCategoryByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		return c.JSON(category)
	}
}

func UpdateCategory(service service.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Category
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		// Retrieve the cart from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetCategoryByID(cart.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		// Update the cart
		err = service.UpdateCategory(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		// Return the updated cart
		return c.JSON(responses.UpdateCategorySuccessResponse(&cart))
	}
}

func DeleteCategory(service service.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the cart
		err = service.DeleteCategory(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CategoryErrorResponse(err.Error()))
		}

		// Return the deleted cart ID
		// FIXME: Return the deleted cart id
		return c.JSON(responses.DeleteCategorySuccessResponse())
	}
}
