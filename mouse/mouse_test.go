package mouse

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	mouse := Mouse_new("./data")
	key := "key"
	value := "value"

	_ = mouse.Put(key, bytes.NewBufferString(value).Bytes())
	actual, _ := mouse.Get(key)

	assert.Equal(t, value, string(actual))
}

func Test_1000(t *testing.T) {
	mouse := Mouse_new("./data")
	for i := 0; i < 2; i++ {
		key := fmt.Sprintf("key_%v", i)
		value := fmt.Sprintf("value_%v", i)

		_ = mouse.Put(key, bytes.NewBufferString(value).Bytes())
		actual, _ := mouse.Get(key)

		assert.Equal(t, value, string(actual))
	}
}

func Test_error_key(t *testing.T) {
	mouse := Mouse_new("./data")
	key := "KEY"
	value := "value"

	err := mouse.Put(key, bytes.NewBufferString(value).Bytes())

	assert.NotNil(t, err)
	assert.Equal(t, "invalid key", err.Error())
}
