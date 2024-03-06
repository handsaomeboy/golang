package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("begin")
	go service()
	time.Sleep(1 * time.Second)
	fmt.Println("end")

}

func service(){
	fmt.Println("service begin..")
	go happenPanic()
	fmt.Println("service end...")
}
func happenPanic() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("defer")
	//	}
	//}()
	fmt.Println("1234")
	panic("happen panic")
	fmt.Println("234")
}
