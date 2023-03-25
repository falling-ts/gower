package tests

import "testing"

func TestAuth(t *testing.T) {
	assert := getAssert(t)
	userId := "1520"
	aud := []string{"192.168.10.11:25468"}

	token, err := auth.Sign("test_model", userId, aud)
	assert.Nil(err)
	assert.NotEmpty(token)
	assert.True(auth.IsToken(token))

	id, newToken, err := auth.Check(token, aud...)
	assert.Nil(err)
	assert.Equal(userId, id)
	assert.Empty(newToken)

	err = auth.Black(token)
	assert.Nil(err)

	_, _, err = auth.Check(token)
	assert.Equal(err.Error(), "token 已拉黑")
}
