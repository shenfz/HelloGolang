package t_test

import (
	"net/url"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/3/25 13:38
 * @Email 1328919715@qq.com
 * @Description: QueryString 和 Fragment解析表现 ,面对多 ？ 和 # 的情况
 **/

var (
	baseURL = "http://example.com/api/v1"
)

func assert(t *testing.T) func(template string, a string, b string, fragment string) {
	return func(template string, a string, b string, fragment string) {

		parsedURL, err := url.Parse(baseURL + template)
		if err != nil {
			t.Fatalf("parse template: %s ,ERR : %s", template, err.Error())
		}
		if queryA := parsedURL.Query().Get("a"); queryA != a {
			t.Fatalf("template[%s]  err a value: [%s] ,expected: [%s] ", template, queryA, a)
		}

		if queryB := parsedURL.Query().Get("b"); queryB != b {
			t.Fatalf("template[%s] err b value: [%s] ,expected: [%s] ", template, queryB, b)
		}

		if fragmentN := parsedURL.EscapedFragment(); fragmentN != fragment {
			t.Fatalf("template[%s] err fragment get: [%s] ,expected:[%s]", template, fragmentN, fragment)
		}
	}
}

func Test_MultipleQuestionMark(t *testing.T) {
	assert(t)(`?a=1?b=2#xxx`, `1?b=2`, ``, `xxx`)
}

func Test_MultipleSameNameKey(t *testing.T) {
	assert(t)(`?a=1&a=2#xxx`, `1`, ``, `xxx`)
}

func Test_multipleAndMark(t *testing.T) {
	assert(t)(`?a=1&&b=2#xxx`, `1`, `2`, `xxx`)
}

func Test_multipleAndMark2(t *testing.T) {
	assert(t)(`?a=1&b=&2&#xxx`, `1`, ``, `xxx`)
}
func Test_multipleFragment(t *testing.T) {
	assert(t)(`?a=1&b=2###xxx##cc#dd`, `1`, `2`, `%23%23xxx%23%23cc%23dd`)
}

func Test_multipleFragment01(t *testing.T) {
	assert(t)(`?a=1#xx?b=2`, `1`, ``, `xx?b=2`)
}

func Test_multipleFragment02(t *testing.T) {
	assert(t)(`?a=1&b=2###xxx##cc#dd`, `1`, `2`, `%23%23xxx%23%23cc%23dd`)
}
