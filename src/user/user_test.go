package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_username_valid(t *testing.T) {
	usernames := []string{
		"athirah",
		"amir1",
		"imran_1",
	}
	for _, username := range usernames {
		user := User_new(username, "")
		assert.True(t, user.Ok_username(), fmt.Sprintf("username should be valid: %v", username))
	}
}

func Test_username_invalid(t *testing.T) {
	usernames := []string{
		"",
		" ",
		"A",
		"1",
		"aA",
	}
	for _, username := range usernames {
		user := User_new(username, "")
		assert.False(t, user.Ok_username(), fmt.Sprintf("username should be invalid: %v", username))
	}
}
