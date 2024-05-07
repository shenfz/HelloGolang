package Assertion_test

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 15:32
 * @Desc: 利用一些编译时候用到的断言技巧，保证条件
 */

// declare biggest value of uint and int
// 声明一个最大的 int和uint常量
const MaxUint = ^uint(0)
const MaxInt = int(^uint(0) >> 1)

// 判定系统位数
const Is64bitArch = ^uint(0)>>63 == 1
const Is32bitArch = ^uint(0)>>63 == 0
const WordBits = 32 << (^uint(0) >> 63) // 64或32

var (
	N int = 1
	M int = 9
)

//为了避免包级变量消耗太多的内存，可以把断言代码放在一个名为空标识符的函数体中

// n >= m
func _() {
	var _ = map[bool]int{false: 0, N >= M: 1}
}

//// n == m
//func _()  {
//	var _ = map[bool]int{false: 0, M==N: 1}
//}

// n <= m
//
//	func _()  {
//		var _ = map[bool]int{false: 0, M<=N: 1}
//	}
func Test_GetMaxOrArch(t *testing.T) {
	t.Log(MaxUint, MaxInt, WordBits, Is32bitArch, Is64bitArch)
}

func Test_BigDataWithInterface(t *testing.T) {
	var a [100]int
	a[0] = 1
	// 这两行的开销相对较大，因为数组a中的元素都将被复制。
	fmt.Println(a)                   // 复制副本
	fmt.Printf("Type of a: %T\n", a) // 复制副本

	// 这两行的开销较小，数组a中的元素没有被复制。
	fmt.Printf("%v\n", a[:])                             // 传的切片
	fmt.Println("Type of a:", fmt.Sprintf("%T", &a)[1:]) // [100]int
	fmt.Println("Type of a:", fmt.Sprintf("%T", &a)[:])  // *[100]int
}
