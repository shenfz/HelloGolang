package bench

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 11:00
 * @Desc:
 */

/*
   1. 空结构体 ：空结构体不占任何内存
   2. 使用场景： map 占位 ， channel 仅作通知 ，仅包含方法的结构体
   3. 一个结构体实例所占据的空间等于各字段占据空间之和，再加上内存对齐的空间大小
   4. 需要字节对齐： cpu 访问内存，并非逐个字节读取，而是按字长（32位cpu读取字长为4字节），增加吞吐量，内存对齐对实现变量的原子性操作也是有好处的，每次内存访问是原子的
   5. 空 struct{} 作为其他 struct 的字段，一般不需要内存对齐.但放在最后一个字段时，需要填充额外的内存保证安全（如未填充，当有指针指向空struct字段, 返回的地址将在结构体之外）
*/

type Door struct{}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}

func Test_SizeOFEmptyStruct(t *testing.T) {
	var (
		i int = 8
	)
	t.Log(unsafe.Sizeof(struct{}{}))  // 0
	t.Log(unsafe.Sizeof(i))           // 8
	t.Log(unsafe.Sizeof(&struct{}{})) // 8
}

type Args struct {
	num1 int // 8
	num2 int // 8
}

type Flag struct {
	num1 int16 // 2  +2 （字节对齐）
	num2 int32 // 4
}

// 对齐
func Test_SizeOfStruct(t *testing.T) {
	fmt.Println(unsafe.Sizeof(Args{})) // 16
	fmt.Println(unsafe.Sizeof(Flag{})) // 8

	// 返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数
	// 对于任意类型的变量 x ，unsafe.Alignof(x) 至少为 1
	// 对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，unsafe.Alignof(x) 等于其中的最大值
	t.Log(unsafe.Alignof(Args{})) // 8
	t.Log(unsafe.Alignof(Flag{})) // 4
}

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

// 空struct   在最后
func Test_EmptyStructAtEnd(t *testing.T) {
	fmt.Println(unsafe.Sizeof(demo3{})) // 8
	fmt.Println(unsafe.Sizeof(demo4{})) // 4
}
