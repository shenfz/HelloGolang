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
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rr.Int())
	}
	return nums
}

func generate(n int) []int {
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rr.Int())
	}
	return nums
}

// BenchmarkGenerateWithCap-16          348           3 455 271 ns/op
func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

// BenchmarkGenerate-16                 166           7 122 513 ns/op
func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}
