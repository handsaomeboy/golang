package main

import (
	"fmt"
	"reflect"
)

type Meta struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
type Response struct {
	Meta *Meta   `json:"meta"`
	Data *Person `json:"data"`
}

type RpcMeta interface {
	GetMeta() *Meta
}

func main() {
	res := &Response{
		Meta: &Meta{
			Code: 10,
			Msg:  "success",
		},
		Data: &Person{
			Name: "test",
			Age:  101,
		},
	}
	objRes := reflect.ValueOf(res).Elem()
	objMeta := objRes.Field(0).Elem()

	//objMeta := objRes.FieldByName("Meta").Elem()
	code := objMeta.Field(0).Int()
	msg := objMeta.Field(1).String()
	fmt.Println(code, msg)

	m1 := GetMeta(res)
	fmt.Println(m1.Msg, m1.Code)

}

func GetMeta(i interface{}) *Meta {
	if m, ok := i.(RpcMeta); ok {
		return m.GetMeta()
	}
	return nil
}
