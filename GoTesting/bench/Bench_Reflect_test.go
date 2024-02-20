package bench

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/15 10:38
 * @Desc:
 */

/*
   1. 创建对象： 反射创建则耗时为new的 1.5 倍
   2. 修改字段： 直接赋值只消耗零点几纳秒，Field 和 FieldByName ，前者通过序号赋值，后者则是通过遍历，性能差10倍
   3. 热点代码尽量避免使用reflect包 ， 利用字典将 Name 和 Index 的映射缓存起来。避免每次反复查找，耗费大量的时间

*/
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

// 	46195243	        26.70 ns/op
func BenchmarkNew(b *testing.B) {
	var config *Config
	for i := 0; i < b.N; i++ {
		config = new(Config)
	}
	_ = config
}

// 28686516	        39.53 ns/op
func BenchmarkReflectNew(b *testing.B) {
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}

func BenchmarkSet(b *testing.B) {
	config := new(Config)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.Name = "name"
		config.IP = "ip"
		config.URL = "url"
		config.Timeout = "timeout"
	}
}

func BenchmarkReflect_FieldSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(0).SetString("name")
		ins.Field(1).SetString("ip")
		ins.Field(2).SetString("url")
		ins.Field(3).SetString("timeout")
	}
}

func BenchmarkReflect_FieldByNameSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.FieldByName("Name").SetString("name")
		ins.FieldByName("IP").SetString("ip")
		ins.FieldByName("URL").SetString("url")
		ins.FieldByName("Timeout").SetString("timeout")
	}
}
