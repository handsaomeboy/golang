package main

import "fmt"

func add(a int) func(int) int {
	sum := 1
	return func(i int) int {
		return sum + i
	}
}

func main() {
 cal:=add(1 )
 fmt.Println(cal(1))
 fmt.Println(cal(1))
}
