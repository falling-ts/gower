package tests

import (
	"fmt"
	"testing"
)

func TestSymCrypt(t *testing.T) {
	fmt.Println("----------------TestSymCrypt 开始----------------")

	assert := getAssert(t)
	plaintext := "123456"
	encrypt, err := symCrypt.Encrypt(plaintext)
	assert.Nil(err)
	decrypt, err := symCrypt.Decrypt(encrypt)
	assert.Nil(err)
	assert.Equal(plaintext, decrypt)

	fmt.Println("----------------TestSymCrypt 结束----------------")
}
