package daos

import "github.com/arikama/koran-backend/models"

type UserDao interface {
	UpsertUser(user models.User) (*models.User, error)
}
