package main

import "fmt"

/**
 * @Author shenfz
 * @Date 2024/5/15 17:24
 * @Email 1328919715@qq.com
 * @Description: 包级元素的初始化 ： 常量 -> 变量 -> init
 **/
/*
  在go1.22可能会出现 变量优先初始化的表现特征
  fix: 常量初始化放置最前

*/

const (
	c1 = "c1"
	c2 = "c2"
)

var (
	v0 = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 has been initialized")
	}
	if c1 != "" {
		fmt.Println("main: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("main: first init func invoked")
}

func init() {
	fmt.Println("main: second init func invoked")
}

func main() {
	// do nothing
}
