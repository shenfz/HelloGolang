package Slice

import (
	"fmt"
	"sync"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 10:47
 * @Desc:
 */

type ArrayP struct {
	array []int      // 固定大小的数组，用满容量和满大小的切片来代替
	len   int        // 真正长度
	cap   int        // 容量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func Make(len, cap int) *ArrayP {
	s := new(ArrayP)
	if cap < len {
		panic("len large than cap")
	}
	if cap == 0 {
		cap = 1
	}
	// 把切片当数组用
	ar := make([]int, cap, cap)

	// 元数据
	s.array = ar
	s.len = 0
	s.cap = cap
	return s
}

func (a *ArrayP) Append(ele int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := 2 * a.cap
		newArry := make([]int, newCap, newCap)
		for k, v := range a.array {
			newArry[k] = v
		}
		a.array = newArry
		a.cap = newCap
	}

	a.array[a.len] = ele
	a.len++
}

func (a *ArrayP) AppendMany(data ...int) {
	for _, v := range data {
		a.Append(v)
	}
}

func (a *ArrayP) Len() int {
	return a.len
}

func (a *ArrayP) Cap() int {
	return a.cap
}

func (a *ArrayP) GetByIndex(index int) int {
	if index < 0 || index > a.len {
		panic("index over len")
	}
	return a.array[index]
}

// 辅助打印
func Print(array *ArrayP) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.GetByIndex(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, array.GetByIndex(i))
	}
	result = result + "]"
	return
}
