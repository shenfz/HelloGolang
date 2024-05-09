package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

/**
 * @Author shenfz
 * @Date 2021/12/1 16:46
 * @Email 1328919715@qq.com
 * @Description: 触发垃圾收集器，执行回调终结
 **/

type File struct{ d syscall.Handle }

func main() {
	absPath, err := filepath.Abs("./123.txt")
	if err != nil {
		panic(err)
	}
	println("get abs path: ", absPath)
	p := openFile(absPath)
	// set  SetFinalizer
	runtime.SetFinalizer(p, func(p *File) {
		println("exec finializer")
		syscall.Close(p.d)
	})

	//
	content := readFile(p)
	println("Here is the content: \n" + content)
	runtime.GC()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("GBDeviceControl exit by signal")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func openFile(path string) *File {
	d, err := syscall.Open(path, syscall.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	p := &File{d}
	return p
}

func readFile(f *File) string {
	doSomeAllocation()

	var buf [1000]byte
	_, err := syscall.Read(f.d, buf[:])
	if err != nil {
		panic(err)
	}

	return string(buf[:])
}

func doSomeAllocation() {
	var a *int

	// memory increase to force the GC
	for i := 0; i < 10000000; i++ {
		i := 1
		a = &i
	}

	_ = a
}
