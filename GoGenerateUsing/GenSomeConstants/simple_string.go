// Code generated by "stringer -type=LogType -output=./simple_string.go -linecomment"; DO NOT EDIT.

package GenSomeConstants

import (
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LogTypeDailyMission-1]
	_ = x[LogTypeKnightStarup-2]
	_ = x[LogTypeKnightLayerup-3]
}

const _LogType_name = "日常任务武将升星武将突破"

var _LogType_index = [...]uint8{0, 12, 24, 36}

func (i LogType) String() string {
	i -= 1
	if i >= LogType(len(_LogType_index)-1) {
		return "LogType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _LogType_name[_LogType_index[i]:_LogType_index[i+1]]
}
