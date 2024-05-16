package main

import (
	"fmt"
	"golang.org/x/exp/trace"
	"io"
	"log"
	"os"
	"strings"
)

/**
 * @Author shenfz
 * @Date 2024/5/16 13:01
 * @Email 1328919715@qq.com
 * @Description: 用于检测因为等待网络而阻塞的 goroutine 的比例
 **/

func main() {

	f, err := os.Open("E:\\GithubGoPath\\src\\HelloGolang\\GoX\\Trace\\trace.out")
	if err != nil {
		log.Fatal(err)
	}

	r, err := trace.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	var blocked int
	var blockedOnNetwork int
	for {
		ev, err := r.ReadEvent()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if ev.Kind() == trace.EventStateTransition {
			st := ev.StateTransition()
			if st.Resource.Kind == trace.ResourceGoroutine {
				// id := st.Resource.Goroutine()
				// from, to := st.GoroutineTransition()
				from, to := st.Goroutine()

				if from.Executing(); to == trace.GoWaiting {
					blocked++
					if strings.Contains(st.Reason, "network") {
						blockedOnNetwork++
					}
				}
			}
		}
	}

	p := 100 * float64(blockedOnNetwork) / float64(blocked)
	fmt.Printf("%2.3f%% instances of goroutines blocking were to block on the network\n", p)
}
