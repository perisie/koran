package user

import (
	"bytes"
	"encoding/json"
	"regexp"
)

var REGEX_USERNAME = regexp.MustCompile(`^[a-z]([a-z]|[0-9]|_)+$`)

type User struct {
	Username string
	Password string
}

func (u *User) Ser() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(u)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (u *User) De(b []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Ok_username() bool {
	return REGEX_USERNAME.MatchString(u.Username)
}

func User_new(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

func User_new_empty() *User {
	return &User{}
}
