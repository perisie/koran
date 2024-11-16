package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	mngr := Mngr_impl_new()
	username := "faithes"
	password := "password"

	user, err := mngr.Create(username, password)

	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	assert.NotEqual(t, password, user.Password)

	user, err = mngr.Get(username)

	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	assert.NotEqual(t, password, user.Password)
}
