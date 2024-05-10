package bench

import (
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/14 17:32
 * @Desc:
 */
/*
  1. slice在循环开始之前，仅计算一次，在循环过程中 修改切片长度 则 不影响本次循环次数 ,但是 内容修改 则会 影响输出
  2. map迭代 ，未迭代到的键值对被删除，则无迭代输出；新增键值对，可能会被迭代。针对nil 切片 ，迭代次数为0
  3. 遍历较简单的数据结构，for-i 和 for-range 性能差不多
  4. 当遍历结构里面有占用较大内存的结构，则尽量规避迭代导致的值复制，使用索引，性能更好
  5. 若切片存储的是指针， []*struct ， for-range 和 for-i 性能无差 ， 指针更便于直接修改结构体值

*/
func addOne(nums []int) {
	nums = append(nums, 1)
}

func Test_RangeSlice(t *testing.T) {
	var (
		intS = []int{1, 2, 3, 4, 5, 6}
	)
	for num, val := range intS {

		t.Log(num, val)
	}
}

func Test_RangeMap(t *testing.T) {
	var (
		intsMap = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}
	)
	for key, val := range intsMap {
		delete(intsMap, "2")
		t.Log(key, val)
	}
}

type Item struct {
	id  int
	val [4096]byte
}

// 5088499	       236.0 ns/op
func BenchmarkForStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// 5088058	       237.6 ns/op
func BenchmarkRangeIndexStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// 12212	    108981 ns/op
func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
