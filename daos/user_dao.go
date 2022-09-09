package daos

import "github.com/arikama/koran-backend/models"

//go:generate mockgen -package=daos -mock_names=UserDao=UserDaoMock -source=./user_dao.go -destination=./user_dao_mock.go
type UserDao interface {
	UpsertUser(user *models.User) (*models.User, error)
	QueryUser(token string) (*models.User, error)
}
