package bench

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

// go test -bench .           : default all
// go test -bench='FIB$'     : get bench example with regexp
// go test -bench='FIB$' test.cpu=2,4 .   : running bench with 2th 4th cpu
// go test -bench='FIB$' -test.benchtime=5s .
// go test -bench="Fib$" -cpuprofile=cpu.pprof .  :  结合命令 go tool pprof -text cpu.pprof
func Benchmark_FIBO(b *testing.B) {
	// get ready,cost time for prepare
	time.Sleep(3 * time.Second)
	// clean prepare time
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// generate
		b.StopTimer()
		rand.Seed(time.Now().Unix())
		randNum := rand.Intn(10)
		b.StartTimer()
		// start
		fibo(randNum)
	}
}

func Test_NumsOfCPU(t *testing.T) {
	t.Log(runtime.NumCPU()) // 16
}
