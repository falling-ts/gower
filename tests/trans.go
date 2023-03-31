package tests

import (
	"fmt"
	"testing"
)

var _ = []string{
	"Error 1062 (23000): Duplicate entry 'aaa' for key 'users.idx_users_username'",
}

func TestTrans(*testing.T) {
	fmt.Println("----------------TestTrans 开始----------------")

	//assert := getAssert(t)
	//for _, msg := range messages {
	//	err := trans.DBError(errors.New(msg))
	//	assert.NotEmpty(err.Error())
	//	assert.NotEqual(err.Error(), msg)
	//}

	fmt.Println("----------------TestTrans 结束----------------")
}
