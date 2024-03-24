package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(service service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllOrders()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListOrdersSuccessResponse(result))
	}
}

func CreateOrder(service service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var order domain.Order
		// Parse the request body
		err := c.BodyParser(&order)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		// Create the order
		err = service.CreateOrder(&order)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		// Return the created order
		return c.JSON(responses.GetOrderSuccessResponse(&order))
	}
}

func GetOrderByID(service service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the order from the service based on the parsed 'id'.Verify if the order exists.
		order, err := service.GetOrderByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		return c.JSON(responses.GetOrderSuccessResponse(order))
	}
}

func UpdateOrder(service service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var order domain.Order
		// Parse the request body
		err := c.BodyParser(&order)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		// Retrieve the order from the service based on the parsed 'id'. Verify if the order exists.
		_, err = service.GetOrderByID(order.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		// Update the order
		err = service.UpdateOrder(&order)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		// Return the updated order
		return c.JSON(responses.UpdateOrderSuccessResponse(&order))
	}
}

func DeleteOrder(service service.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the order
		err = service.DeleteOrder(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.OrderErrorResponse(err.Error()))
		}

		// Return the deleted order ID
		// FIXME: Return the deleted order id
		return c.JSON(responses.DeleteOrderSuccessResponse())
	}
}
