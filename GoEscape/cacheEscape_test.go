package T_test

import (
	"log"
	"testing"
	"unsafe"
)

/*
 内存逃逸
*/

/* 典型
1. 在方法内把局部变量指针返回
2. 发送指针或带有指针的值到 channel （编译时，是没有办法知道哪个 goroutine 会在 channel 上接收数据。所以编译器没法知道变量什么时候才会被释放）
3. 在一个切片上存储指针或带指针的值。 一个典型的例子就是 []*string 。这会导致切片的内容逃逸。尽管其后面的数组可能是在栈上分配的，但其引用的值一定是在堆上。
4. slice 的背后数组被重新分配了，因为 append 时可能会超出其容量( cap )。
   （slice 初始化的地方在编译时是可以知道的，它最开始会在栈上分配。如果切片背后的存储要基于运行时的数据进行扩充，就会在堆上分配）
5. 在 interface 类型上调用方法。 在 interface 类型上调用方法都是动态调度的 —— 方法的真正实现只能在运行时知道
  （一个 io.Reader 类型的变量 r , 调用 r.Read(b) 会使得 r 的值和切片b 的背后存储都逃逸掉）
*/

/* 介绍
栈区：主要存储函数的入参、局部变量、出参当中的资源由编译器控制申请和释放。

堆区：内存由程序员自己控制申请和释放，往往存放一些占用大块内存空间的变量，或是存在于函数局部但需供全局使用的变量。
*/

/* 分配情况
 Go的内存分配由编译器决定对象真正的存储位置是在栈上还是在堆上，并管理他的生命周期。
 内存逃逸是指原本应该被存储在栈上的变量，因为一些原因被存储到了堆上

GO即使是用new申请的内存，如果编译器发现new出来的内存在函数结束后就没有使用了，且申请内存空间不是很大(64k)，
那么new申请的内存空间还是会被分配在栈上，毕竟栈访问速度更快且易于管理
   num of byte / 8 / 1024 / 1024 = xx MB
*/

// go build -gcflags="-m" main.go
func main() {
	var a = 123
	escapes(a) // 帮助逃逸函数

	var b = 24
	noescape(unsafe.Pointer(&b)) // 避免逃逸函数
	select {}
}

// 内置类型uintptr是一个真正的指针类型，但是在编译器那里只是一个存储了地址的int值
// 遮蔽输入和输出的依赖关系。使编译器不认为 p 会通过 x 逃逸， 因为 uintptr() 产生的引用是编译器无法理解的
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func escapes(x interface{}) {
	if dummy.b {
		dummy.x = x // 使用了接口类型，肯定逃逸
	}
}

var dummy struct {
	b bool
	x interface{}
}

func testNoUse() {
	testC := new(int)
	*testC = 1
}

//逃逸分析命令： go build -gcflags="-m" main.go
// 外部未使用,未从栈逃逸到堆
func main1() {
	testNoUse()
}

/*
 # command-line-arguments
.\cacheEscape.go:23:6: can inline testNoUse                内联
.\cacheEscape.go:28:6: can inline main
.\cacheEscape.go:29:11: inlining call to testNoUse        内联调用
.\cacheEscape.go:24:12: new(int) does not escape          未逃逸
.\cacheEscape.go:29:11: new(int) does not escape

*/

func testUseOutSide() *int {
	num := 1
	point := &num
	return point
}

/*
 可以看到局部变量num从栈逃逸到了堆上。
 在main函数中对返回的指针point做了 解引用操作，而point指向的变量num如果存储在栈上会在函数showpoint结束时被释放，那么在main函数中也就无法对指针point做解引用的操作了，所以变量num必须要被放在堆上
*/
// 变量在函数外部存在引用，必定放在堆中
func Test_UsingOutSide(t *testing.T) {
	var point *int
	point = testUseOutSide()
	log.Println(*point)
}

/*
 .\cacheEscape.go:24:6: can inline testNoUse
.\cacheEscape.go:29:6: can inline main1
.\cacheEscape.go:30:11: inlining call to testNoUse
.\cacheEscape.go:43:6: can inline testUseOutSide
.\cacheEscape.go:52:26: inlining call to testUseOutSide
.\cacheEscape.go:25:12: new(int) does not escape
.\cacheEscape.go:30:11: new(int) does not escape
.\cacheEscape.go:44:2: moved to heap: num                                 num逃逸到堆
.\cacheEscape.go:53:15: ... argument does not escape
.\cacheEscape.go:53:16: *point escapes to heap

*/

func testSilceWithSize(n int) {
	nums := make([]*int, n)
	a := 1
	nums[0] = &a
}

func testSilce() {
	nums := make([]*int, 2)
	a := 1
	nums[0] = &a
}

/*
 超过64k的内存占用放到堆上
1.频繁的栈扩缩容会导致性能下降
2.make申请的切片大小为 一个变量 时也会在堆上申请内存
*/

/*
 假设这里创建的切片存储了大量的指针，那么对于当中的每一个指针都需要做变量在外部是否被引用的验证，
这样大量的切片取指针，验证操作都会带来性能的损耗，
所以当切片中存储的是指针时，索性将切片中 指针指向的 栈上的变量 全部放到堆上
*/
func Test_Slice(t *testing.T) {
	testSilce()
	//testSilceWithSize(10)
}

/*
.\cacheEscape.go:46:2: moved to heap: num
.\cacheEscape.go:60:15: ... argument does not escape
.\cacheEscape.go:60:16: *point escapes to heap
.\cacheEscape.go:78:2: moved to heap: a
.\cacheEscape.go:77:12: make([]*int, 1) does not escape
.\cacheEscape.go:88:12: moved to heap: a                      指针指向的栈上的变量 全部放在堆上
.\cacheEscape.go:88:12: make([]*int, 1) does not escape        发现无外部引用，切片最后还是没逃逸
*/

/*
.\cacheEscape.go:46:2: moved to heap: num
.\cacheEscape.go:60:15: ... argument does not escape
.\cacheEscape.go:60:16: *point escapes to heap
.\cacheEscape.go:78:2: moved to heap: a                     指针指向的栈上的变量 全部放在堆上
.\cacheEscape.go:77:12: make([]*int, n) escapes to heap      用一个变量来作为切片的大小，这个时候，切片逃逸到堆上
.\cacheEscape.go:94:12: moved to heap: a

*/
