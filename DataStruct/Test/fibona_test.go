package Test

import "testing"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/12 17:26
 * @Desc:
 */
/*
 算法的优先级排列如下，一般排在上面的要优于排在下面的：

常数复杂度：O(1)
对数复杂度：O(logn)
一次方复杂度：O(n)
一次方乘对数复杂度：O(nlogn)
乘方复杂度：O(n^2)，O(n^3)
指数复杂度：O(2^n)
阶乘复杂度：O(n!)
无限大指数复杂度：O(n^n)

*/
//斐波那契数列  后一个是前俩个数列的和
//eg: 1 1 2 3 5 8 13 21 ... N  N+1 2N+1
func Test_Suan(t *testing.T) {
	//第N个 数
	var count = 7

	t.Logf("Count : %d , Data : %d", count, Filbonaqie(count, 1, 1))
}

func Filbonaqie(n int, ender int, beforer int) int {
	if n == 0 {
		return ender
	}
	//如果一个函数中所有递归形式的调用都出现在函数的末尾，
	// 我们称这个递归函数是尾递归的
	return Filbonaqie(n-1, beforer+ender, ender)
}
