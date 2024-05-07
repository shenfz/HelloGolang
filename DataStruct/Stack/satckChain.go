package Stack

import "sync"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 15:49
 * @Desc:
 */
// 链表栈，后进先出
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

func (s *LinkStack) Push(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.root == nil {
		s.root = new(LinkNode)
		s.root.Value = str
	} else {
		preNode := s.root
		newRoot := new(LinkNode)
		newRoot.Value = str
		newRoot.Next = preNode
		s.root = newRoot
	}
	s.size++
}

func (s *LinkStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.root == nil {
		return "empty"
	} else {
		popVal := s.root.Value
		s.root = s.root.Next
		s.size--
		return popVal
	}
}

func (s *LinkStack) Peek() string {
	if s.root == nil {
		return "empty"
	} else {
		return s.root.Value
	}
}
