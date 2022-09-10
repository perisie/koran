package managers_test

import (
	"testing"

	"github.com/arikama/koran-backend/managers"
	"github.com/stretchr/testify/assert"
)

func TestQuranManagerErr(t *testing.T) {
	assert.Equal(t, "verse does not exist", managers.ErrVerseDoesNotExist())
	assert.Equal(t, "surah does not exist", managers.ErrSurahDoesNotExist())
}
