package SliceQ

import (
	"sort"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/7/19 16:02
 * @Email 1328919715@qq.com
 * @Description:
 **/

// 切片排序
func Test_SortInsert(t *testing.T) {
	IsExisted([]string{"xc", "cx"}, "xc")
}

func IsExisted(ex []string, in string) (index int, b bool) {
	index = sort.SearchStrings(ex, in)
	b = len(ex) != index && ex[index] == in
	return
}

// 切片 数组
func Test_AS(t *testing.T) {
	var (
		arrays = [6]int{1, 2, 3, 4, 5, 6} // 数组是内置类型，是一组同类型数据的集合，它是值类型，通过从0开始的下标索引访问元素值。在初始化后长度是固定的，无法修改其长度
		silces = []int{1, 2, 3, 4, 5}     // 切片是地址传递；切片可以通过数组来初始化
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
