package Slice

import "testing"

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/4/14 10:49
 * @Desc:
 */
//切片长度大于 1024 后，会以接近于 1.25 倍进行容量扩容
func Test_INIT(t *testing.T) {
	array := make([]int, 0, 2)
	t.Logf("cap %d len %d ,Data[%v]", cap(array), len(array), array)

	//切片无法原地 append
	_ = append(array, 1)

}
