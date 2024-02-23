package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListOrdersSuccessResponse(data *[]domain.Order) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Orders successfully listed",
		Data:    data,
	}
}

func CreateOrderSuccessResponse(data *domain.Order) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Order successfully retrieved",
	}
}

func GetOrderSuccessResponse(data *domain.Order) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Order successfully retrieved",
	}
}

func UpdateOrderSuccessResponse(data *domain.Order) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Order successfully updated",
		Data:    data,
	}
}

func DeleteOrderSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Order successfully deleted",
		Data:    nil,
	}
}

func OrderErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
