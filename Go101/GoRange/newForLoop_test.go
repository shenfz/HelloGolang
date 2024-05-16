package GoRange

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/16 13:13
 * @Email 1328919715@qq.com
 * @Description: go1.22 遍历特性
 **/

// 闭包
func Test_LoopVar_Closure(t *testing.T) {
	sl := []int{11, 12, 13, 14, 15}
	var wg sync.WaitGroup
	for i, v := range sl {
		wg.Add(1)
		go func() {
			fmt.Printf("%d : %d\n", i, v)
			wg.Done()
		}()
	}
	wg.Wait()
	/*
		 4 : 15
		1 : 12
		0 : 11
		2 : 13
		3 : 14
	*/
}

func Test_LoopVar(t *testing.T) {
	sl := []int{11, 12, 13, 14, 15}
	var wg sync.WaitGroup
	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		// i := i   创建中间变量 在go1.22下不影响 ，go1.22之前版本会出现越界
		go func() {
			v := sl[i]
			fmt.Printf("%d : %d\n", i, v)
			wg.Done()
		}()
	}
	wg.Wait()
	/*
		 4 : 15
		3 : 14
		0 : 11
		1 : 12
		2 : 13

	*/
}

// 整形表达式
func Test_IntExpression(t *testing.T) {
	//n := 5
	// 如果n <= 0，则循环不运行任何迭代
	// 理解为是一种“语法糖”
	//for i := range n {
	//	fmt.Println(i)
	//}
	//for i := 0; i < 5; i++ {
	//
	//}
}

// 反向迭代切片
func Test_backwards(t *testing.T) {
	sl := []string{"hello", "world", "golang"}
	// ===================>   go1.22之前
	Backward(sl)(func(i int, s string) bool {
		fmt.Printf("%d : %s\n", i, s)
		return true
	})

	//======================> go1.22开启实验特性 $GOEXPERIMENT=rangefunc

	//for i, s := range Backward(sl) {
	//	fmt.Printf("%d : %s\n", i, s)
	//}
}

func Backward[E string | int](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
		return
	}
}
