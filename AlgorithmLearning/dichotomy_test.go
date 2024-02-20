package AlgorithmLearning

import "testing"

/**
 * @Author shenfz
 * @Date 2022/7/17 23:42
 * @Email 1328919715@qq.com
 * @Description: 二分法
 **/

/*
 输入: nums = [-1,0,3,5,9,12],
输出: 4
解释: 9 出现在 nums 中并且下标为 4

*/
func Test_Dichotomy(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	tar := 9
	t.Log(DichotomyArrayV1(nums, tar))
	t.Log(DichotomyArrayV2(nums, tar))
}

// 左闭右闭
func DichotomyArrayV1(nums []int, arm int) int {
	if len(nums) == 0 || len(nums) > 3000 {
		return -1
	}
	var (
		left  = 0
		right = len(nums) - 1
	)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > arm {
			right = mid - 1 //TODO here is diff
		} else if nums[mid] < arm {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 因为左闭右开，所以直接赋值middle 不会影响
func DichotomyArrayV2(nums []int, arm int) int {
	if len(nums) == 0 || len(nums) > 3000 {
		return -1
	}
	var (
		left  = 0
		right = len(nums)
	)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > arm {
			right = mid //TODO here is diff
		} else if nums[mid] < arm {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
