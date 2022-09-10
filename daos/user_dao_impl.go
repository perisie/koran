package daos

import (
	"context"
	"database/sql"

	"github.com/arikama/koran-backend/beans"
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

func (u *UserDaoImpl) CreateUser(email, token string) error {
	user := models.User{
		Email: email,
		Token: token,
	}
	err := user.Insert(*u.context, u.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDaoImpl) QueryUserByEmail(email string) (*beans.User, error) {
	user, err := models.Users(qm.Where("email = ?", email)).One(*u.context, u.db)
	if err != nil {
		return nil, err
	}
	return &beans.User{
		Email:          user.Email,
		Name:           user.Name,
		Token:          user.Token,
		Picture:        user.Picture,
		CurrentPointer: user.CurrentPointer,
	}, nil
}

func (u *UserDaoImpl) QueryUserByToken(token string) (*beans.User, error) {
	user, err := models.Users(qm.Where("token = ?", token)).One(*u.context, u.db)
	if err != nil {
		return nil, err
	}
	return &beans.User{
		Email:          user.Email,
		Name:           user.Name,
		Token:          user.Token,
		Picture:        user.Picture,
		CurrentPointer: user.CurrentPointer,
	}, nil
}

func (u *UserDaoImpl) UpdateUserToken(email, token string) error {
	user, err := models.Users(qm.Where("email = ?", email)).One(*u.context, u.db)
	if err != nil {
		return err
	}
	user.Token = token
	_, err = user.Update(*u.context, u.db, boil.Infer())
	return err
}

func (u *UserDaoImpl) UpdateUserCurrentPointer(email, currentPointer string) error {
	user, err := models.Users(qm.Where("email = ?", email)).One(*u.context, u.db)
	if err != nil {
		return err
	}
	user.CurrentPointer = currentPointer
	_, err = user.Update(*u.context, u.db, boil.Infer())
	return err
}
