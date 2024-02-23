package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

// GetCarts is a function to get all services from the database
func GetAllCarts(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllCarts()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListCartsSuccessResponse(result))
	}
}

func CreateCart(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Cart
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		// Create the cart
		err = service.CreateCart(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		// Return the created cart
		return c.JSON(responses.GetCartSuccessResponse(&cart))
	}
}

func GetCartByID(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the cart from the service based on the parsed 'id'.Verify if the cart exists.
		cart, err := service.GetCartByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		return c.JSON(responses.GetCartSuccessResponse(cart))
	}
}

func UpdateCart(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cart domain.Cart
		// Parse the request body
		err := c.BodyParser(&cart)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		// Retrieve the cart from the service based on the parsed 'id'. Verify if the cart exists.
		_, err = service.GetCartByID(cart.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		// Update the cart
		err = service.UpdateCart(&cart)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		// Return the updated cart
		return c.JSON(responses.UpdateCartSuccessResponse(&cart))
	}
}

func DeleteCart(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the cart
		err = service.DeleteCart(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CartErrorResponse(err.Error()))
		}

		// Return the deleted cart ID
		// FIXME: Return the deleted cart id
		return c.JSON(responses.DeleteCartSuccessResponse())
	}
}
