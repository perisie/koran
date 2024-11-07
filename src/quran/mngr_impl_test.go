package quran

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_get_verse(t *testing.T) {
	var mngr Mngr

	mngr, err := Mngr_impl_new("../../qurancsv")
	assert.Nil(t, err)

	s1v1, err := mngr.Get_verse(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ", s1v1.Text)
	assert.Equal(t, 2, len(s1v1.Translations))
	assert.Equal(t, "In the name of Allah, the Beneficent, the Merciful.", s1v1.Translations["pickthall"])
	assert.Equal(t, "In the name of God, the Gracious, the Merciful.", s1v1.Translations["clearquran"])

	s114v6, err := mngr.Get_verse(114, 6)
	assert.Nil(t, err)
	assert.Equal(t, "مِنَ الْجِنَّةِ وَالنَّاسِ", s114v6.Text)
	assert.Equal(t, 2, len(s114v6.Translations))
	assert.Equal(t, "Of the jinn and of mankind.", s114v6.Translations["pickthall"])
	assert.Equal(t, "From among jinn and among people.”", s114v6.Translations["clearquran"])
}

func Test_get_verse_not_exist(t *testing.T) {
	var mngr Mngr

	mngr, err := Mngr_impl_new("../../qurancsv")
	assert.Nil(t, err)

	_, err = mngr.Get_verse(114, 7)
	assert.NotNil(t, err)
}

func Test_get_surah(t *testing.T) {
	var mngr Mngr

	mngr, err := Mngr_impl_new("../../qurancsv")
	assert.Nil(t, err)

	surah, err := mngr.Get_surah(114)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(surah.Verses))
	assert.Equal(t, "مِنَ الْجِنَّةِ وَالنَّاسِ", surah.Verses[5].Text)
}

func Test_get_surah_not_exist(t *testing.T) {
	var mngr Mngr
	var err error

	mngr, err = Mngr_impl_new("../../qurancsv")
	assert.Nil(t, err)

	_, err = mngr.Get_surah(115)
	assert.NotNil(t, err)
}

func Test_get_surah_infos(t *testing.T) {
	var mngr Mngr
	var err error

	mngr, err = Mngr_impl_new("../../qurancsv")
	assert.Nil(t, err)

	surahInfos, err := mngr.Get_surah_infos()
	assert.Nil(t, err)
	assert.Equal(t, 114, len(surahInfos))
	assert.Equal(t, "The Opening", surahInfos[0].English)
	assert.Equal(t, "Mankind", surahInfos[113].English)
}
