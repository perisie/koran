package managers

//go:generate mockgen -package=managers -mock_names=GoogleAuthManager=GoogleAuthManagerMock -source=./google_auth_manager.go -destination=./google_auth_manager_mock.go
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
