package main

import (
	"context"
	"github.com/google/uuid"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 17:05
 * @Email 1328919715@qq.com
 * @Description: 本地生产prof文件
 **/

/*
go tool pprof cpu.pprof

go run main.go
go tool pprof -http=:8080 cpu.prof
http://127.0.0.1:8000/debug/pprof/allocs?debug=1

*/
func main() {
	// Setup CPU profiling.
	f, err := os.Create("./cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	ctx := context.Background()
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	// Run workload for tenant1
	iteratePerTenant(ctx, "tenant1", uuid)
	// Run workload for tenant2
	iteratePerTenant(ctx, "tenant2", uuidWithHyphen.String())
}

var (
	iterationsPerTenant = map[string]int{
		"tenant1": 10_000_000_000,
		"tenant2": 1_000_000_000,
	}
)

func iteratePerTenant(ctx context.Context, tenant, uuid string) {
	pprof.Do(ctx, pprof.Labels("tenant", tenant, "uuid", uuid), func(ctx context.Context) {
		iterate(iterationsPerTenant[tenant])
	})
}

func iterate(iterations int) {
	for i := 0; i < iterations; i++ {
	}
}
