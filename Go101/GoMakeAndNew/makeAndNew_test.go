package MakeAndNewQ__test

/*
  make 和 new 关键字区别
*/

/*
 值类型：int，float，bool，string，struct和array.
变量直接存储值，分配栈区的内存空间，这些变量所占据的空间在函数被调用完后会自动释放。

引用类型：slice，map，chan和值类型对应的指针.
变量存储的是一个地址（或者理解为指针），指针指向内存中真正存储数据的首地址。内存通常在堆上分配，通过GC回收。
*/

/*
	 slice切片 底层调用
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
