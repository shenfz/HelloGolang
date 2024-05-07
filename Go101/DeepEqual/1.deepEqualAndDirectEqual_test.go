package DeepEqual__test

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 11:02
 * @Desc: 直接比较 和 深比较
 */

/*
1. reflect.DeepEqual(x, y)和x == y的结果可能会不同
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

	var f1, f2 func() = nil, func() {}
	fmt.Println(reflect.DeepEqual(f1, f1)) // true
	fmt.Println(reflect.DeepEqual(f2, f2)) // false

	var a, b interface{} = []int{1, 2}, []int{1, 2}
	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(a == b)                  // panic
}
