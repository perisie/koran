package user

import (
	"bytes"
	"errors"
	"github.com/perisie/mouse"
	"golang.org/x/crypto/bcrypt"
)

type Mngr_impl struct {
	mouse_fs *mouse.Mouse_fs
}

func (m *Mngr_impl) Create(username string, password string) (*User, error) {
	b, _ := bcrypt.GenerateFromPassword(bytes.NewBufferString(password).Bytes(), bcrypt.MinCost)
	password_hash := string(b)
	user := User_new(username, password_hash)
	if !user.Ok_username() {
		return nil, errors.New("username invalid")
	}
	user_b, err := user.Ser()
	if err != nil {
		return nil, err
	}
	err = m.mouse_fs.Put(username, user_b)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Mngr_impl) Get(username string) (*User, error) {
	user := User_new_empty()
	b, err := m.mouse_fs.Get(username)
	if err != nil {
		return user, err
	}
	err = user.De(b)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Mngr_impl_new() *Mngr_impl {
	return &Mngr_impl{
		mouse_fs: mouse.Mouse_fs_new("data"),
	}
}
