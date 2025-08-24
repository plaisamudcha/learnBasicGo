package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type myError struct{}

func (e myError) Error() string {
	return "myError"
}

var myErr = errors.New("my custom error message")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// err := myError{}
		// err := fmt.Errorf("can't divide '%.2f' by zero", a)
		return 0, myErr
	}
	r := a / b
	return r, nil
}

func main() {
	r, err := divide(1, 0)
	if err != nil {
		fmt.Println("handle error:", err)
		return
	}
	fmt.Println(r, err)
}
