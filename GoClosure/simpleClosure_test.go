package ClosureQ

import (
	"log"
	"os"
	"testing"
	"time"
)

// 返回函数使用了当前函数的局部变量，这种引用环境，结合函数使用，就是闭包
// 流程：1.days分配内存
//      2.定义闭包结构体
//         type closure struct {
//             F         uintptr   // 函数指针，代表内部匿名函数
//             days      *int      // 变量days指针，代表对外部环境的引用 ，检测到闭包环境对该变量‘写’
//             onlyRead  int      // 变量onlyRead值拷贝，检测到闭包对该变量‘只读’ ，避免内存逃逸的发生
//          }
//      3.为闭包结构体对象   申请空间
//      4.通过闭包结构体对象 进行调用
func work() func() (int, int) {
	var days int = 0
	onlyRead := 0
	return func() (int, int) {
		days++
		return days, onlyRead
	}
}

//  闭包调用 ，写和读
func Test_PointClosure(t *testing.T) {
	closureF := work()
	t.Log(closureF()) // 1 0
	t.Log(closureF()) // 2 0
}

func func1() (i int) {
	i = 100
	defer func() {
		i += 1
	}()
	return 5
}

func func2() int {
	var i int
	defer func() {
		i += 1

	}()
	i += 100

	return i
}

// defer  和 闭包结合使用
func Test_DeferAndClosure(t *testing.T) {
	/*
		 首先变量i声明在返回值中，根据go的caller-save模式，变量i会被存储在调用者(main)的栈空间中
		其次defer执行的是一个闭包，闭包中的匿名函数有对外部变量i的写操作，所以闭包结构体中存的是变量i的指针。
		因此func1执行的顺序为先为变量i赋值为100，执行return时会为变量i赋值为5，最后执行defer通过指针将i的值加1，最终返回值为6
	*/
	t.Log(func1()) // 6
	t.Log(func1()) // 6

	/*
	   首先变量i声明在func2中，所以变量i会被放在func2的栈空间中
	  其次defer执行的是一个闭包，闭包中的匿名函数有对外部变量i的写操作，所以闭包结构体中存的是变量i的指针。
	  因此func2的执行顺序为先为变量i赋值为100，再执行return将变量i的值赋给返回值（放在main的栈空间），最后执行defer的闭包将func2栈空间的i加1，最终返回值仍为100
	*/
	t.Log(func2()) // 100
	t.Log(func2()) // 100
}

// 连续 循环 给值闭包
func Test_ForRangeClosure(t *testing.T) {

	// 使用局部变量
	for i := 0; i < 5; i++ {
		// 闭包侦测到外部有 ‘写’ 操作 ，保存了i的指针，最后通过取值的操作打印
		// 如果不每次赋值新变量，则会打印五个 5
		i := i
		defer func() {
			println(i)
		}()
	}

	println("=====")

	// 使用传参
	for i := 6; i < 10; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
	/* 先入后出
	=====
	9
	8
	7
	6
	4
	3
	2
	1
	0
	*/
}

// 循环使用闭包 会导致资源延迟释放
// 这里使用的事局部函数 使用defer释放资源没问题
func Test_RangeClosure(t *testing.T) {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("/path/to/file")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}()
	}
}

func Test_ForSelect(t *testing.T) {
	exitChan := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		exitChan <- "stop"
	}()
	startN := time.Now()
	ForSelect(exitChan)
	t.Logf("Cost== [%v]", time.Since(startN))
}

func ForSelect(exitChan chan string) {
	idieTimeDuration := 2 * time.Second

	idieTimeOnce := time.NewTimer(idieTimeDuration)
	defer idieTimeOnce.Stop()

	idieTimeTricker := time.NewTicker(idieTimeDuration)
	defer idieTimeTricker.Stop()

EXIT:
	for {
		select {
		case <-exitChan:
			break EXIT
		case <-idieTimeOnce.C:
			goto EXIT2
		case <-idieTimeTricker.C:
			goto EXIT3
		case <-time.After(idieTimeDuration):
			goto EXIT4
		default:
		}
	}
	log.Println("exit 1")
	return
EXIT2:
	log.Println("exit 2")
	return
EXIT3:
	log.Println("exit 3")
	return
EXIT4:
	log.Println("exit 4")
	return
}
