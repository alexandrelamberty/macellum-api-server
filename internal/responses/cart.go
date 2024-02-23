package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListCartsSuccessResponse(data *[]domain.Cart) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Carts successfully listed",
		Data:    data,
	}
}

func CreateCartSuccessResponse(data *domain.Cart) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Cart successfully retrieved",
	}
}

func GetCartSuccessResponse(data *domain.Cart) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Cart successfully retrieved",
	}
}

func UpdateCartSuccessResponse(data *domain.Cart) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Cart successfully updated",
		Data:    data,
	}
}

func DeleteCartSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Cart successfully deleted",
		Data:    nil,
	}
}

func CartErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
