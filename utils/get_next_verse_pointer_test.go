package utils_test

import (
	"testing"

	"github.com/arikama/koran-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetNextVersePointer(t *testing.T) {
	assert.Equal(t, "1:2", utils.GetNextVersePointer("1:1"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer("114:6"))
}

func TestGetNextVersePointerCompleteQuran(t *testing.T) {
	pointer := "1:1"
	for i := 0; i < 6234; i++ {
		pointer = utils.GetNextVersePointer(pointer)
	}
	assert.Equal(t, "1:1", pointer)
}

func TestGetNextVersePointerInvalid(t *testing.T) {
	assert.Equal(t, "1:1", utils.GetNextVersePointer("1000:1000"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer("1:1000"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer("1000:1"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer("1:"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer(":1"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer(":"))
	assert.Equal(t, "1:1", utils.GetNextVersePointer(""))
}
