package main

import (
	"fmt"
	"sort"
)

type IPerson interface {
	Hello()
	Sing()
}

type Person1 struct {
	Name string
	Age  int
	IPerson
}

//func (p1 *Person1) Hello() {
//	//p1.IPerson.Hello()
//	fmt.Println("p1 hello")
//}

type Person2 struct {
	Name string
	Age  int
	p    IPerson
}

func (p2 Person2) Hello() {

}

type Man struct {
}

func (m Man) Hello() {
	fmt.Println("man hello")
}

func (m Man) Sing() {
	fmt.Println("man sing")
}

type reverse struct {
	sort.Interface
}

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// 构造reverse Interface
func Reverse(data sort.Interface) sort.Interface {
	return &reverse{data}
}

type Array []int

func (arr Array) Len() int {
	return len(arr)
}
func (arr Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}
func (arr Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	m1 := Man{}
	//p1 := &Person1{
	//	Name: "person1",
	//	Age:  1,
	//	m1}
	p1 := &Person1{
		"person1",
		1,
		m1}
	p1.Hello()

	p2 := Person2{
		Name: "person2",
		Age:  2,
		p:    m1,
	}
	p2.p.Hello()
	arr := Array{1, 2, 3}
	rarr := Reverse(arr)
	fmt.Println(arr.Less(0, 1))
	fmt.Println(rarr.Less(0, 1))
}
