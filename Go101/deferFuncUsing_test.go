package Go101

import (
	"fmt"
	"os"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 14:34
 * @Desc: https://gfw.go101.org/article/defer-more.html#kind-of-resource-leaking
 */

/*
   1. 自定义延迟调用函数的返回结果可以舍弃
   2. 大多内置函数结果不可被舍弃,所以大多内置函数需要包在一个匿名函数内部延迟执行
   3. 延迟方法在压入调用栈之前完成调用值估值 ， 属主实参同理
   4. 大量延迟调用导致的暂时性内存泄露，如批量读写文件，函数推出前所有文件句柄不能及时释放。可结合匿名方法改进
*/
func Test_DeferAppend(t *testing.T) {
	s := []string{"a", "b", "c", "d"}
	defer fmt.Println(s) // [a x y d]
	// defer append(s[:1], "x", "y") // 编译错误
	// defer len(s)                  编译错误
	// defer cap(s)                  编译错误
	defer func() {
		_ = append(s[:1], "x", "y")
		_ = len(s)
		_ = cap(s)
	}()
}

/*
  一个被延迟调用的 函数值 是在其调用被推入 延迟调用堆栈 之前 被估值的
*/
func Test_DeferValuate(t *testing.T) {
	var f = func() {
		fmt.Println(false)
	}
	defer f() // false 此处被估值，推入延迟调用栈
	f = func() {
		fmt.Println(true)
	}
}

type Tint int

func (t Tint) M(n int) Tint {
	print(n)
	return t
}

/*
  tt.M(1) 属主实参   打印出 1342
*/
func Test_MainBelongParams(t *testing.T) {
	var tt Tint
	// tt.M(1)是方法调用M(2)的属主实参，因此它
	// 将在M(2)调用被推入延迟调用堆栈之前被估值。
	defer tt.M(1).M(2)
	tt.M(3).M(4)
}

/*
  结合匿名函数 梅调用一个文件句柄，在该匿名函数完成后都会及时释放
*/
func Test_ReleaseFDInTime(t *testing.T) {
	files := []os.File{}
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.Name())
			if err != nil {
				return err
			}
			defer f.Close() // 将在此循环步内执行

			_, err = f.WriteString("xx")
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			t.Log(err)
		}
	}
}
