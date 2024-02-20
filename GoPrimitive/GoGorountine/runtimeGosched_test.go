package GoGorountine

import (
	"fmt"
	"runtime"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 17:09
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
 执行到runtime.Gosched( )时会暂停向下执行，直到其它协程执行完后，再回到该位置
*/

func Test_GoSched(t *testing.T) {
	//匿名函数
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		//时间出让
		runtime.Gosched()
		fmt.Println("hello")
	}
}

/* output:
world
world
hello
hello
*/
