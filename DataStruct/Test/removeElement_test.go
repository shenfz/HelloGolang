package Test

import "testing"

/**
 * @Author shenfz
 * @Date 2022/7/17 23:58
 * @Email 1328919715@qq.com
 * @Description: 移除元素
 **/

func Test_Remove(t *testing.T) {
	nums := []int{12, 454, 67, 889, 9, 5649, 5, 8, 9, 35, 8, 9}
	move := 9
	t.Log(RemoveByPoint(nums, move))
}

/*
  移除元素 返回剩余长度
*/

/*
快慢指针 把与移除目标不同的元素 移动 到前方索引 ，统计得来的索引长度就是剩余长度
*/
func RemoveByPoint(nums []int, mVal int) int {
	if len(nums) == 0 || len(nums) > 3000 {
		return -1
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != mVal {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
