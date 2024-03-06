package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome string

func (w welcome) Do(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

//func main() {
//	persons := make(map[interface{}]interface{})
//	persons["张三"] = 20
//	persons["李四"] = 23
//	persons["王五"] = 26
//
//	var w welcome = "大家好"
//
//	Each(persons, w)
//}

// 接口型函数
type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

/*
还是差不多原来的实现，只是把方法名Do改为selfInfo。HandlerFunc(w.selfInfo)不是方法的调用，
而是转型，因为selfInfo和HandlerFunc是同一种类型，所以可以强制转型。
转型后，因为HandlerFunc实现了Handler接口，所以我们就可以继续使用原来的Each方法了。
*/
type str string

func (s str) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

//func main() {
//	persons := make(map[interface{}]interface{})
//	persons["张三"] = 20
//	persons["李四"] = 23
//	persons["王五"] = 26
//
//	var s str = "大家好"
//
//	Each(persons, HandlerFunc(s.selfInfo))
//
//}

/*
新增了一个EachFunc函数，帮助调用者强制转型，调用者就不用自己做了。
现在我们发现EachFunc函数接收的是一个func(k, v interface{})类型的函数，
没有必要实现Handler接口了，所以我们新的类型可以去掉不用了。
*/
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

func selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)

	a := func(a, b int) {

	}
	a(1, 2)

}
