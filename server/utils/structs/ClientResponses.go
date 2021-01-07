package structs

type ClientResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewClientResponse(err bool, message string) *ClientResponse {
	return &ClientResponse{
		Error:   err,
		Message: message,
	}
}
