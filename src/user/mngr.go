package user

type Mngr interface {
	Create(username string, password string) (*User, error)
	Get(username string) (*User, error)
}
