package user

import (
	"errors"

	pkg_mouse "github.com/perisie/mouse"
	"perisie.com/koran/src/util"
)

type Mngr_impl struct {
	mouse pkg_mouse.Mouse
}

func (m *Mngr_impl) Create(username string, password string) (*User, error) {
	password_hash, _ := util.Hash_password(password)
	user := User_new(username, password_hash)
	if !user.Ok_username() {
		return nil, errors.New("username invalid")
	}
	user_b, err := user.Ser()
	if err != nil {
		return nil, err
	}
	key := m.get_key(username)
	err = m.mouse.Put(key, user_b)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Mngr_impl) Get(username string) (*User, error) {
	user := User_new_empty()
	key := m.get_key(username)
	b, err := m.mouse.Get(key)
	if err != nil {
		return user, err
	}
	err = user.De(b)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *Mngr_impl) Update_surah_verse(username string, surah, verse int) error {
	user, err := m.Get(username)
	if err != nil {
		return err
	}
	user.Surah = surah
	user.Verse = verse
	user_ser, err := user.Ser()
	if err != nil {
		return err
	}
	key := m.get_key(username)
	err = m.mouse.Put(key, user_ser)
	return err
}

func (m *Mngr_impl) Update_setting(username string, name string, value string) error {
	user, err := m.Get(username)
	if err != nil {
		return err
	}
	if name == "surah_verse" {
		if value == "true" {
			user.Setting.Surah_verse = true
		} else {
			user.Setting.Surah_verse = false
		}
	}
	if name == "surah_translation" {
		if value == "true" {
			user.Setting.Surah_translation = true
		} else {
			user.Setting.Surah_translation = false
		}
	}
	if name == "bookmark_verse" {
		if value == "true" {
			user.Setting.Bookmark_verse = true
		} else {
			user.Setting.Bookmark_verse = false
		}
	}
	if name == "bookmark_translation" {
		if value == "true" {
			user.Setting.Bookmark_translation = true
		} else {
			user.Setting.Bookmark_translation = false
		}
	}
	user_ser, err := user.Ser()
	if err != nil {
		return err
	}
	key := m.get_key(username)
	err = m.mouse.Put(key, user_ser)
	return err
}

func (m *Mngr_impl) get_key(username string) string {
	return "v1__user__" + username
}

func Mngr_impl_new(dir_data string) *Mngr_impl {
	return &Mngr_impl{
		mouse: pkg_mouse.Mouse_fs_new(dir_data),
	}
}

func Mngr_impl_fake() *Mngr_impl {
	return &Mngr_impl{
		mouse: pkg_mouse.Mouse_memory_new(),
	}
}
