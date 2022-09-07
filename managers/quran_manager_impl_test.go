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
	assert.Equal(t, 1, len(s1v1.Translations))
	assert.Equal(t, "In the name of Allah, the Beneficent, the Merciful.", s1v1.Translations["pickthall"])

	s114v6, err := quranManager.GetVerse(114, 6)
	assert.Nil(t, err)
	assert.Equal(t, "مِنَ الْجِنَّةِ وَالنَّاسِ", s114v6.Text)
	assert.Equal(t, 1, len(s114v6.Translations))
	assert.Equal(t, "Of the jinn and of mankind.", s114v6.Translations["pickthall"])
}

func TestGetSurah(t *testing.T) {
	var quranManager managers.QuranManager
	var err error

	quranManager, err = managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	surah, err := quranManager.GetSurah(114)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(surah.Verses))
	assert.Equal(t, "مِنَ الْجِنَّةِ وَالنَّاسِ", surah.Verses[5].Text)
}

func TestGetSurahNotExist(t *testing.T) {
	var quranManager managers.QuranManager
	var err error

	quranManager, err = managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	_, err = quranManager.GetSurah(115)
	assert.NotNil(t, err)
}

func TestGetSurahInfos(t *testing.T) {
	var quranManager managers.QuranManager
	var err error

	quranManager, err = managers.NewQuranManagerImpl("./../qurancsv")
	assert.Nil(t, err)

	surahInfos, err := quranManager.GetSurahInfos()
	assert.Nil(t, err)
	assert.Equal(t, 114, len(surahInfos))
	assert.Equal(t, "The Opening", surahInfos[0].English)
	assert.Equal(t, "Mankind", surahInfos[113].English)
}
