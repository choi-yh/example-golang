package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeHash(t *testing.T) {
	str := "string"

	hash, err := EncodeHash(str)
	assert.NoError(t, err)

	t.Log(hash)
}
