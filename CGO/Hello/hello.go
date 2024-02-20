package main
import "C"
import "fmt"

/*export SayHello 指令将Go语言实现的函数 SayHello 导出为C语言函数
们可以将SayHello当作一个标准库的函数使用*/

//export SayHelloByGo
func SayHelloByGo(s *C.char) {
	fmt.Print(C.GoString(s))
}