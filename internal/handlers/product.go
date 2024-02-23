package handlers

import (
	"net/http"

	"github.com/alexandrelamberty/macellum-api-server/internal/responses"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

// GetProducts is a function to get all services from the database
func GetAllProducts(service service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllProducts()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListProductsSuccessResponse(result))
	}
}

func CreateProduct(service service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var product domain.Product
		// Parse the request body
		err := c.BodyParser(&product)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		// Create the product
		err = service.CreateProduct(&product)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		// Return the created product
		return c.JSON(responses.GetProductSuccessResponse(&product))
	}
}

func GetProductByID(service service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Retrieve the product from the service based on the parsed 'id'.Verify if the product exists.
		product, err := service.GetProductByID(uintID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		return c.JSON(responses.GetProductSuccessResponse(product))
	}
}

func UpdateProduct(service service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var product domain.Product
		// Parse the request body
		err := c.BodyParser(&product)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		// Retrieve the product from the service based on the parsed 'id'. Verify if the product exists.
		_, err = service.GetProductByID(product.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		// Update the product
		err = service.UpdateProduct(&product)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		// Return the updated product
		return c.JSON(responses.UpdateProductSuccessResponse(&product))
	}
}

func DeleteProduct(service service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the 'id' parameter
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}
		uintID := uint(id)

		// Delete the product
		err = service.DeleteProduct(uintID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ProductErrorResponse(err.Error()))
		}

		// Return the deleted product ID
		// FIXME: Return the deleted product id
		return c.JSON(responses.DeleteProductSuccessResponse())
	}
}
