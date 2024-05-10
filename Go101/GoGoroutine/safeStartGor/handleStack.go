package safeStartGor

import (
	"bufio"
	"bytes"
	"reflect"
	"runtime"
	"strings"
)

/**
 * @Author shenfz
 * @Date 2024/5/10 17:38
 * @Email 1328919715@qq.com
 * @Description:
 **/

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func ShortFunctionName(i interface{}) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	index := strings.LastIndex(funcName, "/")
	if index <= 0 {
		return funcName
	}
	return funcName[index+1:]
}

func Stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

func HandleStack(info []byte) string {
	respData := bytes.NewBufferString("")
	reader := bufio.NewReader(bytes.NewReader(info))
	for {
		lineD, _, err := reader.ReadLine()
		if err != nil {
			return respData.String()
		}
		if bytes.Contains(lineD, []byte("panic")) {
			break
		}
	}

	line := 3
	for line > 0 {
		lineD, _, err := reader.ReadLine()
		if err != nil {
			return respData.String()
		}
		if line < 3 {
			respData.WriteString("\n")
			respData.Write(lineD)
		}
		line--

	}
	return respData.String()
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
