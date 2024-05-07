package Tree

import (
	"fmt"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 16:13
 * @Desc:
 */
/*

每一个节点本身以及它的后代也是一棵树，是一个递归的结构。
没有后代的节点称为叶子节点，没有节点的树称为空树。

二叉树：每个节点最多只有两个儿子节点的树。

满二叉树：叶子节点与叶子节点之间的高度差为 0 的二叉树，即整棵树是满的，树呈满三角形结构。在国外的定义，非叶子节点儿子都是满的树就是满二叉树。我们以国内为准。

完全二叉树：完全二叉树是由满二叉树而引出来的，设二叉树的深度为 k，除第 k 层外，其他各层的节点数都达到最大值，且第 k 层所有的节点都连续集中在最左边。
*/

/*
 对于一棵有 n 个节点的完全二叉树，从上到下，从左到右进行序号编号，
对于任一个节点，编号 i=0 表示树根节点，
编号 i 的节点的左右儿子节点编号分别为：2i+1,2i+2，
父亲节点编号为：i/2，整除操作去掉小数
*/

/*
 先序遍历：先访问根节点，再访问左子树，最后访问右子树。
后序遍历：先访问左子树，再访问右子树，最后访问根节点。
中序遍历：先访问左子树，再访问根节点，最后访问右子树。
层次遍历：每一层从左到右访问每一个节点。
*/
// 二叉树
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右字树
}

// 先序遍历
func PreOrderPrint(t *TreeNode) {
	if t == nil {
		return
	}
	//root
	fmt.Printf("%s \n", t.Data)
	//left
	PreOrderPrint(t.Left)
	//right
	PreOrderPrint(t.Right)
}

// 中序遍历
func MiddleOrderPrint(t *TreeNode) {
	if t == nil {
		return
	}
	//left
	MiddleOrderPrint(t.Left)
	//root
	fmt.Printf("%s \n", t.Data)
	//right
	MiddleOrderPrint(t.Right)
}

// 后序遍历
func AfterOrderPrint(t *TreeNode) {
	if t == nil {
		return
	}
	AfterOrderPrint(t.Left)
	AfterOrderPrint(t.Right)
	//root
	fmt.Printf("%s \n", t.Data)
}
