package services

//go:generate mockgen -package=services -mock_names=GoogleAuthService=GoogleAuthServiceMock -source=./google_auth_service.go -destination=./google_auth_service_mock.go
type GoogleAuthService interface {
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
