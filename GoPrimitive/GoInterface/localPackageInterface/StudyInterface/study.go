package StudyInterface

/**
 * @Author shenfz
 * @Date 2021/12/9 17:15
 * @Email 1328919715@qq.com
 * @Description:
 **/

type Study interface {
	Listen(message string) string
	i()
}

func Speak(s Study) string {
	return s.Listen("abc")
}

type StudentLocal struct {
	Name string
}

func (s *StudentLocal) Listen(msg string) string {
	return s.Name
}

func (s *StudentLocal) i() {

}
