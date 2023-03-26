package tests

import (
	"fmt"
	"testing"
)

func TextPasswd(t *testing.T) {
	fmt.Println("----------------TextPasswd 开始----------------")
	
	assert := getAssert(t)
	str := "123456"
	hash, _ := passwd.Hash(str)
	assert.Nil(passwd.Check(str, hash))

	fmt.Println("----------------TextPasswd 结束----------------")
}
