package utils_test

import (
	"testing"

	"github.com/arikama/koran-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestReadSurahsInfo(t *testing.T) {
	surahInfos, err := utils.ReadSurahsInfo("./../qurancsv/surahs.csv")
	assert.Nil(t, err)
	assert.Equal(t, 114, len(surahInfos))
	assert.Equal(t, "Al-Fatihah", surahInfos[0].Title)
	assert.Equal(t, "An-Naas", surahInfos[113].Title)
}

func TestReadSurahsInfoFileNotFound(t *testing.T) {
	_, err := utils.ReadSurahsInfo("./../qurancsv/x.csv")
	assert.NotNil(t, err)
	assert.Equal(t, "open ./../qurancsv/x.csv: no such file or directory", err.Error())
}
