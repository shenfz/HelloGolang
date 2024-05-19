package GoGenericity

import (
	"reflect"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/19 16:12
 * @Email 1328919715@qq.com
 * @Description: 类型
 **/

func Test_TypeOf(t *testing.T) {
	tFor := reflect.TypeFor[string]()
	t.Log(tFor.Name())
	tOf := reflect.TypeOf("xxx")
	t.Log(tOf)
	/*
	    typeOf_test.go:17: string
	   typeOf_test.go:19: string
	*/
}
