package user

import (
	"bytes"
	"encoding/json"
)

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

func User_new(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

func User_new_empty() *User {
	return &User{}
}
