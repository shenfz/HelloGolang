# test

## bench 命令
```shell
// go test -bench .           : default all
// go test -bench='FIB$'     : get bench example with regexp
// go test -bench='FIB$' test.cpu=2,4 .   : running bench with 2th 4th cpu
// go test -bench='FIB$' -test.benchtime=5s .
// go test -bench="Fib$" -cpuprofile=cpu.pprof .  :  结合命令 go tool pprof -text cpu.pprof
```
