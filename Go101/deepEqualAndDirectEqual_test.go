package Go101

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 11:02
 * @Desc:
 */

/*
1. reflect.DeepEqual(x, y)和x == y的结果可能会不同
2. Go类型系统禁止切片类型[]MyByte的值转换为类型[]byte
   但是，当前的reflect.Value类型的Bytes方法的实现可以帮我们绕过这个限制。
   此实现应该是违反了Go类型系统的规则
*/
func Test_DeepEqual(t *testing.T) {

	type Book struct{ page int }
	x := struct{ page int }{123}
	y := Book{123}
	fmt.Println(reflect.DeepEqual(x, y)) // false  比较类型
	fmt.Println(x == y)                  // true

	z := Book{123}
	fmt.Println(reflect.DeepEqual(&z, &y)) // true
	fmt.Println(&z == &y)                  // false  比较指针，值不同

	type T struct{ p *T }

	var f1, f2 func() = nil, func() {}
	fmt.Println(reflect.DeepEqual(f1, f1)) // true
	fmt.Println(reflect.DeepEqual(f2, f2)) // false

	var a, b interface{} = []int{1, 2}, []int{1, 2}
	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(a == b)                  // panic
}

type MyByte byte

func Test_PrivateType(t *testing.T) {
	var mybs = []MyByte{'a', 98, 'c'}
	var bs []byte

	// bs = []byte(mybs) // this line fails to compile

	v := reflect.ValueOf(mybs)
	bs = v.Bytes()                                     // okay. Violating Go type system.
	fmt.Println(bytes.HasPrefix(bs, []byte{'a', 'b'})) // true

	bs[1], bs[2] = 'r', 't'
	fmt.Printf("%s \n", mybs) // art
}
