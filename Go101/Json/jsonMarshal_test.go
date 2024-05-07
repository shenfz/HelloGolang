package Json

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/12/20 16:44
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*解析陷阱
1. 解析取地址 ： 赋值不成功
2. 大小写 ： 内部使用反射，小写的内部成员变量 访问不到
3. 非UTF-8 : 加上反斜杠转义可以成功，或者使用base64编码成字符串.如果不是UTF-8格式，那么Go会用 � (U+FFFD) 来代替无效的 UTF8
4. 数字转interface ： json解析值 默认  float64 ，直接断言int会出错
*/

// 解析出的 原始类型是 float64
func Test_Json_BaseType(t *testing.T) {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}
	armVar, ok := result["status"]
	if !ok {
		t.Fatal("not existed this value")
	}
	typeArmVal := reflect.TypeOf(armVar)
	t.Log(typeArmVal.Kind().String()) // float64
}

type People struct {
	Name string `json:"name"`
}

// 非UTF-8 : 加上反斜杠转义可以成功，或者使用base64编码成字符串.如果不是UTF-8格式，那么Go会用 � (U+FFFD) 来代替无效的 UTF8
func Test_Other(t *testing.T) {

	//	raw1 := []byte(`{"name":"\xc2"}`)  //    invalid character 'x' in string escape code
	raw1 := []byte(`{"name":"\\xc2"}`) //   jsonMarshal_test.go:53: \xc2
	var person People

	if err := json.Unmarshal(raw1, &person); err != nil {
		log.Fatalln(err)
	}
	t.Log(person.Name)
}

/* 比如Age在版本 1 中是int在版本 2 中是string，解析的过程中就会出错
1. 无论反射获得的是哪种类型都会去调用相应的解析接口 UnmarshalJSON
2. 实现此方法 ，解析的时候源码就会调用自实现接口方法
*/

// 定义临时类型。用来接受非json:"_"的字段
type tmp Student

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"_"` // 不解析
}

// 自定义解析规则
func (p *Student) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, p)
}

// 使用外部的 age字段去拦截解析，然后再根据类型做转换，最终给到需要呈现的
func Test_StructChange(t *testing.T) {

	b := []byte(`{"age":"200","name":"tom"}`)
	// b:=[]byte(`{"age":200,"name":"tom"}`)   // age  是200或“200” 结果都对
	var s = &struct {
		tmp // 别的类型 字段填充？
		// interface{}类型，这样才可以接收任意字段
		Age interface{} `json:"age"`
	}{}
	// 解析
	err := json.Unmarshal(b, &s)
	switch da := s.Age.(type) {
	case string:
		var age int
		age, err = strconv.Atoi(da)
		if err != nil {
			t.Fatal(err)
		}
		s.tmp.Age = age
	case float64:
		s.tmp.Age = int(da)
	}

	p := Student(s.tmp)
	t.Log(p)
}
