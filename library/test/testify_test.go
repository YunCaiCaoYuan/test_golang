package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	//断言相等
	assert.Equal(t, 123, 123, "they should be equal")

	//断言不相等
	assert.NotEqual(t, 123, 456, "they should not be equal")

	//对于nil的断言
	//assert.Nil(t, object)

	//对于非nil的断言
	//if assert.NotNil(t, object) {
	//	// now we know that object isn't nil, we are safe to make
	//	// further assertions without causing any errors
	//	assert.Equal(t, "Something", object.Value)
	//}
}
