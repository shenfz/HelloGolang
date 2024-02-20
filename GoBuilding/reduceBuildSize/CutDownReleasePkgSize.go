package main

import (
	"log"
	"net/http"
	"net/rpc"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 17:24
 * @Desc: 借助命令或第三方工具  削减发行包的体积
 */

/*
   1. Normal         [9.8M]       : go build -o server main.go
   2. DropDebugInfo  [7.8M]       : go build -ldflags="-s -w" -o server main.go
   3. UPX            [5.0M]       : go build -ldflags="-s -w" -o server main.go && upx -9 server
      upx 压缩等级： 1--9
      upx 原理： 在程序开头或其他合适的地方插入解压代码，将程序的其他部分压缩。执行时完成解压
*/

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

func main() {
	rpc.Register(new(Calc))
	rpc.HandleHTTP()
	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
