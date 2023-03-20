package tests

import (
	"errors"
	"testing"
)

var messages = []string{
	"Error 1062 (23000): Duplicate entry 'aaa' for key 'users.idx_users_username'",
}

func TestTrans(t *testing.T) {
	assert := getAssert(t)
	for _, msg := range messages {
		err := trans.DBError(errors.New(msg))
		assert.NotEmpty(err.Error())
		assert.NotEqual(err.Error(), msg)
	}

}
