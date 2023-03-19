package tests

import "testing"

func TextPasswd(t *testing.T) {
	assert := getAssert(t)
	str := "123456"
	hash, _ := passwd.Hash(str)
	assert.Nil(passwd.Check(str, hash))
}
