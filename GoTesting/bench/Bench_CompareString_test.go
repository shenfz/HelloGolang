package bench

import (
	"strings"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/12/13 14:14
 * @Email 1328919715@qq.com
 * @Description:
 **/
/*
 比较大小写敏感字符串：  耗时差不多
  1. a = b
  2. strings.Compare(a,b)
 比较大小写忽略字符串： 使用Equalfold较快，比较可以半路下车
  1. strings.Tolower(a) == strings.Tolower(b)
  2. strings.Equalfold(a,b)

*/

//  	614534666	         1.957 ns/op
func Benchmark_Direct_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareDirect("aa", "aa")
	}
}

func CompareDirect(a, b string) bool {
	return a == b
}

//	650660419	         1.847 ns/op
func Benchmark_StringsPKG_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareByStringsPKG("aa", "aa")
	}
}

func CompareByStringsPKG(a string, b string) bool {
	if strings.Compare(a, b) == 0 {
		return true
	}
	return false
}

// 120207826	         9.867 ns/op
func Benchmark_Ignore_Letter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareByIgnoreLetterV1("AcBac", "acbac")
	}
}

func CompareByIgnoreLetterV1(a, b string) bool {
	return strings.EqualFold(a, b)
}

// 37604941	        30.73 ns/op
func Benchmark_IgnoreLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareByIgnoreLetterV2("AcBac", "acbac")
	}
}

func CompareByIgnoreLetterV2(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}
