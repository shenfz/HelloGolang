package Go101

import "testing"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 14:01
 * @Desc: 强制声明 带字段名
 */

type Config struct {
	_    [0]int
	Name string
	Size int
}

func Test_MustWithName(t *testing.T) {
	//	_ = Config{[0]int{}, "bar", 123} // 编译不通过
	_ = Config{Name: "bar", Size: 123} // 带字段名 ， 编译没问题
}
