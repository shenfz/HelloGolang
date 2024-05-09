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

func Test_Try(t *testing.T) {
	t.Log(Try(MustPanicFunc, Clear))
	// 2021/07/21 15:00:08 im clear
	//   startGor_test.go:29: PublicByGoV2/StartGoroutineQ.MustPanicFunc ==> : runtime error: invalid memory address or nil pointer dereference
}

func Test_Try2(t *testing.T) {
	Try2(MustPanicFunc2, Clear)
	/*TryGoroutine14:58:34 xgo.go:43: recover.err 'runtime error: invalid memory address or nil pointer dereference' ,line 'F:/MycodeGO/src/PublicByGoV2/StartGoroutineQ/xgo/xgo.go:52'
	  TryGoroutine14:58:34 xgo.go:44: Stack[ ==>
	  		PublicByGoV2/StartGoroutineQ.MustPanicFunc2()
	  F:/MycodeGO/src/PublicByGoV2/StartGoroutineQ/startGor_test.go:21 +0x26
	  	<==]
	  	2021/07/21 14:58:34 im clear*/
}

func Test_Try_Debug(t *testing.T) {
	t.Log(TryDebug(MustPanicFunc, Clear))
	/*output:
	   | => Gor: 6
	   | => LocalAt: xgo.go
	   | => Fun: github.com/shenfz/HelloGolang/GoGorountine/safeStartGor.MustPanicFunc
	   | => Stack:
	  github.com/shenfz/HelloGolang/GoGorountine/safeStartGor.MustPanicFunc(0xf, 0x48db6f)
	  	E:/MyCodeGo/src/github.com/shenfz/HelloGolang/GoGorountine/safeStartGor/xgo_test.go:19 +0x26
	   | => Other:
	   runtime error: invalid memory address or nil pointer dereference

	*/
}
