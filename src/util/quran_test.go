package util

import (
	"errors"
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
		{114, 6, 1, 1, 1},
	}
	for _, test := range tests {
		surah := test[0]
		verse := test[1]
		direction := test[2]
		expected_surah_next := test[3]
		expected_verse_next := test[4]

		surah_next, verse_next, _ := Move_surah_verse(surah, verse, direction)

		assert.Equal(t, expected_surah_next, surah_next)
		assert.Equal(t, expected_verse_next, verse_next)
	}
}

func Test_move_surah_verse_error(t *testing.T) {
	type Test struct {
		Surah     int
		Verse     int
		Direction int
		Error     error
	}
	var tests = []*Test{
		&Test{0, 1, 1, errors.New("invalid surah")},
		&Test{115, 1, 1, errors.New("invalid surah")},
		&Test{1, 0, 1, errors.New("invalid verse")},
		&Test{1, 8, 1, errors.New("invalid verse")},
	}
	for _, test := range tests {
		_, _, err := Move_surah_verse(test.Surah, test.Verse, test.Direction)

		assert.Equal(t, test.Error, err)
	}
}
