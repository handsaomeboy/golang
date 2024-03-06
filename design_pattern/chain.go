package main

import (
	"fmt"
	"reflect"
)

// 责任链模式
type Handler interface {
	Handle(s int)
}

type HandlerOne struct {
	//Handler
}

func (h *HandlerOne) Handle(s int) {
	fmt.Println("handler one:", s)
}

type HandlerTwo struct {
	//Handler
}

func (h *HandlerTwo) Handle(s int) {

}

type HandlerThree struct {
	//Handler
}

func (h *HandlerThree) Handle(s int) {

}

type ChainHandler struct {
	list []Handler
}

func (c *ChainHandler) Handle(s int) {
	for _, handler := range c.list {
		handler.Handle(s)
	}
}

func (c *ChainHandler) AddHandler(h Handler) {
	c.list = append(c.list, h)
}

//  第二种
type HandleFunc func(s int)

type FilterFunc = HandleFunc

func (h HandleFunc) Handle(s int) {
	h(s)
}

type IHandlerFunc func(h Handler, s int)

type Test struct {
	t IHandlerFunc
}

func main() {

	h1 := &HandlerOne{}
	r := reflect.TypeOf(Handler.Handle)
	fmt.Println(r.String())
	//r := reflect.TypeOf(IHandlerFunc)
	t := Test{}
	h := IHandlerFunc(func(h Handler, s int) {
		h.Handle(s)
		fmt.Println("heiheiheiheieh")
	})
	t.t = h
	h(h1,2)
	t.t(h1, 1)
}
