package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

const (
	pprofAddr string = ":7890"
)

func StartHTTPDebuger() {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: pprofAddr, Handler: pprofHandler}
	go server.ListenAndServe()
	//go http.ListenAndServe(pprofAddr, nil)
}
func main() {
	//StartHTTPDebuger()
	ch := make(chan int)

	ch <- 1
	//var sywg sync.WaitGroup
	//sywg.Wait()

	fmt.Println(1)
}
