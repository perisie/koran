package daos_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"perisie.com/koran/daos"
)

func TestErr(t *testing.T) {
	assert.Equal(t, "sql: no rows in result set", daos.ErrSqlNoRowsInResultSet())
}
