package safeStartGor

import (
	"runtime"
	"strconv"
	"strings"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 17:07
 * @Email 1328919715@qq.com
 * @Description:
 **/

func GetGorID() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	idField := strings.Fields(stk)[0]
	id, _ := strconv.Atoi(idField)
	return int64(id)
}
