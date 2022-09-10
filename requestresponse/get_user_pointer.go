package requestresponse

type GetUserPointerRequest struct {
	Email string `json:"email"`
}

type GetUserPointerResponse struct {
	CurrentPointer string `json:"current_pointer"`
}
