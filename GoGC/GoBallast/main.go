package main

import (
	"runtime"
)

/**
 * @Author shenfz
 * @Date 2021/12/1 16:21
 * @Email 1328919715@qq.com
 * @Description:
 **/
/*
 Go ballast，其实很简单就是初始化一个生命周期贯穿整个 Go 应用生命周期的超大 slice

ps -eo pmem,comm,pid,maj_flt,min_flt,rss,vsz --sort -rss | numfmt --header --to=iec --field 5 | numfmt --header --from-unit=1024 --to=iec --field 6 | column -t | egrep "[t]est|[P]I"

RSS = 344M
VSZ = 10*1024*1024  M
*/

func main1() {
	ballast := make([]byte, 10*1024*1024*1024) // 10G
	// do something
	runtime.KeepAlive(ballast) // 利用 runtime.KeepAlive 来保证 ballast 不会被 GC 给回收掉
}
