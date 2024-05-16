package sliceConcat__test

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/16 16:13
 * @Email 1328919715@qq.com
 * @Description:  Go 1.21 才添加slices ，但是对源切片处理，go1.22 和 go1.21表现不同
 **/

/*
 拼接切片之前先计算了 新切片所需的长度 ，然后利用 Grow 函数初始化新切片。
 这样做的好处是避免了后续 append 操作中因为切片扩容而导致的内存重新分配和复制问题，使得函数更加高效
*/

func Test_Concat(t *testing.T) {
	s := []int{1, 2, 3, 4}
	t.Log(s[:0])
	t.Log(s[3:4])
	t.Log(s[2:3])
	t.Log(s[1:2])
	t.Log(s[0:1])
	S := slices.Concat(s[:0], s[3:4], s[2:3], s[1:2], s[0:1])
	t.Log(s)
	t.Log(S)
	/*
	    concat_test.go:24: []
	   concat_test.go:25: [4]
	   concat_test.go:26: [3]
	   concat_test.go:27: [2]
	   concat_test.go:28: [1]
	   concat_test.go:30: [1 2 3 4]
	   concat_test.go:28: [4 3 2 1]
	*/
}

/*
go 1.21       不会更改源切片内容   [1, 2, 3, 4, 5]
go1.22        会把源切片删除部分 置为零值 且 尾置   [1 2 3 5 0]
*/
func Test_Delete(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.Delete(s1, 3, 4)
	fmt.Println(s1)
	fmt.Println(s2)
}

func Test_DeleteFunc(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.DeleteFunc(s1, func(x int) bool {
		return x%2 == 0
	})
	t.Log(s1)
	t.Log(s2)
	/* go1.22
	   concat_test.go:58: [1 3 5 0 0]
	   concat_test.go:59: [1 3 5]
	*/
}

func Test_Compact(t *testing.T) {
	s1 := []string{"Gopher", "MingYong Chen", "mingyong chen"}
	s2 := slices.CompactFunc(s1, func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	})
	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)
	/* go1.22
	 []string{"Gopher", "MingYong Chen", ""}
	[]string{"Gopher", "MingYong Chen"}
	*/
}

func Test_Replace(t *testing.T) {
	s1 := []int{1, 6, 7, 4, 5}
	s2 := slices.Replace(s1, 1, 3, 2)
	fmt.Println(s1)
	fmt.Println(s2)
	/*go1.22
	[1 2 4 5 0]
	[1 2 4 5]
	*/
}

func Test_Insert(t *testing.T) {
	s1 := []string{"程序员", "陈明勇"}
	t.Log(len(s1), cap(s1))
	s2 := slices.Insert(s1, 1, "xxx")
	fmt.Println(s2)
	t.Log(len(s1), cap(s1))
	t.Log(len(s2), cap(s2))
	/* go1.22
	[程序员 陈明勇 xxx]
	*/
}
