package managers

import "perisie.com/koran/beans"

//go:generate mockgen -package=managers -mock_names=UserManager=UserManagerMock -source=./user_manager.go -destination=./user_manager_mock.go
type UserManager interface {
	CreateUser(email, token string) (*beans.User, error)
	GetUser(token string) (*beans.User, error)
	AdvanceUserCurrentPointer(email, token string) (string, error)
	ReverseUserCurrentPointer(email, token string) (string, error)
}

func ErrUserTokenMismatch() string         { return "error user token mismatch" }
func ErrUserWithEmailDoesNotExist() string { return "user with email does not exist" }
