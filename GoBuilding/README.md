# 构建GUI窗口，将控制台隐藏
```shell
  go build -ldflags="-s -w -H=windowsgui" -o  Ws.exe
```

# upx压缩和剔除调试信息
```go
/*
   1. Normal         [9.8M]       : go build -o server main.go
   2. DropDebugInfo  [7.8M]       : go build -ldflags="-s -w" -o server main.go
   3. UPX            [5.0M]       : go build -ldflags="-s -w" -o server main.go && upx -9 server
      upx 压缩等级： 1--9
      upx 原理： 在程序开头或其他合适的地方插入解压代码，将程序的其他部分压缩。执行时完成解压
*/
```
