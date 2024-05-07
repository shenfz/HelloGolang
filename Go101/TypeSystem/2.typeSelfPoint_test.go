package TypeSystem_

import "testing"

/**
 * @Author shenfz
 * @Date 2024/5/7 18:59
 * @Email 1328919715@qq.com
 * @Description: 指针类型的 基类型 是自己的情况
 **/

func Test_SelfPointer(t *testing.T) {
	// 一个指针类型的 基类型可以为 此指针类型自身
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
	s := S{0: nil}
	s[0] = s
	_ = s[0][0][0][0][0][0][0][0]

	type M map[string]M
	m := M{"Go": nil}
	m["Go"] = m
	_ = m["Go"]["Go"]["Go"]["Go"]

	type C chan C
	c := make(C, 3)
	c <- c
	c <- c
	c <- c
	<-<-<-c

	type F func(F) F
	var f F
	f = func(F) F { return f }
	f(f(f(f(f))))
}
