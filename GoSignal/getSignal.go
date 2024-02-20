package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author shenfz
 * @Date 2021/10/11 16:41
 * @Email 1328919715@qq.com
 * @Description: 获取信号量 ， kill -l
 **/

func main() {
	log.Println("now start working...")
	//监听服务退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("TransAgreement_GateWay Get a Signal: [%s]", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("TransAgreement_GateWay Exit By Signal: [%s]", s.String())
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
