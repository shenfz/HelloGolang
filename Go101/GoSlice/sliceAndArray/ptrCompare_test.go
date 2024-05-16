package sliceAndArray

import (
	"fmt"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/10 18:07
 * @Email 1328919715@qq.com
 * @Description: 数组和切片比较 取地址的情况对比
 **/

type Student struct {
	Age int // 8
}

// ====================================> 数组和切片 取地址的情况对比
func Test_ArrayAndSlicePoint(t *testing.T) {

	s := []Student{{Age: 21}, {Age: 45}}

	a := [...]Student{{Age: 23}, {Age: 90}}
	bb := a[:]

	fmt.Printf("%p , %p ,%p \n", s, &s[0], &s[1])     // 0xc00000a330 , 0xc00000a330 ,0xc00000a338
	fmt.Printf("%p  , %p ,%p \n", &a, &a[0], &a[1])   // 0xc00000a340  , 0xc00000a340 ,0xc00000a348
	fmt.Printf("%p  , %p ,%p \n", bb, &bb[0], &bb[1]) // 0xc00000a340  , 0xc00000a340 ,0xc00000a348
}
