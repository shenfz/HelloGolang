package safeStartGor

/**
 * @Author shenfz
 * @Date 2021/12/20 17:08
 * @Email 1328919715@qq.com
 * @Description:
 **/

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"runtime"
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
	return fmt.Sprintf("\n | => Gor: %d \n | => At: %s \n | => Func: %s \n | => Stack:%s \n | => Other:\n %s", p.GorID, p.File, p.Fun, p.Stack, p.Other)
}

var _logger = new(log.Logger)

func init() {
	_logger = log.New(os.Stdout, "Xgo", log.Lshortfile|log.Ltime)
}

func Try_Debug(fn func() error, hooks ...HookFunction) (ret error) {
	if len(hooks) > 0 {
		for i, _ := range hooks {
			defer hooks[i](context.Background())
		}
	}
	defer func() {
		_, file, _, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			tmp := PanicInfo{
				GorID: GetGorID(),
				File:  chopPath(file),
				Fun:   ShortFunctionName(fn),
				Stack: HandleStack(Stack()),
			}
			if _, ok := err.(error); ok {
				tmp.Other = err.(error).Error()
			} else {
				tmp.Other = fmt.Sprintf("%+v", err)
			}
			ret = errors.New(tmp.String())
		}
	}()
	return fn()
}

func Try(fn func() error, hooks ...HookFunction) (ret error) {
	if len(hooks) > 0 {
		for i, _ := range hooks {
			defer hooks[i](context.Background())
		}
	}
	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
			ret = errors.Wrap(ret, fmt.Sprintf("At: %s ", ShortFunctionName(fn)))
		}
	}()
	return fn()
}

func Try_Stack(fn func(), hooks ...HookFunction) (ret error) {
	if len(hooks) > 0 {
		for i, _ := range hooks {
			defer hooks[i](context.Background())
		}
	}
	defer func() {
		runtime.Caller(5)
		// 	_, file, line, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			_logger.Printf("recover.err => '%v' \n", err)
			_logger.Printf("Stack[ ==> %s \n<==]\n", HandleStack(Stack()))
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
