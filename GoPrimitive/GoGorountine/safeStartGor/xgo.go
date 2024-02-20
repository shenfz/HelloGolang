package safeStartGor

/**
 * @Author shenfz
 * @Date 2021/12/20 17:08
 * @Email 1328919715@qq.com
 * @Description:
 **/

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

/*
 Caller报告当前go程调用栈所执行的函数的文件和行号信息。实参skip为上溯的栈帧数，0表示Caller的调用
*/

type PanicInfo struct {
	GorID int64  `json:"gorID"`
	File  string `json:"file"`
	Fun   string `json:"fun"`
	Stack string `json:"stack"`
	Other string `json:"other"`
}

func (p *PanicInfo) String() string {
	return fmt.Sprintf("\n | => Gor: %d \n | => LocalAt: %s \n | => Fun: %s \n | => Stack:%s \n | => Other:\n %s", p.GorID, p.File, p.Fun, p.Stack, p.Other)
}

var _logger = new(log.Logger)

func init() {
	_logger = log.New(os.Stdout, "Xgo", log.Lshortfile|log.Ltime)
}

func TryDebug(fn func() error, cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		_, file, _, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			tmp := PanicInfo{
				GorID: GetGorID(),
				File:  chopPath(file),
				Fun:   FunctionName(fn),
				Stack: HandleStack(Stack()),
			}
			if _, ok := err.(error); ok {
				tmp.Other = err.(error).Error()
			} else {
				tmp.Other = fmt.Sprintf("%+v", err)
			}
			ret = errors.New(tmp.String())
			//TODO print it
			//s,_ := prettyjson.Marshal(&tmp)
			//_logger.Printf(string(s))
		}
	}()
	return fn()
}

func Try(fn func() error, cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
			ret = errors.Wrap(ret, fmt.Sprintf("FuntionName: %s ==> ", FunctionName(fn)))
		}
	}()
	return fn()
}

func Try2(fn func(), cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		_, file, line, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			_logger.Printf("recover.err '%v' ,line '%s'\n", err, fmt.Sprintf("%s:%d", file, line))
			_logger.Printf("Stack[ ==> %s <==]\n", HandleStack(Stack()))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
		}
	}()
	fn()
	return nil
}

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
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
