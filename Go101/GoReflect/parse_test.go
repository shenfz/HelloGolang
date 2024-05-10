package GoReflect

import (
	"reflect"
	"sync"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2022/1/10 11:23
 * @Email 1328919715@qq.com
 * @Description:
 **/
type Item struct {
	name string `json:"name"`
}

type service struct {
	name   string                 // name of service
	rcvr   reflect.Value          // receiver of methods for the service
	typ    reflect.Type           // type of the receiver
	method map[string]*methodType // registered methods
}

type methodType struct {
	sync.Mutex // protects counters
	method     reflect.Method
	ArgType    reflect.Type
	ReplyType  reflect.Type
	numCalls   uint
}

func Test_GetItem(t *testing.T) {
	var rcvr interface{} = &Item{name: "xxx"}
	s := new(service)
	s.typ = reflect.TypeOf(rcvr)
	s.rcvr = reflect.ValueOf(rcvr)
	// sname := reflect.Indirect(s.rcvr).Type().Name() // Item
	t.Logf("%+v", s)
	/*
	  &{name: rcvr:{typ_:0x955440 ptr:0xc00008a250 flag:22} typ:0x955440 method:map[]}
	*/
}
