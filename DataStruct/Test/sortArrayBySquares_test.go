package Test

import "testing"

/**
 * @Author shenfz
 * @Date 2022/7/18 0:16
 * @Email 1328919715@qq.com
 * @Description: 有序数组的平方排序
 **/

/*
 示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
*/

func TestName(t *testing.T) {
	t.Log()
}

/*
一样采用双指针法，前后比较
*/
func SortSquaresArray(nums []int) []int {
	var (
		res    = make([]int, len(nums))
		left   = 0
		right  = len(nums) - 1
		iPoint = len(nums) - 1
	)
	for left <= right && iPoint > 0 {
		lD, rD := nums[left]*nums[left], nums[right]*nums[right]
		if lD > rD { // 当左平方大于右平方，”大者入位“ ，指针右移
			res[iPoint] = lD
			left++
		} else {
			res[iPoint] = rD
			right--
		}
		iPoint--
	}
	return res
}
