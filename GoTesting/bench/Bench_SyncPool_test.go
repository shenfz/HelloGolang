package bench

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 16:38
 * @Desc:
 */

/*
Sync.Pool :

 1. 可伸缩的，同时也是并发安全的，其大小仅受限于内存的大小

 2. 用于存储那些被分配了但是没有被使用，而未来可能会使用的值

 3. 存放在池中的对象如果不活跃了会被自动清理 （两次GC周期）

    直接使用初始对象指针和使用sync.pool对比 ：

 1. Student 结构体内存占用较小，内存分配几乎不耗时间

 2. 两者耗时相近 ， 主要是json序列化反射耗时占用大头

 3. 内存占用 sync.pool 小了一个数量级

    bytes.Buffer 使用 sync.Pool 对比：

 1. 使用 sync.Pool 无论在耗时和内存占用上面都 表现更佳 ,复用对象 0 B/op
*/
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var (
	poolLocal = sync.Pool{New: func() interface{} {
		return new(Student)
	}}
	buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
)

func Test_SetGet(t *testing.T) {
	stu := poolLocal.Get().(*Student)
	stu.Age = 10
	poolLocal.Put(stu)
}

// 16917	     70545 ns/op	    1384 B/op	       7 allocs/op
func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

// 17061	       69936 ns/op	      232 B/op	       6 allocs/op
func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := poolLocal.Get().(*Student)
		json.Unmarshal(buf, stu)
		poolLocal.Put(stu)
	}
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

// 13985931	        85.61 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset() // buf.Truncate(0)
		bufferPool.Put(buf)
	}
}

// 1590661	       762.7 ns/op	   10240 B/op	       1 allocs/op
func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
