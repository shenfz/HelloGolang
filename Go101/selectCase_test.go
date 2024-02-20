package Go101

import (
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 9:35
 * @Desc:
 */
/*
   1. switch流程控制代码块中的switch表达式的缺省默认值为类型确定值true（其类型为预定义类型bool）
   2. go 插补分号规则导致的问题 ， 可以使用 go fmt 命令来格式化 ， 避免一些格式错误导致的意想不到问题
*/

func Test_Boolean(t *testing.T) {
	switch {
	case true:
		t.Log("缺省值为true") // 缺省值为true
	default:
		t.Log("default")
	case false:
		t.Log("缺省值为false")
	}
}

/*
 在Go代码中，注释除外，如果一个代码行的最后一个语法词段（token）为下列所示之一，则一个分号将自动插入在此字段后（即行尾）：
1. 一个整数、浮点数、虚部、码点或者字符串字面量
2. 这几个跳转关键字之一：break、continue、fallthrough和return
3. 自增运算符++或者自减运算符--    一个右括号：)、]或}
4. 为了允许一条复杂语句完全显示在一个代码行中，分号可能被插入在一个右小括号)或者右大括号}之前

*/
func False() bool {
	return false
}

// 插补分号规则导致的
func Test_BreakLine(t *testing.T) {
	// eg1.
	switch False(); {
	case true:
		t.Log("值为true") //值为true
	default:
		t.Log("default")
	case false:
		t.Log("值为false")
	}

	// eg2. 与eg1等价 ，缺省值为true，所以case true 命中
	switch False(); true {
	case true:
		t.Log("值为true") //值为true
	default:
		t.Log("default")
	case false:
		t.Log("值为false")
	}
}

//func f() {
//	a := 0
//	println(a++) //插分号之后: println(a++;);
//	println(a--) //插分号之后: println(a--;);
//}

func f(x int) {
	switch x {
	case 1:
		{
			goto A
		A: // 这里编译没问题
		}
	case 2:
		goto B
	B:
		; // syntax error: 跳转标签后缺少语句 , 插补规则使然 ， 这里不会自动加分号 ， 需要手动加一个
	case 0:
		goto C
	C: // 这里编译没问题
	}
}

func demo(n, m int) (r int) {
	switch n {
	case 123:
		{
			if m > 0 {
				goto End
			}
			r++

		End: // syntax error: 标签后缺少语句 插补规则使然 也可用{} 显式包括
		}
	default:
		r = 1
	}
	return
}
