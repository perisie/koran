package mouse

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

var REGEX_KEY = regexp.MustCompile("([a-z|[0-9]|_)+")

type Mouse struct{}

func (m *Mouse) Put(key string, value []byte) error {
	key_n := m.normalize_key(key)
	if !m.is_key_ok(key_n) {
		return errors.New("invalid key")
	}
	f, err := os.Create(fmt.Sprintf("./data/%v.data", key_n))
	if err != nil {
		return err
	}
	_, err = f.Write(value)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *Mouse) Get(key string) ([]byte, error) {
	key_n := m.normalize_key(key)
	if !m.is_key_ok(key_n) {
		return nil, errors.New("invalid key")
	}
	f, err := os.Open(fmt.Sprintf("./data/%v.data", key_n))
	if err != nil {
		return nil, err
	}
	value, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (m *Mouse) is_key_ok(key string) bool {
	return REGEX_KEY.Match([]byte(key))
}

func (m *Mouse) normalize_key(key string) string {
	key_n := ""
	for _, k := range key {
		for _, ch := range "abcdefghijklmnopqrstuvwxyz0123456789_" {
			if k == ch {
				key_n += string(ch)
			}
		}
	}
	return key_n
}

func Mouse_new() *Mouse {
	_ = os.Mkdir("data", 0755)
	return &Mouse{}
}

func To_byte(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
