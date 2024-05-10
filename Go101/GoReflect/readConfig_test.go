package GoReflect

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/5/10 18:30
 * @Email 1328919715@qq.com
 * @Description:  环境遍历中读取配置  通过反射
 **/

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

func readConfig() *Config {
	conf := Config{}
	typeConf := reflect.TypeOf(conf)
	valConf := reflect.Indirect(reflect.ValueOf(&conf))
	for i := 0; i < typeConf.NumField(); i++ {
		fieldTmp := typeConf.Field(i)
		if v, ok := fieldTmp.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, ex := os.LookupEnv(key); ex {
				valConf.FieldByName(fieldTmp.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &conf
}

// 从环境变量获取配置
func Test_UseReflect_Example1(t *testing.T) {
	os.Setenv("CONFIG_SERVER_NAME", "global_server")
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	os.Setenv("CONFIG_SERVER_URL", "geektutu.com")
	c := readConfig()
	fmt.Printf("%+v", c)
}
