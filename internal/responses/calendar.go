package responses

import "github.com/alexandrelamberty/macellum-api-server/pkg/domain"

func ListCalendarsSuccessResponse(data *[]domain.Calendar) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Calendars successfully listed",
		Data:    data,
	}
}

func CreateCalendarSuccessResponse(data *domain.Calendar) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Calendar successfully retrieved",
	}
}

func GetCalendarSuccessResponse(data *domain.Calendar) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Calendar successfully retrieved",
	}
}

func UpdateCalendarSuccessResponse(data *domain.Calendar) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Calendar successfully updated",
		Data:    data,
	}
}

func DeleteCalendarSuccessResponse() ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: "Calendar successfully deleted",
		Data:    nil,
	}
}

func CalendarErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: "Error: " + msg,
		Data:    nil,
	}
}
