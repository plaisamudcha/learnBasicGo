package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var m map[string]int = map[string]int{}
// 	fmt.Println("the value:", m)

// 	m["Answer"] = 42
// 	fmt.Printf("the value: %#v\n", m)

// 	v := m["Answer"]
// 	fmt.Println("the value:", v)

// 	m["Pai"] = 15
// 	fmt.Println("the value:", m)

// 	m["Pai"] = 20
// 	fmt.Println("the value:", m)

// 	delete(m, "Answer")
// 	fmt.Printf("the value: %#v\n", m)

// 	n, ok := m["Answer"]
// 	if ok {
// 		fmt.Println("the value:", n)
// 	} else {
// 		fmt.Println("no")
// 	}
// }

//////////////////////////////////////////////////////////
// workshop

func wordCount(s string) map[string]int {
	w := strings.Fields(s)
	r := map[string]int{}
	for _, val := range w {
		r[val] += 1
	}
	return r
}

func main() {
	s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck"
	w := wordCount(s)
	fmt.Printf("%#v\n", w)
}
