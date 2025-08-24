package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	c := cap(sk)
	fmt.Printf("%s: len: %d cap: %d -- show: %v\n", tag, l, c, sk)
}

// func main() {
// 	skills := [3]string{"js", "go", "python"}

// 	s1 := skills[0:2]
// 	show("s1", s1)

// 	s2 := skills[1:3]
// 	show("s2", s2)

// 	s1[1] = "gopher"
// 	show("s1", s1)
// 	show("s2", s2)

// 	fmt.Printf("skills: %#v\n", skills)
// }

// func main() {
// 	skills := []string{"js", "go", "python"}

// 	s1 := skills[0:2]
// 	show("s1", s1)

// 	s2 := skills[1:3]
// 	s2 = append(s2, "c++")
// 	show("s2", s2)

// 	s2[0] = "gopher"
// 	show("s1", s1)
// 	show("s2", s2)

// 	show("skills", skills)

// }

func main() {
	skills := [3]string{"js", "go", "python"}

	show("s1", skills[:])
	// var s1 = make([]string, 3)

}
