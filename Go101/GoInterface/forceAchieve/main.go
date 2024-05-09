package main

/**
 * @Author shenfz
 * @Date 2021/12/9 17:10
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  force `Person` achieve `Man` interface
*/

import "fmt"

var _ Person = (*Man)(nil)

type Man struct {
	Name string
}

type Person interface {
	Listen(message string) string
}

func (m *Man) Listen(msg string) string {
	return m.Name
}

func main() {
	fmt.Println("hello world")
}
