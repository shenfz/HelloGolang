package GoMap

import (
	MySafeMap2 "github.com/shenfz/HelloGolang/GoPrimitive/GoMap/MySafeMap"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 16:39
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  原生map,在删除键值对的时候，并不会真正的删除，而是标记。
  那么随着键值对越来越多，会造成大量内存浪费，甚至导致OOM（OutOfMemory）
*/

/* SafeMap
1.预设一个 删除阈值，如果触发迁移，会放到一个新预设好的 newmap 中
2.两个 map 是一个整体，所以 key 只能留一份
3. 将原先的 dirtyOld 清空，存储的 key/value 通过 for-range 重新存储到 dirtyNew，然后将 dirtyNew 指向 dirtyOld
   [其实在 for-range 过程中，会过滤掉 tophash <= emptyOne 的 key]

*/

func Test_Safemap(t *testing.T) {
	smap := MySafeMap2.NewSafeMap()
	for i := 0; i < 10000; i++ {
		smap.Set(i, i)
	}
}
