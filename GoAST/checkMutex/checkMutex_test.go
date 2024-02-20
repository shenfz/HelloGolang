package main

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2022/1/21 10:33
 * @Email 1328919715@qq.com
 * @Description:
 **/

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer, "example.go")
}
