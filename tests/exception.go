package tests

import (
	"errors"
	"testing"
)

func TestException(t *testing.T) {
	assert := getAssert(t)
	msg := "test message"
	err := errors.New(msg)
	exceptions := excp.BadRequest(err)

	assert.Equal(exceptions.Error(), msg)
	assert.Equal(exceptions.Service.RawErr, err)

	assert.NotEqual(excp, exceptions)
	assert.NotEqual(excp.Exception, exceptions.Exception)
}
