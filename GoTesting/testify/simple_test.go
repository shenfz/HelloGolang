package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/10/11 11:16
 * @Email 1328919715@qq.com
 * @Description: testify 用法
 **/

func Test_Sample(t *testing.T) {
	var (
		a = 100

		b = 100
	)
	assert.Equal(t, a, b)
}

// 数组、切片 包含 值
// 字典  包含 键
func Test_Contains(t *testing.T) {
	var (
		sub = "x"
		c1  = []string{"x", "1", "a", "b"}
		c2  = map[string]string{"x": "x", "a": "a", "b": "b"}
		//c3  = struct {
		//	A string
		//	B int
		//}{"x",1}
		c4 = [4]string{"x", "1"}
	)

	assert.Contains(t, c1, sub)
	assert.Contains(t, c2, sub)
	//assert.Contains(t, c3,sub)
	assert.Contains(t, c4, sub)
}

// 目录存在
func Test_IsDirExists(t *testing.T) {
	var (
		path = "./static"
	)
	assert.DirExists(t, path)
}

// 包含相同的元素，不论顺序，重复元素的次数也必须相同
func Test_ListSame(t *testing.T) {
	var (
		listA = []interface{}{"x", 2, 3, 4, 4, 55, 66, 77}
		listB = []interface{}{"x", 2, 3, 4, 4, 55, 77, 66}
	)
	assert.ElementsMatch(t, listA, listB)
}

//断言 object 是空 ， 但类型不同含义也不同
func Test_IsEmpty(t *testing.T) {
	var (
		ePointer = new(string)
		eInt     int
		eFloat   float64
		eStr     string
		eBoolean bool
		eChannel chan string
		eSlice   []string
	)
	assert.Empty(t, ePointer)
	assert.Empty(t, eInt)
	assert.Empty(t, eFloat)
	assert.Empty(t, eStr)
	assert.Empty(t, eBoolean)
	assert.Empty(t, eChannel)
	assert.Empty(t, eSlice)
}

func Test_EqualError(t *testing.T) {

}
