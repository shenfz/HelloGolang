package Go101

import "testing"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 10:59
 * @Desc:  栈溢出的错误不可被恢复
 */

func fcv() {
	fcv()
}

// 在目前的主流Go编译器实现中，栈溢出是致命错误。一旦栈溢出发生，程序将不可恢复地崩溃
func Test_RecoverCannotStopStackOver(t *testing.T) {
	defer func() {
		recover()
	}()
	fcv()
}
