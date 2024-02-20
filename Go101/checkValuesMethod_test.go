package Go101

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 14:21
 * @Desc: 检测值是否具有某种方法，不使用反射
 */

type StructA int
type StructB int

func (b StructB) M(x int) string {
	return fmt.Sprint(b, ": ", x)
}

// 使用了断言 ， 判定v是否实现了 匿名接口中的M方法
func check(v interface{}) bool {
	_, has := v.(interface{ M(int) string })
	return has
}

func Test_CheckedWithoutReflect(t *testing.T) {
	var a StructA = 123
	var b StructB = 789
	fmt.Println(check(a)) // false
	fmt.Println(check(b)) // true
}
