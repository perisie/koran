package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hash_password(t *testing.T) {
	password := "super_strong_password"
	hash, _ := Hash_password(password)
	assert.True(t, Hash_password_check(password, hash))
}
