package requestresponse

type PostAuthGoogleRequest struct {
	AuthCode string `json:"auth_code"`
}

type PostAuthGoogleResponse struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Picture string `json:"picture"`
}
