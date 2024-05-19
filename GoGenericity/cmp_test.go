package GoGenericity

import (
	"cmp"
	"os"
	"slices"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/19 16:15
 * @Email 1328919715@qq.com
 * @Description:
 **/

//  cmp.Or 的主要用途是获取字符串并返回第一个非空白字符串

func Test_CMP_OR(t *testing.T) {
	// 获取环境变量，如果为空就返回默认
	goVersion := cmp.Or(os.Getenv("GO_VERSION"), "default")
	t.Log(goVersion)
}

func Test_Cmp_Compare(t *testing.T) {
	type Order struct {
		Product  string
		Customer string
		Price    float64
	}
	orders := []Order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}
	// Sort by customer first, product second, and last by higher price
	slices.SortFunc(orders, func(a, b Order) int {
		return cmp.Or(
			//	cmp.Compare(a.Customer, b.Customer),
			//	cmp.Compare(a.Product, b.Product),
			cmp.Compare(b.Price, a.Price),
		)
	})
	for i, order := range orders {
		t.Log("index: ", i, " %v ", order)
	}

}

func Default[T comparable](value T, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}

// DefaultWithFunc returns defaultValue if value is zero, otherwise value.
//
//	DefaultWithFunc("", func() string { return "foo" }) // "foo"
//	DefaultWithFunc("bar", func() string { return "foo" }) // "bar"
func DefaultWithFunc[T comparable](value T, defaultValue func() T) T {
	var zero T
	if value == zero {
		return defaultValue()
	}
	return value
}
