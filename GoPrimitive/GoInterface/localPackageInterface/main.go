package main

import (
	StudyInterface2 "github.com/shenfz/HelloGolang/GoPrimitive/GoInterface/localPackageInterface/StudyInterface"
)

/**
 * @Author shenfz
 * @Date 2021/12/9 17:13
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  interface have lowercase method ,like  `i()`, only suit local package struct to achieve it
*/

type StudentOutside struct {
	Name string
}

func (s *StudentOutside) Listen(msg string) string {
	return s.Name
}

func (s *StudentOutside) i() {

}

func main() {
	// StudyInterface.Speak(&StudentOutside{"im outside"})  // wrong
	StudyInterface2.Speak(&StudyInterface2.StudentLocal{Name: "im local"}) // right
}
