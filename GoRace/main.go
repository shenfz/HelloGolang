package main

import (
	"flag"
	"github.com/shenfz/HelloGolang/GoRace/mapMan"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/**
 * @Author shenfz
 * @Date 2022/2/15 16:12
 * @Email 1328919715@qq.com
 * @Description: 竞争态检测
 **/

var (
	racer int
)

func init() {
	flag.IntVar(&racer, "r", 5, "how many racer")
	flag.Parse()
}

//go:generate go run -race main.go
func main() {
	cc := mapMan.GetConcurrentMap()

	for i := 0; i < racer; i++ {
		go func(num int) {
			for {
				randIn := rand.Intn(1000)
				cc.Set(strconv.Itoa(randIn), randIn)
				time.Sleep(200 * time.Millisecond)
				log.Printf("Num:%d  Get:%v", num, cc.Get(strconv.Itoa(randIn)))
			}
		}(i)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println(" exit by signal")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
