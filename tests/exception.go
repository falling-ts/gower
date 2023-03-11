package tests

import (
	"errors"
	"testing"
)

func TestException(t *testing.T) {
	assert := getAssert(t)
	msg := "test message"
	err := errors.New(msg)
	exception := excp.BadRequest(err)

	assert.Equal(exception.Error(), msg)
	assert.Equal(exception.Struct.RawErr, err)

	assert.NotEqual(excp, exception)
	assert.NotEqual(excp.Struct, exception.Struct)
}
