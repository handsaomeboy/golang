package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"sync"
)

var wg2 sync.WaitGroup
var nums []int
var l sync.Mutex

func startHttpDebug() {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: ":7890", Handler: pprofHandler}
	go server.ListenAndServe()
	fmt.Println("start go")
	//go http.ListenAndServe(pprofAddr, nil)
}
func main() {
	startHttpDebug()
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go process(ch)
	}
	for i := 0; i < 20; i++ {
		ch <- i
	}
	close(ch)
	wg2.Wait()
	fmt.Println(fmt.Sprintf("main is end,and len is:%d", len(nums)))
}

func process(ch chan int) {
	defer wg2.Done()
	fmt.Println("begin")
	for v := range ch {
		l.Lock()
		nums = append(nums, v)
		fmt.Println(v)
		l.Unlock()
	}
	fmt.Println("end")

}
