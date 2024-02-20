package convey

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2022/7/17 23:38
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
 /*
 //     Convey(description string, t *testing.T, mode FailureMode, action func())
 //      Convey(description string, mode FailureMode, action func())
 //      Convey(description string, action func(c Context))
 //      Convey(description string, action func())
*/

func Test_Fibonacci(t *testing.T) {
	var count = 3
	convey.Convey("比较俩种斐波那契数列", t, func(c convey.C) {
		c.So(fibonacciV1(count), convey.ShouldEqual, fibonacciV2(1, 1, count))
	})
}

func fibonacciV1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciV1(n-1) + fibonacciV1(n-2)
}

func fibonacciV2(first int, sec int, n int) int {
	if n <= 0 {
		return 0
	}
	if n < 3 {
		return 1
	} else if n == 3 {
		return first + sec
	}
	return fibonacciV2(sec, first+sec, n-1)
}
