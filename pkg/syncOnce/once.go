package syncOnce

import (
	"sync"
	"sync/atomic"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 17:10
 * @Desc:
 */

/*
   done 放在第一个字段(hot path):
     1. 编译后的机器码指令更少，更直接，必然是能够提升性能
     2. 第一个字段的地址和结构体的指针是相同,直接对结构体的指针解引用即可
     3. 其他字段还需要计算与第一个值的偏移量
     4. 偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址
         因此，访问第一个字段的机器代码更紧凑，速度更快
*/

type Once struct {
	done int32
	m    sync.Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadInt32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreInt32(&o.done, 1)
		f()
	}
}
