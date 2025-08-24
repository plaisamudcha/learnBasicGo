package main

import "fmt"

// func show(sk [3]string) {
// 	sk[2] = "ruby"
// 	fmt.Printf("show: %#v\n", sk)
// }

// func main() {
// 	var skills [3]string = [3]string{"js", "go", "python"}
// 	fmt.Printf("%#v\n", skills)

// 	s := skills[1]
// 	fmt.Printf("%#v\n", s)

// 	l := len(skills)
// 	fmt.Printf("%#v\n", l)

// 	skills[1] = "golang"
// 	fmt.Printf("%#v\n", skills)

// 	show(skills)

// 	var xs [3]string
// 	show(xs)
// }

/////////////////////////////////////////
// workshop

func main() {
	genres := [...]string{"Action", "Adventure", "Fantasy"}
	fmt.Println(len(genres))
	fmt.Printf("genres[0]: %v\n", genres[0])
	fmt.Printf("genres[1]: %v\n", genres[1])
	fmt.Printf("genres[2]: %v\n", genres[2])
	genres[1] = "Sci-Fi"
	fmt.Printf("genres[1]: %v\n", genres[1])
}
