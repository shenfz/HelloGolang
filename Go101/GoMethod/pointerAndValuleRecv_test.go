package GoMethod

import "testing"

/**
 * @Author shenfz
 * @Date 2021/12/20 16:47
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  1.接收者实现了指针类型的方法，（语法糖）隐含地实现接收者是值类型的方法
  2.指针类型方法，无论调用者是对象还是对象指针，可以影响调用者内部参数
  3.指针类型方法，值调用（struct{}）等价于（&struct{}）,指针传参
*/

type PointerRecv struct {
	A int8
	B int8
}

func (v *PointerRecv) GetA() int8 {
	return v.A
}

func (v *PointerRecv) Exchange() {
	v.B, v.A = v.A, v.B
}

func Test_Pointer_Recv(t *testing.T) {
	V := PointerRecv{
		A: 1,
		B: 2,
	}
	//
	V.Exchange()
	t.Logf("%+v", V) // {A:2,B:1}
	(&V).Exchange()
	t.Logf("%+v", V) // {A:1,B:2}
}

/*
  1.接收者实现了值类型的方法
  2.值类型方法，无论调用者是对象还是对象指针，可以不能影响对象内部参数，因为是拷贝的副本
*/

func Test_Value_Recv(t *testing.T) {
	V := ValRecv{
		A: 1,
		B: 2,
	}
	V.Exchange()
	t.Logf("%+v", V) // {A:1,B:2}
	(&V).Exchange()
	t.Logf("%+v", V) // {A:1,B:2}
}

type ValRecv struct {
	A int8
	B int8
}

func (v ValRecv) GetA() int8 {
	return v.A
}

func (v ValRecv) Exchange() {
	v.B, v.A = v.A, v.B
}
