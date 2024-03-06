package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var w sync.WaitGroup

func main() {
	fmt.Println("main begin")

	for i := 0; i < 10000; i++ {
		stu := &Student{
			20,
			"wang",
		}
		w.Add(3)
		go AddAgeC(stu)
		go AddAgeA(stu)
		go AddAgeB(stu)

		w.Wait()
		stu.String()
	}
	fmt.Println("main end")

}

type Student struct {
	Age  int32
	Name string
}

func (stu *Student) String() {
	fmt.Println(fmt.Sprintf("age is:%d and name is  %s", stu.Age, stu.Name))
}

func AddAgeA(stu *Student) {
	//fmt.Println("AddAgeA begin")
	time.Sleep(1*time.Second)
	res := atomic.AddInt32(&stu.Age, 1)
	fmt.Println(res)
	//fmt.Println("AddAgeA end")
	w.Done()
}

func AddAgeB(stu *Student) {
	//fmt.Println("AddAgeB begin")
	time.Sleep(1*time.Second)
	res := atomic.AddInt32(&stu.Age, 1)
	fmt.Println(res)
	//fmt.Println("AddAgeB end")
	w.Done()

}
func AddAgeC(stu *Student) {
	//fmt.Println("AddAgeC begin")
	//time.Sleep(1*time.Second)
	res := atomic.AddInt32(&stu.Age, 1)
	fmt.Println(res)
	//fmt.Println("AddAgeC end")
	flag:=false
	for; !flag ;  {
		flag =atomic.CompareAndSwapInt32(&stu.Age,stu.Age,30)
		fmt.Println("spin ing -----")
	}
	fmt.Println(flag)
	w.Done()
}
