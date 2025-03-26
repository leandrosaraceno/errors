package response

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func ReturnError(message string) Response {
	return Response{
		Success: false,
		Message: message,
	}
}
