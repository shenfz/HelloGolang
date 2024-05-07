package Chain

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/12 17:46
 * @Desc:
 */

//一个数据关联到另外一个数据  可以使用链表
//最基础的数据结构

/*
  链表由一个个数据节点组成的，它是一个递归结构，要么它是空的，要么它存在一个指向另外一个数据节点的引用。
*/

// 单链表
type LinkNode struct {
	Data     interface{} //数据
	NextNode *LinkNode   //指向
}

func Test_SingleChainTable(t *testing.T) {

	// 新的节点
	node := new(LinkNode)
	node.Data = 2
	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 // node1 链接到 node 节点上
	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 // node2 链接到 node1 节点上
	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			t.Log(nowNode.Data.(int))
			// 获取下一个节点
			nowNode = nowNode.NextNode
		}
		// 如果下一个节点为空，表示链表结束了
		break
	}
}

func Test_RingChain(t *testing.T) {
	rChain := new(Ring)
	rChain.prev = rChain
	rChain.next = rChain
	rChain.Value = "back to home"
}

func Test_LinkAndUnlink(t *testing.T) {
	deleteTest()
	/*output:
	 -1
	1
	------
	4
	3
	2
	*/
}

func deleteTest() {
	// 第一个节点
	r := &Ring{Value: -1}
	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	//  -1 -> 4 ->3 -> 2 -> 1 -> -1

	//  -1 -> 4 ->3 -> 2 -> ##cut## 1 -> -1
	//   2 -> 4 -> 3 -> 2
	//temp = &Ring{Value:4}
	temp := r.UnLink(3) // 解除了后面两个节点
	// 打印原来的节点

	node := r

	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.GetNext()
		//  如果节点回到了起点，结束
		if node == r {
			break
		}
	}

	fmt.Println("------")
	// 打印被切断的节点
	node = temp
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.GetNext()
		//  如果节点回到了起点，结束
		if node == temp {
			break
		}
	}
}
