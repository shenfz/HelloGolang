package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2022/7/17 21:54
 * @Email 1328919715@qq.com
 * @Description: 使用pprof分析性能
 **/

/*

// go test -bench=. -cpuprofile=cpu.prof
// go tool pprof -http=:8080 cpu.prof
*/

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
		c.So(FibonacciV1(count), convey.ShouldEqual, FibonacciV2(1, 1, count))
	})
}

// 1 1 2 3 5 8 13 21
func TestName(t *testing.T) {
	t.Log(FibonacciV1(8))       // 21
	t.Log(FibonacciV2(1, 1, 8)) //21
}

// BenchmarkFiboV1-16    	       2	 540087000 ns/op
func BenchmarkFiboV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciV1(40)
	}
}

// BenchmarkFiboV2-16    	17026831	        77.33 ns/op
func BenchmarkFiboV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciV2(1, 1, 40)
	}
}
