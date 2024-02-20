package bench

import (
	"math/rand"
	"testing"
	"time"
)

/*
  =====> 提前分配容量cap，有效避免扩容造成的性能损失

*/
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

//  186	   6422871 ns/op
func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

//  129	  8893428 ns/op
func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}
