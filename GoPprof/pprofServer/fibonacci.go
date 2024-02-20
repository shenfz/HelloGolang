package main

/**
 * @Author shenfz
 * @Date 2022/7/17 21:57
 * @Email 1328919715@qq.com
 * @Description:
 **/

func FibonacciV1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciV1(n-1) + FibonacciV1(n-2)
}

func FibonacciV2(first int, sec int, n int) int {
	if n <= 0 {
		return 0
	}
	if n < 3 {
		return 1
	} else if n == 3 {
		return first + sec
	}
	return FibonacciV2(sec, first+sec, n-1)
}
