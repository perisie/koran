package managers_test

import (
	"testing"

	"github.com/arikama/koran-backend/managers"
	"github.com/stretchr/testify/assert"
)

func TestGetVerse(t *testing.T) {
	var quranManager managers.QuranManager
	var err error

	quranManager, err = managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	s1v1, err := quranManager.GetVerse(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ", s1v1.Text)

	s114v6, err := quranManager.GetVerse(114, 6)
	assert.Nil(t, err)
	assert.Equal(t, "مِنَ الْجِنَّةِ وَالنَّاسِ", s114v6.Text)
}
