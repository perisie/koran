package requestresponse

type PatchUserPointerAdvanceRequest struct {
	Email string `json:"email"`
}

type PatchUserPointerAdvanceResponse struct {
	CurrentPointer string `json:"current_pointer"`
}
