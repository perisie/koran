package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"perisie.com/koran/utils"
)

func TestReadQuranCsv(t *testing.T) {
	records, err := utils.ReadQuranCsv("./../qurancsv/quran.csv")
	assert.Nil(t, err)
	assert.Equal(t, 6236, len(records))
	assert.Equal(t, 3, len(records[0]))
	assert.Equal(t, "1", records[0][0])
	assert.Equal(t, "1", records[0][1])
}

func TestReadQuranCsvFileNotFound(t *testing.T) {
	_, err := utils.ReadQuranCsv("./../qurancsv/x.csv")
	assert.NotNil(t, err)
	assert.Equal(t, "open ./../qurancsv/x.csv: no such file or directory", err.Error())
}
