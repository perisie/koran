package daos_test

import (
	"testing"

	"github.com/arikama/koran-backend/daos"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	assert.Equal(t, "sql: no rows in result set", daos.ErrSqlNoRowsInResultSet())
}
