package Go101

import (
	"strings"
	"sync"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 13:54
 * @Desc: 一些不希望被复制的值
 */

/*
  1. strings.Builder 和 bytes,buffer : 前者复制使用panic , 后者不会被go vet 检测到
  2. sync 包里面的东西： 复制sync中类型的值会被go vet命令检测到并被警告

*/

//strings.Builder的实现会在运行时刻探测到非法的strings.Builder值复制。
//一旦这样的复制被发现，就会产生恐慌
func Test_CopyStringsBuilder(t *testing.T) {
	var b strings.Builder
	b.WriteString("hello ")
	var b2 = b
	b2.WriteString("world!") // 一个恐慌将在这里产生
	// 	panic: strings: illegal use of non-zero Builder copied by value
}

func UseCopySyncMutex(locker sync.Mutex) {
	locker.Lock()
	defer locker.Unlock()
	// do something ...
}
