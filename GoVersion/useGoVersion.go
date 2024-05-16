package main

import (
	"go/version"
	"log"
	"runtime"
)

/**
 * @Author shenfz
 * @Date 2024/5/16 16:03
 * @Email 1328919715@qq.com
 * @Description: go1.21 新增go/version包 ，定义兼容规则和toolchain ，也能自用
 **/

func main() {
	log.Println(version.Lang(runtime.Version()))
}
