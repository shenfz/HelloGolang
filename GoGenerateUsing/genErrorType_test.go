package GoGenerateUsing_test

import (
	"github.com/shenfz/HelloGolang/GoGenerateUsing/genErrorType"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/4/24 16:56
 * @Email 1328919715@qq.com
 * @Description:
 **/

func Test_UsingFuncByGenerate(t *testing.T) {
	t.Log(genErrorType.LogTypeDailyMission)
	t.Logf("%s", genErrorType.LogTypeKnightStarup)
	t.Logf("%+v", genErrorType.LogTypeKnightLayerup)
	t.Log(genErrorType.LogType(9))
}

/* output:
   genErrorType_test.go:16: 日常任务
   genErrorType_test.go:17: 武将升星
   genErrorType_test.go:18: 武将突破
   genErrorType_test.go:19: LogType(9)
*/
