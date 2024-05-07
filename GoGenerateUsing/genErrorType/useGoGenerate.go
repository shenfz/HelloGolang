package genErrorType

/**
 * @Author shenfz
 * @Date 2021/11/19 10:02
 * @Email 1328919715@qq.com
 * @Description: 使用
 **/

/*
 stringer默认使用常量名称作为描述，使用-linecomment标志可以指定其使用行后的注释。
 stringer生成的文件名默认为类型名小写_string.go，可以通过-output=filename.go标志指定
*/

type LogType uint32

// stringer -type=LogType -output=./genErrorType/genError_string.go -linecomment

const (
	LogTypeDailyMission  LogType = 1 // 日常任务
	LogTypeKnightStarup  LogType = 2 // 武将升星
	LogTypeKnightLayerup LogType = 3 // 武将突破
)
