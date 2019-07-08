package testifydemo

import (
	"github.com/stretchr/testify/assert"
	"gopractice/projects/testdemo/goconveydemo"
	"testing"
)

func TestIsEqual(t *testing.T){
	ok1 := goconveydemo.IsEqual(1, 1)
	ok2 := goconveydemo.IsEqual(1, 2)

	ok3, err3 := goconveydemo.IsEqualWithErr(1, 1)
	ok4, err4 := goconveydemo.IsEqualWithErr(1, 2)

	assert.True(t, ok1, "a:1 b:1 true")
	assert.False(t, ok2, "a:1 b:2 false")

	assert.Equal(t, true, ok3, "a:1 b:1 true")
	assert.Nil(t, err3, "a:1 b:1 nil")

	assert.Equal(t, false, ok4, "a:1 b:2 false")
	assert.NotNil(t, err4, "a:1 b:2 not nil")
}