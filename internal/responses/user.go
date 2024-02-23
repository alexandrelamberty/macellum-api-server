package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListUsersSuccessResponse(data *[]domain.User) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Users successfully listed",
		Data:    data,
	}
}

func CreateUserSuccessResponse(data *domain.User) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "User successfully retrieved",
	}
}

func GetUserSuccessResponse(data *domain.User) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "User successfully retrieved",
	}
}

func UpdateUserSuccessResponse(data *domain.User) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "User successfully updated",
		Data:    data,
	}
}

func DeleteUserSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "User successfully deleted",
		Data:    nil,
	}
}

func UserErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
