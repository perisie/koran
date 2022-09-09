package daos

import (
	"context"
	"database/sql"

	"github.com/arikama/koran-backend/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserDaoImpl struct {
	context *context.Context
	db      *sql.DB
}

func NewUserDaoImpl(db *sql.DB) (*UserDaoImpl, error) {
	context := context.Background()
	return &UserDaoImpl{
		context: &context,
		db:      db,
	}, nil
}

func (u *UserDaoImpl) UpsertUser(user *models.User) (*models.User, error) {
	err := user.Upsert(*u.context, u.db, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDaoImpl) QueryUser(token string) (*models.User, error) {
	user, err := models.Users(qm.Where("token = ?", token)).One(*u.context, u.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
