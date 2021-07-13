package inmemory

type Request struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type Response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Message: err.Error()}
}
