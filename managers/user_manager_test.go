package managers_test

import (
	"testing"

	"github.com/arikama/koran-backend/managers"
	"github.com/stretchr/testify/assert"
)

func TestUserManagerErr(t *testing.T) {
	assert.Equal(t, "error user token mismatch", managers.ErrUserTokenMismatch())
}
