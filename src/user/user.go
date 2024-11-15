package user

type User struct {
	Username string
	Password string
}

func User_new(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
