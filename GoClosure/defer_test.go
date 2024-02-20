package ClosureQ

import "fmt"

/**
 * @Author shenfz
 * @Date 2021/8/30 16:29
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  须知，return 不是单一性操作
  1. 返回值赋值
  2. 执行defer ,先进后出原则

  defer 函数 执行规则
  1. defer 函数中的return 忽略掉
  2. defer 函数传参和外部函数返回参数同名，作用域问题
  3. defer 值引用 保存规则
*/

// 先对 result赋值为 0 , defer 表达式里 result++有用 ===> 1
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// defer 函数里面传参名 与 外部函数返回值名 相同，但是两者的result不是同一个东西了 === 》 5
func f2() (result int) {
	defer func(result int) {
		result = result + 1
	}(result)
	return 5
}

func f3() {
	// 对for遍历里面的i值引用 ，最终i等于5 ======》  5，5，5，5，5
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	// 值传递，把每次i当作参数传递，结果  =======》  4，3，2，1，0
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
	// 每次都声明一个中间值，保存遍历的i，同时defer值引用也会保存 ===》 4.3.2.1.0
	for i := 0; i < 5; i++ {
		tmp := i
		defer func() {
			fmt.Println(tmp)
		}()
	}

}
