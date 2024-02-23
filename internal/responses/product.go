package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListProductsSuccessResponse(data *[]domain.Product) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Products successfully listed",
		Data:    data,
	}
}

func CreateProductSuccessResponse(data *domain.Product) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Product successfully retrieved",
	}
}

func GetProductSuccessResponse(data *domain.Product) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Product successfully retrieved",
	}
}

func UpdateProductSuccessResponse(data *domain.Product) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Product successfully updated",
		Data:    data,
	}
}

func DeleteProductSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Product successfully deleted",
		Data:    nil,
	}
}

func ProductErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
