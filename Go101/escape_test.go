package Go101

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 10:17
 * @Desc:
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
}

//
func Test_SelfPointer(t *testing.T) {
	// 一个指针类型的基类型可以为此指针类型自身
	type P *P
	var p P
	p = &p
	p = **************p

	/*
		 一个切片类型的元素类型可以是此切片类型自身，
		一个映射类型的元素类型可以是此映射类型自身，
		一个通道类型的元素类型可以是此通道类型自身，
		一个函数类型的输入参数和返回结果值类型可以是此函数类型自身
	*/
	type S []S
	type M map[string]M
	type C chan C
	type F func(F) F

	s := S{0: nil}
	s[0] = s
	m := M{"Go": nil}
	m["Go"] = m
	c := make(C, 3)
	c <- c
	c <- c
	c <- c
	var f F
	f = func(F) F { return f }

	_ = s[0][0][0][0][0][0][0][0]
	_ = m["Go"]["Go"]["Go"]["Go"]
	<-<-<-c
	f(f(f(f(f))))
}
