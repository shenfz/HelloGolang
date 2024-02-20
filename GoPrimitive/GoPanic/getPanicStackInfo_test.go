package GoPanic

import (
	"fmt"
	"log"
	"runtime"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 17:04
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
 recover 必须在 defer 函数中运行。recover 捕获的是 祖父级 调用时的异常，直接调用时无效
*/
func a() {
	fmt.Println("a")
	b()
}

func b() {
	fmt.Println("b")
	c()
}

type Student struct {
	Name int
}

func c() {
	defer RecoverFromPanic("fun c")
	fmt.Println("c")
	var a *Student
	fmt.Println(a.Name)
}

func main() {
	a()
}

func RecoverFromPanic(funcName string) {
	log.SetFlags(19)
	if e := recover(); e != nil {
		buf := make([]byte, 64<<10)
		buf = buf[:runtime.Stack(buf, false)]
		log.Printf("func_name:[%s] ,PanicErr[ %v ] \n Stack ========> [ %s ]", funcName, e, string(buf))
	}
	return
}
