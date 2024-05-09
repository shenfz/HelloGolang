package GoGC__test

import (
	"runtime"
	"testing"
	"time"
)

/**
 * @Author shenfz
 * @Date 2024/5/9 15:13
 * @Email 1328919715@qq.com
 * @Description:
 **/

// get last gc time point

func Test_GetLastGc(t *testing.T) {
	// GC 的触发时间可以利用 runtime.Memstats 的 LastGC 来获取
	go doSomeAllocation()
	r := runtime.MemStats{}
	for i := 0; i < 50; i++ {
		time.Sleep(2 * time.Second)
		t.Log("Last Gc : ", time.UnixMicro(int64(r.LastGC)).Format(time.DateTime))
	}
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
