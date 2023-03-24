package tests

import (
	"errors"
	"gower/app/exceptions"
	"testing"
)

func TestException(t *testing.T) {
	assert := getAssert(t)
	msg := "test message"
	err := errors.New(msg)
	exception := excp.BadRequest(err).(*exceptions.Exception)

	assert.Equal(exception.Error(), msg)
	assert.Equal(exception.Service.RawErr, err)

	assert.NotEqual(excp, exception)
	assert.NotEqual(excp.Exception, exception.Exception)
}
