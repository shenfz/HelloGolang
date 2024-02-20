package main

import (
	"runtime"
	"syscall"
)

/**
 * @Author shenfz
 * @Date 2021/12/1 16:48
 * @Email 1328919715@qq.com
 * @Description:
 **/

type File struct{ d syscall.Handle }

func main() {
	p := openFile("t.txt")
	content := readFile(p.d)

	runtime.KeepAlive(p) // runtime.KeepAlive 能阻止 runtime.SetFinalizer 延迟发生，保证我们的变量不被 GC 所回收

	println("Here is the content: " + content)
}

func openFile(path string) *File {
	d, err := syscall.Open(path, syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	p := &File{d}
	runtime.SetFinalizer(p, func(p *File) {
		syscall.Close(p.d)
	})

	return p
}

func readFile(descriptor syscall.Handle) string {
	doSomeAllocation()

	var buf [1000]byte
	_, err := syscall.Read(descriptor, buf[:])
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
