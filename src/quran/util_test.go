package quran

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_read_quran_csv(t *testing.T) {
	records, err := read_quran_csv("../../qurancsv/quran.csv")
	assert.Nil(t, err)
	assert.Equal(t, 6236, len(records))
	assert.Equal(t, 3, len(records[0]))
	assert.Equal(t, "1", records[0][0])
	assert.Equal(t, "1", records[0][1])
}

func Test_read_quran_csv_not_found(t *testing.T) {
	_, err := read_quran_csv("../../qurancsv/x.csv")
	assert.NotNil(t, err)
	assert.Equal(t, "open ../../qurancsv/x.csv: no such file or directory", err.Error())
}
