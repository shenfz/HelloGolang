package safeStartGor

import (
	"log"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 17:10
 * @Email 1328919715@qq.com
 * @Description:
 **/

func MustPanicFunc() error {
	var (
		err error
	)
	log.Printf(err.Error())
	return err
}

func MustPanicFunc2() {
	var (
		err error
	)
	log.Printf(err.Error())
}

func Clear() {
	log.Println("im clear")
}

// 简单打印方法名
func Test_Try(t *testing.T) {
	t.Log(Try(MustPanicFunc))
	/*
	   xgo_test.go:36: At: safeStartGor.MustPanicFunc : runtime error: invalid memory address or nil pointer dereference
	*/
}

func Test_TryStack(t *testing.T) {
	Try_Stack(MustPanicFunc2)
	/*
		 Xgo17:16:38 xgo.go:96: recover.err => 'runtime error: invalid memory address or nil pointer dereference'
		Xgo17:16:38 xgo.go:97: Stack[ ==>
		github.com/shenfz/HelloGolang/Go101/GoGorountine/safeStartGor.MustPanicFunc2()
			E:/GithubGoPath/src/HelloGolang/Go101/GoGorountine/safeStartGor/xgo_test.go:27 +0x10
		<==]
	*/
}

func Test_Try_Debug(t *testing.T) {
	t.Log(Try_Debug(MustPanicFunc))
	/*output:
	  xgo_test.go:54:
	       | => Gor: 6
	       | => At: xgo.go
	       | => Func: safeStartGor.MustPanicFunc
	       | => Stack:
	      github.com/shenfz/HelloGolang/Go101/GoGorountine/safeStartGor.MustPanicFunc()
	      	E:/GithubGoPath/src/HelloGolang/Go101/GoGorountine/safeStartGor/xgo_test.go:19 +0x10
	       | => Other:
	       runtime error: invalid memory address or nil pointer dereference
	*/
}
