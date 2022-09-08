package managers

type GoogleAuthManager interface {
	AuthUserCode(userAuthCode string) (*GoogleUser, error)
}

type GoogleUser struct {
	Email   string
	Name    string
	Token   string
	Picture string
}

type ClientId string
type ClientSecret string
type RedirectUrl string
