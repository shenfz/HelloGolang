package Range

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 14:28
 * @Desc: https://gfw.go101.org/article/container.html#delete-slice-elements
 */

/*
   1. range-array: 巨量数组 ，复制其指针或使用其切片遍历，开销较小
   2. range-array: 遍历空指针 ， 如未忽略第二迭代赋值的字面量 ，则造成panic
*/

// 巨量数组 ，复制其指针或使用其切片遍历，开销较小
func Test_RangeByArrayPointer1(t *testing.T) {
	var a [100]int

	for i, n := range &a { // 复制一个指针的开销很小
		fmt.Println(i, n) // 0 0
	}

	for i, n := range a[:] { // 复制一个切片的开销很小
		fmt.Println(i, n) // 0 0
	}
}

// 遍历空指针 ， 如未忽略第二迭代赋值的字面量 ，则造成panic
func Test_RangeByArrayPointer2(t *testing.T) {
	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	for i, n := range p { // panic
		fmt.Println(i, n)
	}
}

// 通过数组的指针来访问和修改此数组中的元素 , 此指针是一个nil指针，将导致一个恐慌
func Test_ChangeValByPointer(t *testing.T) {
	a := [5]int{2, 3, 5, 7, 11}
	p := &a
	p[0], p[1] = 17, 19
	fmt.Println(a) // [17 19 5 7 11]
	p = nil
	_ = p[0] // panic
}
