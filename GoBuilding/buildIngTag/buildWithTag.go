package main

import (
	"github.com/shenfz/HelloGolang/GoBuilding/buildIngTag/needTags"
	"net/http"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 9:00
 * @Desc:
 */
/*
   死码消除：
   1. 编译时期，编译器做内联优化，计算能确定的分支，并消除
   2. 在声明全局变量时，如果能够确定为常量，尽量使用 const 而非 var
   3. 死码消除后，既减小了二进制的体积，又可以提高运行时的效率
   4. 在声明局部变量时，编译器死码消除是生效的（能确定的值）
   5. 包(package)级别的变量和函数内部的局部变量的推断难度是不一样的，很难做到优化

   调试模式：
   1. 当设置常量Debug=false后编译，调试语句在编译的时候会被消除，对最终二进制大小和运行效率都有好处

   条件编译：
   1. 结合 build tags 来实现条件编译 eg: go build -tags Debug -o main
   2. 同行 +build 表示或的关系，异行表示与的关系
   3. 集合全局布尔值常量（这里是debug）实现条件编译
*/

//go:generate go run -tags debug main.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!! ==> " + needTags.PrintDebug()))
	})
	http.ListenAndServe(":8081", nil)
}
