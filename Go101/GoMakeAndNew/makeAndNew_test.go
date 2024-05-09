package MakeAndNewQ

import (
	"fmt"
	"testing"
)

/*
  make 和 new 关键字区别
*/

/*
 值类型：int，float，bool，string，struct和array.
变量直接存储值，分配栈区的内存空间，这些变量所占据的空间在函数被调用完后会自动释放。

引用类型：slice，map，chan和值类型对应的指针.
变量存储的是一个地址（或者理解为指针），指针指向内存中真正存储数据的首地址。内存通常在堆上分配，通过GC回收。
*/

/* slice切片 底层调用
  type slice struct {
    array  unsafe.Pointer 存储切片数据的指针
    len    int          长度
    cap    int          容量
}
 1.new ,runtime.NewObject() 返回复合类型的指针【*Type】，mallocgc 第一个参数是type.size ，如果传入类型是结构体，只会申请slice结构体的内存！！！
        解引操作会造成panic ,如： *(new([]int))[0] = 1;

 2.make ,汇编用的是 runtime.makeSlice() 返回的是整个复合类型【Type】 ，mallocgc第一个参数是mem，从MulUintptr源码中可以看出mem是slice的容量cap乘以type.size，
  因此使用makeslice可以成功的为切片申请内存
*/
type Student struct {
	Age int // 8
}

//func Test_CompareSomeStruct(t *testing.T) {
//	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})// false   比较的是指针值
//	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"}) // true     无指针 比较的是结构体字段字段
//}

func Test_CompareArrayAndSlice(t *testing.T) {
	fmt.Println([...]string{"1"} == [...]string{"1"}) // 数组类型 可比较
	//fmt.Println([]string{"1"} == []string{"1"})      // 切片类型 不可比较
}

func Test_MapGet(t *testing.T) {
	kv := map[string]Student{"menglu": {Age: 21}}
	s := []Student{{Age: 21}, {Age: 45}}
	a := [...]Student{{Age: 23}, {Age: 90}}
	s[0].Age = 22
	a[0].Age = 24

	bb := a[:]

	fmt.Println(kv)
	fmt.Printf("%p , %p ,%p \n", s, &s[0], &s[1])     // 0xc00000a330 , 0xc00000a330 ,0xc00000a338
	fmt.Printf("%p  , %p ,%p \n", &a, &a[0], &a[1])   // 0xc00000a340  , 0xc00000a340 ,0xc00000a348
	fmt.Printf("%p  , %p ,%p \n", bb, &bb[0], &bb[1]) // 0xc00000a340  , 0xc00000a340 ,0xc00000a348
}
