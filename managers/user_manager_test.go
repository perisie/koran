package managers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"perisie.com/koran/managers"
)

func TestUserManagerErr(t *testing.T) {
	assert.Equal(t, "error user token mismatch", managers.ErrUserTokenMismatch())
}
