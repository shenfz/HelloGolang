package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/**
 * @Author shenfz
 * @Date 2021/11/19 10:36
 * @Email 1328919715@qq.com
 * @Description:
 **/

type Result struct {
	Num, Ans int
}

type Calc int

// Square calculates the square of num
func (calc *Calc) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

var (
	port string = "1234"
)

func main() {
	err := rpc.Register(new(Calc))
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	log.Printf("Serving RPC server on port :%s", port)
	if err := http.ListenAndServe(net.JoinHostPort("", port), nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
