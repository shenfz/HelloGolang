package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

/**
 * @Author shenfz
 * @Date 2022/7/17 20:51
 * @Email 1328919715@qq.com
 * @Description: 开启pprof路由端口
 **/

//  https;//localhost:6060/debug/pprof
//  https;//localhost:6060/debug/pprof/profile?seconds=60
func main() {
	var (
		execCount = 40
	)
	go func() {
		for i := 0; i < execCount; i++ {
			FibonacciV1(execCount)
		}
		log.Println("gor1 exited")
	}()

	go func() {
		for i := 0; i < execCount; i++ {
			FibonacciV2(1, 2, execCount)
		}
		log.Println("gor2 exited")
	}()

	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatalln(err)
	}
}
