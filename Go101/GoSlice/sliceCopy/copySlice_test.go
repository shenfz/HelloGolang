package sliceCopy

import (
	"reflect"
	"testing"
	"unsafe"
)

/**
 * @Author shenfz
 * @Date 2021/8/5 13:25
 * @Email 1328919715@qq.com
 * @Description:
 **/

// source: https://github.com/lifei6671/interview-go/blob/master/question/q020.md
// 拷贝切片 如何不发生内存拷贝 ？
// 1. 只要是发生类型强转都会发生内存拷贝
// 2. 如果想要在底层转换二者，只需要把 StringHeader 的地址强转成 SliceHeader 就行
// 3. unsafe.Pointer(&a)方法可以得到变量a的地址
// 4. (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
// 5. (*[]byte)(unsafe.Pointer(&ssh)) 可以把底层结构体转成byte的切片的指针。
// 6. 再通过 *转为指针指向的实际内容。

func Test_CopySliceByChangePointer(t *testing.T) {
	str := "aaa"
	t.Logf("Get str[ %v ] Len()=%d Addr= %#p \n", str, len(str), &str)
	pointerTemp := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	bTemp := *(*[]byte)(unsafe.Pointer(&pointerTemp))
	t.Logf("Get Byte[ %v ] Len()=%d Addr= %#p ", bTemp, len(bTemp), &bTemp)
}

/*
   copySlice_test.go:27: Get str[ aaa ] Len()=3 Addr= c000036520
   copySlice_test.go:30: Get Byte[ [97 97 97] ] Len()=3 Addr= c000004090
*/
