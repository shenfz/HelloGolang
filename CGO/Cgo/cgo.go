//+build debug
//+build windows,386 cgo

package main

/*
  中 windows,386 中 windows和386用逗号链接表示AND的意思
  而 windows,386 和 !cgo 之 间通过空白分割来表示OR的意思

*/

// #cgo CFLAGS: -D PNG_DEBUG=1 -I ./include
// #cgo LDFLAGS: -L /usr/local/lib -l png

/*
  -D : 宏定义 PNG_DEBUG
  -I : 定义了头文件包含的检索目录
  -L : 指定了链接时库文件检索目录
  -l ： 指定了链接时需要链接png库
*/

// #cgo windows CFLAGS: -DX86=1
// #cgo !windows LDFLAGS: -l m

/*
   支持条件选择，根据环境或者架构来生效
   win环境下，定义宏 DX86
   非win环境下，连接math 数学库
*/

/*
#cgo windows CFLAGS: -D CGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -D CGO_OS_DARWIN=1
#cgo linux CFLAGS: -D CGO_OS_LINUX=1
#if defined(CGO_OS_WINDOWS)
  const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
  static const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
  static const char* os = "linux";
#else
#   error(unknown os)
#endif */
import "C"

func main1() {
	print(C.GoString(C.os)) // 	// windows 处理不同平台之间的代码差异
}
