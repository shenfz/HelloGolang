package main

/**
 * @Author shenfz
 * @Date 2021/9/30 15:48
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  struct A {
  int  i;
  int  _type;   // 这种情况 会被覆盖，导致 type访问不到
  float type;  // type是GO关键字，使用下划线访问 _type
  float arr[]; // 零长数组也访问不到
  };
  union B1 {
   int i;
   float f;
 };
  enum C {
    ONE,
    TWO,
  };
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 访问结构体
func main12() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a._type)
}

// 使用unsafe包强转 ， 完成对c 联合类型访问
func mainx() {
	var b C.union_B1
	fmt.Println("b.i : ", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Println("b.f : ", *(*C.float)(unsafe.Pointer(&b)))
}

// 枚举类型 C语言中，枚举类型底层对应 int 类型，支持负数类型的值 , 直接访问
func main() {
	var c C.enum_C = C.TWO
	fmt.Println(c)     // 1
	fmt.Println(C.ONE) // 0
	fmt.Println(C.TWO) // 1

}
