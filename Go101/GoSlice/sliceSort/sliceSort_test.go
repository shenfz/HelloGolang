package sliceSort

import (
	"sort"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/15 17:47
 * @Email 1328919715@qq.com
 * @Description: 通过sort包
 **/

// 切片排序
func Test_SortInsert(t *testing.T) {
	IsExisted([]string{"xc", "cx"}, "xc")
}

func IsExisted(ex []string, in string) (index int, b bool) {
	index = sort.SearchStrings(ex, in)
	b = len(ex) != index && ex[index] == in
	return
}
