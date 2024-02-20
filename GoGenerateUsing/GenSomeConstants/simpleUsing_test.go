package GenSomeConstants

import (
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/11/19 10:17
 * @Email 1328919715@qq.com
 * @Description:
 **/

func Test_UsingFuncByGenerate(t *testing.T) {
	t.Log(LogTypeDailyMission)
	t.Logf("%s", LogTypeKnightStarup)
	t.Logf("%v", LogTypeKnightLayerup)
}

/* output:
   simpleUsing_test.go:15: 日常任务
   simpleUsing_test.go:16: 武将升星
   simpleUsing_test.go:17: 武将突破
*/
