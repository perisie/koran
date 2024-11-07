package managers

import (
	"errors"

	"github.com/hooligram/kifu"
	"perisie.com/koran/beans"
	"perisie.com/koran/constants"
	"perisie.com/koran/daos"
	"perisie.com/koran/utils"
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
		err := u.userDao.UpdateUserToken(email, token)
		if err != nil {
			kifu.Error("Error updating existing user token: %v", err.Error())
			return nil, err
		}
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
	err = u.userDao.UpdateUserCurrentPointer(user.Email, constants.StartPointer())
	if err != nil {
		kifu.Error("Error updating user current pointer: %v", err.Error())
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

func (u *UserManagerImpl) AdvanceUserCurrentPointer(email, token string) (string, error) {
	user, err := u.userDao.QueryUserByEmail(email)
	if err != nil {
		return "", err
	}
	if token != user.Token {
		return "", errors.New(ErrUserTokenMismatch())
	}
	newPointer := utils.GetNextVersePointer(user.CurrentPointer, 1)
	if newPointer == "" {
		return newPointer, errors.New("failed to advance user pointer")
	}
	err = u.userDao.UpdateUserCurrentPointer(email, newPointer)
	if err != nil {
		return "", err
	}
	return newPointer, nil
}

func (u *UserManagerImpl) ReverseUserCurrentPointer(email, token string) (string, error) {
	user, err := u.userDao.QueryUserByEmail(email)
	if err != nil {
		return "", err
	}
	if token != user.Token {
		return "", errors.New(ErrUserTokenMismatch())
	}
	newPointer := utils.GetNextVersePointer(user.CurrentPointer, -1)
	if newPointer == "" {
		return newPointer, errors.New("failed to reverse user pointer")
	}
	err = u.userDao.UpdateUserCurrentPointer(email, newPointer)
	if err != nil {
		return "", err
	}
	return newPointer, nil
}
