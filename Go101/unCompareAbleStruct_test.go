package Go101

import "testing"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 13:57
 * @Desc: 致使结构不可被比较
 */

// 放置一个非导出的零尺寸的不可比较类型的字段
type UnCompare struct {
	// dummy        [0]func()
	AnotherField int
}

var x map[UnCompare]int // 编译错误：非法的键值类型
func Test_noCompareAbleStruct(t *testing.T) {
	var a, b UnCompare
	_ = a == b // 编译错误：非法的比较
}
