package beans

type User struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Picture string `json:"picture"`
}
