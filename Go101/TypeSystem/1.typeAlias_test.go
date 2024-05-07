package TypeSystem__test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/7 18:26
 * @Email 1328919715@qq.com
 * @Description:  绕过 类型别名 进行操作
 **/

/*
 1. Go类型系统禁止 切片类型[]MyByte的值转换为类型[]byte
   但是，当前的reflect.Value类型的Bytes方法的实现可以帮我们绕过这个限制。
   此实现应该是违反了Go类型系统的规则
*/

type MyByte byte

func Test_PrivateType(t *testing.T) {
	var mybs = []MyByte{'a', 98, 'c'}
	var bs []byte

	// bs = []byte(mybs) // 编译失败

	v := reflect.ValueOf(mybs)
	bs = v.Bytes()                                     // okay. Violating Go type system.
	fmt.Println(bytes.HasPrefix(bs, []byte{'a', 'b'})) // true

	bs[1], bs[2] = 'r', 't'
	fmt.Printf("%s \n", mybs) // art
}
