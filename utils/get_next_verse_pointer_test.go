package utils_test

import (
	"testing"

	"github.com/arikama/koran-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetNextVersePointer(t *testing.T) {
	assert.Equal(t, "1:2", utils.GetNextVersePointer("1:1", 1))
	assert.Equal(t, "1:1", utils.GetNextVersePointer("114:6", 1))
	assert.Equal(t, "114:6", utils.GetNextVersePointer("1:1", -1))
	assert.Equal(t, "113:5", utils.GetNextVersePointer("1:1", -1-6))
}

func TestGetNextVersePointerCompleteQuran(t *testing.T) {
	pointer := "1:1"
	for i := 0; i < 6234; i++ {
		pointer = utils.GetNextVersePointer(pointer, 1)
	}
	assert.Equal(t, "1:1", pointer)
}

func TestGetNextVersePointerInvalid(t *testing.T) {
	assert.Equal(t, "", utils.GetNextVersePointer("1000:1000", 1))
	assert.Equal(t, "", utils.GetNextVersePointer("1:1000", 1))
	assert.Equal(t, "", utils.GetNextVersePointer("1000:1", 1))
	assert.Equal(t, "", utils.GetNextVersePointer("1:", 1))
	assert.Equal(t, "", utils.GetNextVersePointer(":1", 1))
	assert.Equal(t, "", utils.GetNextVersePointer(":", 1))
	assert.Equal(t, "", utils.GetNextVersePointer("", 1))
}

func TestMoveSurahVerse(t *testing.T) {
	var tests = [][]int{
		{1, 1, 1, 1, 2},
		{1, 1, 7, 2, 1},
		{1, 1, 6 + 286 + 1, 3, 1},
		{1, 1, -1, 114, 6},
		{1, 1, -6, 114, 1},
		{1, 1, -6 - 1, 113, 5},
	}
	for _, test := range tests {
		surah := test[0]
		verse := test[1]
		direction := test[2]
		surahNextExpected := test[3]
		verseNextExpected := test[4]
		surahNext, verseNext, _ := utils.MoveSurahVerse(surah, verse, direction)
		assert.Equal(t, surahNextExpected, surahNext)
		assert.Equal(t, verseNextExpected, verseNext)
	}
}
