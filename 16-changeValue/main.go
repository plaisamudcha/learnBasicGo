package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 256
	fmt.Printf("type:%T, val: %d\n", i, i)

	var f float64 = float64(i)
	fmt.Printf("type:%T, val: %f\n", f, f)

	var u uint8 = uint8(f)
	fmt.Printf("type:%T, val: %d\n", u, u)

	v := "72 xxx"
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	i2 := 10
	n := strconv.Itoa(i2)
	fmt.Println(n)
}
