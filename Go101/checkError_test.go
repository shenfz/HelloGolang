package Go101

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 11:13
 * @Desc:
 */
/*
   1. 需要精准确认错误或者错误类型较多，使用对应包提供的错误判断函数或者 errors.Is()
   2. 错误解包
*/

func Test_CheckError(t *testing.T) {
	_, err := os.Stat("a-nonexistent-file.abcxyz")
	fmt.Println(os.IsNotExist(err))             // true
	fmt.Println(err == os.ErrNotExist)          // false
	fmt.Println(errors.Is(err, os.ErrNotExist)) // true
}

func Test_PrintWithPosition(t *testing.T) {
	t.Logf("%[2]v %[1]v %[2]v %[1]v", "a", "m") // m a m a
}
