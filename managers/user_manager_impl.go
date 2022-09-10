package managers

import (
	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/daos"
	"github.com/hooligram/kifu"
)

type UserManagerImpl struct {
	userDao daos.UserDao
}

func NewUserManagerImpl(userDao daos.UserDao) (*UserManagerImpl, error) {
	return &UserManagerImpl{
		userDao: userDao,
	}, nil
}

func (u *UserManagerImpl) CreateUser(email, token string) (*beans.User, error) {
	existing, err := u.userDao.QueryUserByEmail(email)
	if err != nil && err.Error() != daos.ErrSqlNoRowsInResultSet() {
		kifu.Error("Error queriying user by email: %v", err.Error())
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}
	err = u.userDao.CreateUser(email, token)
	if err != nil {
		kifu.Error("Error creating user: %v", err.Error())
		return nil, err
	}
	user, err := u.userDao.QueryUserByToken(token)
	if err != nil {
		kifu.Error("Error queriying user by token: %v", err.Error())
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

func (u *UserManagerImpl) GetUser(token string) (*beans.User, error) {
	user, err := u.userDao.QueryUserByToken(token)
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
