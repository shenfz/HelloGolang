package Stack

import "sync"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 15:34
 * @Desc:
 */

// 切片栈
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func (s *ArrayStack) Push(ele string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//先入放后
	s.array = append(s.array, ele)
	s.size++
}

func (s *ArrayStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.size == 0 {
		return "empty"
	}
	//栈顶
	dstStr := s.array[s.size-1]

	//收缩切片
	s.array = s.array[:s.size-1]
	s.size--

	return dstStr
}

func (s *ArrayStack) Peek() string {
	if s.size == 0 {
		return "empty"
	}

	return s.array[s.size-1]
}
