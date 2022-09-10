package daos

import (
	"github.com/arikama/koran-backend/beans"
)

//go:generate mockgen -package=daos -mock_names=UserDao=UserDaoMock -source=./user_dao.go -destination=./user_dao_mock.go
type UserDao interface {
	CreateUser(email, token string) error
	QueryUserByEmail(email string) (*beans.User, error)
	QueryUserByToken(token string) (*beans.User, error)
	UpdateUserToken(email, token string) error
}

func ErrSqlNoRowsInResultSet() string { return "sql: no rows in result set" }
