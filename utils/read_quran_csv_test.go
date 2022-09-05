package utils_test

import (
	"testing"

	"github.com/arikama/koran-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestReadQuranCsv(t *testing.T) {
	records, err := utils.ReadQuranCsv("quran")
	assert.Nil(t, err)
	assert.Equal(t, 6236, len(records))
	assert.Equal(t, 3, len(records[0]))
	assert.Equal(t, "1", records[0][0])
	assert.Equal(t, "1", records[0][1])
}
