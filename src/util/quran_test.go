package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_move_surah_verse(t *testing.T) {
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
		surahNext, verseNext, _ := Move_surah_verse(surah, verse, direction)
		assert.Equal(t, surahNextExpected, surahNext)
		assert.Equal(t, verseNextExpected, verseNext)
	}
}
