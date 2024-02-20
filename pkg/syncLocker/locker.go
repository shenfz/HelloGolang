package SyncLocker

import (
	"runtime"
	"sync"
	"sync/atomic"
)

/**
 * @Author shenfz
 * @Date 2021/8/11 15:05
 * @Email 1328919715@qq.com
 * @Description:
 **/

type spinLock uint32

//Lock 加锁 一直尝试比较目标值，非0：一直让出时间 ， 0：原子赋值给1
func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

//Unlock 释放锁 原子赋值0
func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock  创建一个锁
func NewSpinLock() sync.Locker {
	return new(spinLock)
}
