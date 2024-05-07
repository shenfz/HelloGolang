package Tree

import (
	"fmt"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 16:35
 * @Desc:
 */
/*
            A
       B         C
   D      E   F

*/
func Test_Print(tr *testing.T) {
	t := new(TreeNode)
	t.Data = "A"
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}

	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}
	fmt.Println("先序排序：")
	PreOrderPrint(t)
	fmt.Println("\n中序排序：")
	MiddleOrderPrint(t)
	fmt.Println("\n后序排序")
	AfterOrderPrint(t)
}
