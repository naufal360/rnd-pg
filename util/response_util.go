package util

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewSuccessResponse(data interface{}) Response {
	return Response{Success: true, Data: data}
}

func NewErrorResponse(message string) Response {
	return Response{Success: false, Error: message}
}
