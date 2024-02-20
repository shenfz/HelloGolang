package bench

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/14 15:50
 * @Desc:
 */

/*   ==============================》   提前分配容量cap，有效避免扩容造成的性能损失
fmt.sprintf() 和 +             ： 最占用内存和耗时，拼接途中需要多次开辟内存
byte.buffer 和 strings.builder :  内存消耗小，性能好，builder比buffer要高，前者是[]byte转string，后者是复制
*/

const (
	CountCat     = 100
	StringConcat = "i9h80du0dhj0hjeh0340nopn"
)

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

//   81331	     15122 ns/op
func Benchmark_direct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusConcat(CountCat, StringConcat)
	}
}

func concatByFmt(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

//  46381	     26191 ns/op
func Benchmark_UseFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByFmt(CountCat, StringConcat)
	}
}

func concatByStringBuilder(n int, str string) string {
	builder := strings.Builder{}
	builder.Grow(n * len(str)) // 申请容量，避免扩容带来性能损失
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 1975567	       593.4 ns/op
func Benchmark_StringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByStringBuilder(CountCat, StringConcat)
	}
}

func byteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str)) // alloc with cap 提前声明容量，减少扩容需求
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// 1619971	       753.4 ns/op
func Benchmark_UseByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteConcat(CountCat, StringConcat)
	}
}
