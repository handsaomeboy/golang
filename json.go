package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type IdpType int64

type Client struct {
	M map[IdpType]bool `json:"m"`
}

func main() {

	w:=sync.WaitGroup{}
	w.Wait()
	c := Client{}
	m := make(map[IdpType]bool, 10)
	m[IdpType(1)] = false
	m[IdpType(2)] = true
	c.M = m
	str, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(str))

}
