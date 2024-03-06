package main

import "sync"

func main() {
	m := sync.Mutex{}
	cond := sync.NewCond(&m)
	cond.L.Lock()

	cond.L.Unlock()

}
