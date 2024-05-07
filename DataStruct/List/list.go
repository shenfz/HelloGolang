package List

import "sync"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 16:00
 * @Desc:
 */

/*
  一组 地址连续的存储单元 依次存储线性表的数据元素，称为线性表的 顺序存储结构
   一组 任意的存储单元 存储线性表中的数据元素，称为线性表的 链式存储结构
*/
// 双端列表，双端队列
//自实现的与标准库 container/list 不同
type DoubleList struct {
	head *ListNode  // 指向链表头部
	tail *ListNode  // 指向链表尾部
	len  int        // 列表长度
	lock sync.Mutex // 为了进行并发安全pop操作
}

// 列表节点
type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}
