package bench

import (
	"fmt"
	"runtime"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/14 16:29
 * @Desc:
 */

/*
   1. 数组之上切片，注意容量问题造成的扩容
   2. 删除后，将空余的位置置为空,有助于垃圾回收 eg del back : a[len(a)-1] = nil , a = a[:len(a)-1]
   3. 切片之上re-slice,不会创建新数组，底层数组一直得不到释放，占据大量空间，此时，最好使用make-copy替代re-slice
*/

func printLenCap(nums []int) {
	fmt.Printf("Pointer: %p ,len: %d, cap: %d %v\n", nums, len(nums), cap(nums), nums)
}

func Test_Change(t *testing.T) {
	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
	printLenCap(nums2) // len: 2, cap: 6 [3 4]

	nums2 = append(nums2, 50, 60)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]
}

func Test_someKindCopy(t *testing.T) {
	// method.1  新版本对make-copy有优化，性能略高
	// b = make([]T,len(a))
	// copy(b,a)

	// method2. 如果a=nil ，复制出来的
	// b = append([]T{},a...)

	// method3. a[:0:0] 第二个0 给定了容量为0，长度为0的切片 ，再往里面依次填充,超过容量，分配到另外一块区域
	a := []int{1, 2, 3, 4, 5}
	printLenCap(a[:0:0]) // len: 0, cap: 0 []
	b := append(a[:0:0], a...)
	printLenCap(a) // len: 5, cap: 5 [1 2 3 4 5]
	printLenCap(b) // len: 5, cap: 6 [1 2 3 4 5]
}

func Test_someKindDelete(t *testing.T) {
	// method1.  del order i
	// a = append(a[:i],a[i+1:])

	// method2. 用a[i+1:] 把 a[i:] 的元素覆盖【345 --》 455】，加上复制个体，再切片
	a := []int{1, 2, 3, 4, 5}
	i := 2
	printLenCap(a)                 // len: 5, cap: 5 [1 2 3 4 5]
	printLenCap(a[i:])             // len: 3, cap: 3 [3 4 5]
	printLenCap(a[i+1:])           // len: 2, cap: 2 [4 5]
	a = a[:i+copy(a[i:], a[i+1:])] // copy完成后 a = [1,2,4,5,5]
	printLenCap(a)                 // len: 4, cap: 5 [1 2 4 5]
}

func Test_SomeKindFilter(t *testing.T) {
	// method1. filter 原切片不再被使用，把过滤成功的元素覆盖切片头部，最后再切出来，节省内存空间
	sums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	index := 0
	for _, val := range sums {
		if val%2 == 0 {
			sums[index] = val
			index++
		}
	}
	sums = sums[:index]
}

// ==================================》 copy replace re-slice
func printMem(t *testing.T) {
	t.Helper() // 将该函数标记为助手函数， 打印行信息和文件信息则跳过
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}
func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
		runtime.GC() // 手动调用gc  将无引用的底层数组回收
	}
	printMem(t)
	_ = ans
}
func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) } // 100.26 MB
func TestLastCharsByCopy(t *testing.T)  { testLastChars(t, lastNumsByCopy) }  //  0.26 MB
