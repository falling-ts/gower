package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetKey(t *testing.T) {
	a := assert.New(t)
	a.Len(getKey(), 21)
}
