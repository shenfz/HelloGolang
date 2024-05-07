package Go101

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 16:15
 * @Desc:  减少边界检查  提升性能
 */
/* Go是一个内存安全的语言。在数组和切片的索引和子切片操作中，Go运行时将检查操作中使用的下标是否越界
1. BCE: bounds check elimination，边界检查消除
2. CSE: common subexpression elimination，公共子表达式消除
3. 从gc 1.11开始，如果x和y是两个切片，即使上例中使用小技巧没有被使用，y[i]中的边界检查也已经被消除了
*/

func f1(s []int) {
	_ = s[0] // 第5行： 需要边界检查
	_ = s[1] // 第6行： 需要边界检查
	_ = s[2] // 第7行： 需要边界检查
}

func f2(s []int) {
	_ = s[2] // 第11行： 需要边界检查
	_ = s[1] // 第12行： 边界检查消除了！
	_ = s[0] // 第13行： 边界检查消除了！
}

func f3(s []int, index int) {
	_ = s[index] // 第17行： 需要边界检查
	_ = s[index] // 第18行： 边界检查消除了！
}

func f4(a [5]int) {
	_ = a[4] // 第22行： 边界检查消除了！
}
