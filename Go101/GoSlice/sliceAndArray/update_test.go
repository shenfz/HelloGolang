package sliceAndArray

import (
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/7/19 16:02
 * @Email 1328919715@qq.com
 * @Description:  传递数组并改变其集合中的某些值 ： 数组传指针或者传自身的”最长切片“（array[:]）、切片传自身
 **/

// 切片 数组 都是值拷贝传递
func Test_AS(t *testing.T) {
	var (
		arrays = [...]int{1, 2, 3, 4, 5, 6}
		silces = []int{1, 2, 3, 4, 5}
		s2     = arrays[:]
	)

	ChangeArray(arrays)
	ChangeSlice(silces)

	t.Log(arrays) // [1 2 3 4 5 6]
	t.Log(silces) // [100 2 3 4 5 6]

	ChangeArrayByPoint(&arrays)
	t.Log(arrays) // [101 2 3 4 5 6]
	t.Log(s2)     // [101 2 3 4 5 6]

	ChangeSlice(s2)
	t.Log(s2)     //  [100 2 3 4 5 6]
	t.Log(arrays) //  [100 2 3 4 5 6]
}

// 传递的是数组的原始值拷贝，此时在函数内部是无法更新该数组的
func ChangeArray(a [6]int) {
	a[0] = 100
}

// 若想更新该数组,传递该数组的指针
func ChangeArrayByPoint(a *[6]int) {
	(*a)[0] = 101 //equal: a[0] = 101
}

// 切片是地址传递
func ChangeSlice(s []int) {
	s[0] = 100
}
