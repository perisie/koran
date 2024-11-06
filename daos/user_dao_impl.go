package daos

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/models"
	"github.com/arikama/koran-backend/mouse"
)

type UserDaoImpl struct {
	mouse *mouse.Mouse
}

func NewUserDaoImpl() (*UserDaoImpl, error) {
	return &UserDaoImpl{
		mouse: mouse.Mouse_new("./data"),
	}, nil
}

func (u *UserDaoImpl) CreateUser(email, token string) error {
	user := models.User{
		Email: email,
		Token: token,
	}
	key := u.key_user_email(email)
	key_token := u.key_user_token(token)
	value, err := mouse.To_byte(user)
	if err != nil {
		return err
	}
	err = u.mouse.Put(key, value)
	if err != nil {
		return err
	}
	return u.mouse.Put(key_token, value)
}

func (u *UserDaoImpl) QueryUserByEmail(email string) (*beans.User, error) {
	key := u.key_user_email(email)
	value, err := u.mouse.Get(key)
	if err != nil {
		return nil, err
	}
	decoder := gob.NewDecoder(bytes.NewReader(value))
	var user beans.User
	err = decoder.Decode(&user)
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
	key := u.key_user_token(token)
	value, err := u.mouse.Get(key)
	if err != nil {
		return nil, err
	}
	decoder := gob.NewDecoder(bytes.NewReader(value))
	var user beans.User
	err = decoder.Decode(&user)
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
	user, err := u.QueryUserByEmail(email)
	if err != nil {
		return err
	}
	user.Token = token
	value, err := mouse.To_byte(user)
	if err != nil {
		return err
	}
	key := u.key_user_email(email)
	key_token := u.key_user_token(token)
	err = u.mouse.Put(key, value)
	if err != nil {
		return err
	}
	return u.mouse.Put(key_token, value)
}

func (u *UserDaoImpl) UpdateUserCurrentPointer(email, currentPointer string) error {
	user, err := u.QueryUserByEmail(email)
	if err != nil {
		return err
	}
	user.CurrentPointer = currentPointer
	value, err := mouse.To_byte(user)
	if err != nil {
		return err
	}
	key := u.key_user_email(email)
	key_token := u.key_user_token(user.Token)
	err = u.mouse.Put(key, value)
	if err != nil {
		return err
	}
	return u.mouse.Put(key_token, value)
}

func (u *UserDaoImpl) key_user_email(email string) string {
	return fmt.Sprintf("user__%v", email)
}

func (u *UserDaoImpl) key_user_token(token string) string {
	return fmt.Sprintf("token__%v", token)
}
