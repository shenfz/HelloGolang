package Chain

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/12 19:37
 * @Desc:
 */

/*
 1. 数组，编程语言高度抽象封装的结构，[下标] 会转换成 [虚拟内存地址]
 2. 数组数据之间是挨着，存放在一个连续的内存区域，每一个固定大小（8字节）的内存片段都有一个【虚拟的地址编号】（程序启动都会有一个虚拟内存空间来映射真正的内存）
 3. 数组的优点是占用空间小，查询快，直接使用索引就可以获取数据元素，缺点是移动和删除数据元素要大量移动空间
 4. 链表的优点是移动和删除数据元素速度快，只要把相关的数据元素重新链接起来，缺点是占用空间大，查找需要遍历
*/

// 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}
func CreateRingChainWithN(N int) *Ring {
	if N <= 0 {
		return nil
	}
	first := new(Ring)
	p := first
	for i := 0; i < N; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = first
	first.prev = p
	return first
}

func (r *Ring) GetNext() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) GetPrev() *Ring {
	if r.prev == nil || r.next == nil {
		return r.init()
	}
	return r.prev
}

// 移动节点，正负区分前后
func (r *Ring) MoveTo(N int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case N > 0:
		for i := 0; i < N; i++ {
			r = r.next
		}
	case N < 0:
		for i := 0; i > N; i-- {
			r = r.prev
		}
	}
	return r
}

// 连接一个节点 ， 返回之前的后节点
func (r *Ring) Link(rn *Ring) *Ring {
	//连接前的后节点
	nextBefore := r.GetNext()
	if rn != nil {
		//传入节点的next,prev为空，则 p==rn
		p := rn.GetPrev()
		r.next = rn
		rn.prev = r
		nextBefore.prev = rn
		//这一步就好分析了
		p.next = nextBefore
	}
	return nextBefore
}

// 解除第n个链表节点
func (r *Ring) UnLink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.MoveTo(n + 1))
}

func (r *Ring) Len() int {
	var length int = 0
	if r != nil {
		//不为空表示至少一个点
		length = 1
		for p := r.GetNext(); p != r; p = p.next {
			length++
		}
	}
	return length
}
