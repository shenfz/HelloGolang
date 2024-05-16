package main

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/trace"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/**
* @Author shenfz
* @Date 2024/5/15 18:08
* @Email 1328919715@qq.com
* @Description:  golang.org/x/exp/trace  设置黑匣子用来记录长 HTTP 请求的示例
                 输出 记录超时的最新次请求 trace.out文件，用于 go tool trace ./trace.out分析
**/

func main1() {
	fr := trace.NewFlightRecorder()
	fr.Start()
	var once sync.Once
	http.HandleFunc("/my-endpoint", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		doWork(w, r)

		if time.Since(start) >= 300*time.Millisecond {

			once.Do(func() {
				var b bytes.Buffer
				_, err := fr.WriteTo(&b)
				if err != nil {
					log.Print(err)
					return
				}

				if err := os.WriteFile(filepath.Clean("E:\\GithubGoPath\\src\\HelloGolang\\GoX\\Trace\\trace.out"), b.Bytes(), 0o755); err != nil {
					log.Print(err)
					return
				} else {
					log.Println("output trace.out success")
				}
			})
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func doWork(w http.ResponseWriter, r *http.Request) {
	virtualCost := time.Duration(rand.IntN(2000)) * time.Millisecond
	time.Sleep(virtualCost)
	w.Write([]byte(fmt.Sprintf("waitTime: %v ms", virtualCost.Milliseconds())))
}
