package tests

import (
	"errors"
	"fmt"
	"gitee.com/falling-ts/gower/app/exceptions"
	"testing"
)

func TestException(t *testing.T) {
	fmt.Println("----------------TestException 开始----------------")

	assert := getAssert(t)
	msg := "test message"
	err := errors.New(msg)
	exception := excp.BadRequest(err).(*exceptions.Exception)

	assert.Equal(exception.Error(), msg)
	assert.Equal(exception.Service.RawErr, err)

	assert.NotEqual(excp, exception)
	assert.NotEqual(excp.Exception, exception.Exception)

	fmt.Println("----------------TestException 结束----------------")
}
