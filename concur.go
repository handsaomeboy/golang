package main
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex
var rw sync.RWMutex

func main() {

	ch1 := make(chan int, 1)
	defer close(ch1)
	ch2 := make(chan int, 1)
	defer close(ch2)
	ch3 := make(chan int, 1)
	defer close(ch3)
	wg.Add(3)
	fmt.Println("begin")
	ch1 <- 1
	go printA(ch1, ch2)
	go printB(ch2, ch3)
	go printC(ch3, ch1)
	wg.Wait()

	fmt.Println("end")
}

func printA(in1 chan int, in2 chan int) {
	for i := 0; i < 10; i++ {
		select {
		case <-in1:
			fmt.Println("A")
			in2 <- 1
		}
	}


	wg.Done()
}

func printB(in1 chan int, in2 chan int) {
	for i := 0; i < 10; i++ {
		select {
		case <-in1:
			fmt.Println("B")
			in2 <- 1
		}
	}
	wg.Done()

}

func printC(in1 chan int, in2 chan int) {
	for i := 0; i < 10; i++ {
		//ok:=<-in1
		select {
		case ok:=<-in1:
			fmt.Println("C")
			fmt.Println(ok)
			in2 <- 1
		}
	}
	wg.Done()

}
