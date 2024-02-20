package main

//#include <hello.h>
import "C"

// 调用c/c++ 函数
func main() {
	C.SayHelloByC(C.CString("xxx"))
	C.SayHelloByCMore(C.CString("xxx"))
}

// 用Go重新实现C函数
func main1() {
 C.SayHelloByGo(C.CString("cgo生成的C语言版本 SayHello函数最终会通过桥接代码调用Go语言版本的SayHello函数"))
}


/*


*/