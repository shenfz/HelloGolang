package Escape

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 10:17
 * @Desc:  逃逸
 */

// 两个零尺寸值的地址可能相等，也可能不相等
// 依赖于具体编译器实现以及具体编译器版本
func Test_InitValEscape(t *testing.T) {
	a := struct{}{}
	b := struct{}{}
	x := struct{}{}
	y := struct{}{}
	m := [10]struct{}{}
	n := [10]struct{}{}
	o := [10]struct{}{}
	p := [10]struct{}{}

	fmt.Println(&x, &y, &o, &p)

	// 对于标准编译器1.17版本，x、y、o和p将
	// 逃逸到堆上，但是a、b、m和n则开辟在栈上。

	fmt.Println(&a == &b) // false
	fmt.Println(&x == &y) // true
	fmt.Println(&a == &x) // false

	fmt.Println(&m == &n) // false
	fmt.Println(&o == &p) // true
	fmt.Println(&n == &p) // false
	/*
		 false
		true
		false
		false
		true
		false
	*/
}
