# go gc

## [GoBallAst](./GoBallast)
> 1. 调整GC： GOGC 或者 debug.SetGCPercent()
> 2. GC回收仅仅是标记内存可以返回给操作系统，并不是立即回收，这就是你看到 Go 应用 RSS 一直居高不下的原因
> 3. GOGC 默认值是 100，也就是下次 GC 触发的 heap 的大小是这次 GC 之后的 heap 的一倍。
> 4. 如果是内存占用比较小的程序(触发gc的门槛低)，可能会因为频繁的GC导致耗时不稳定 
> 5. 此方法是为了解决 GOGC 的问题影响或者周期性的耗时不稳定 在 main 注入即可

## [GoKeepLive](./GoKeepLive)
> * 赋予对象 不被gc的权力

## [GoSetFinalizer](./GoSetFinalizer)
> * 被gc的时候触发回调