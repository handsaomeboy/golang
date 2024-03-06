package main

import "fmt"

type MyError struct {
	Cause error
}

func (m *MyError) Error() string {
	return fmt.Sprintf("error caused by: %s", m.Cause.Error())
}

func WrapError(err error) *MyError {
	if err == nil {
		return nil
	}
	return &MyError{Cause: err}
}

func doSomething() error {
	var err error = underlyingDoSomething()
	// 这里将一个具体的类型值赋值给一个接口，实际也是将这个返回值的
	// type和 value 都写进这个 interface 中
	return WrapError(err)
}
func underlyingDoSomething() error {
	return nil
}

func main() {
	var err = doSomething()
	fmt.Println(err == nil) // 这里为 false
	if err != nil {
		myErr, ok := err.(*MyError)
		if !ok {
			fmt.Println("convert fail")
			return
		}
		fmt.Println(myErr == nil) // 这里为 true
	}
}