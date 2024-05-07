package Heap

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 17:10
 * @Desc:
 */
/*
 优先队列: 是一种能完成以下任务的队列：插入一个数值，取出最小或最大的数值（获取数值，并且删除）。

 1.优先队列可以用二叉树来实现，我们称这种结构为二叉堆。
 2.最小堆和最大堆是二叉堆的一种，是一棵完全二叉树（一种平衡树）
 3.最大堆和最小堆实现方式一样，只不过根节点一个是最大的，一个是最小的
*/
/*【最大堆举例】
  上浮操作：push 一个数入堆，放在树的最尾部，再比较父节点，层层上浮
  下沉操作：pop 一个堆顶数 ,树最尾部的节点放在堆顶，与左右子节点比较，往差值最大方下沉

*/

// 一个最大堆，一棵完全二叉树
// 最大堆要求节点元素都不小于其左右孩子
type Heap struct {
	// 堆的大小
	Size int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

func NewHeap(array []int) *Heap {
	hP := new(Heap)
	hP.Array = array
	hP.Size = len(array)
	return hP
}

func (h *Heap) Push(data int) {
	//成为堆顶
	if h.Size == 0 {
		h.Array[0] = data
		h.Size++
		return
	}
	//待插入的下标
	insertIndex := h.Size
	for insertIndex > 0 {
		//父亲节点
		parentIndex := (insertIndex - 1) / 2
		//父亲节点值大于插入值 满足最大堆规则 break
		if h.Array[parentIndex] >= data {
			break
		}
		//被打败的父节点 互换位置
		h.Array[insertIndex] = h.Array[parentIndex]

		insertIndex = parentIndex
	}
	//放在排好的位置
	h.Array[insertIndex] = data
	h.Size++
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return -1
	}
	//为了之后好删除，先把

}
