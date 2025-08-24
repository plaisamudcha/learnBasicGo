package main

import "fmt"

// func show(sk []string) {
// 	l := len(sk)
// 	fmt.Printf("length: %d -- show: %#v\n", l, sk)
// }

// func main() {
// 	skills := []string{"js", "go", "python"}
// 	show(skills)

// 	genre := []string{"Action", "Adventure", "Fantasy", "Sci-Fi"}
// 	show(genre)

// 	s1 := skills[0:2]
// 	show(s1)
// 	show(skills[0:2])
// 	show(skills[:2])

// 	s2 := skills[1:3]
// 	show(s2)
// 	show(skills[1:3])
// 	show(skills[1:])

// 	l := len(skills)
// 	show(skills[0:l])
// 	show(skills[0:])
// 	show(skills[:l])
// 	show(skills[:])
// }

// func main() {
// 	var ss []int
// 	fmt.Printf("%#v\n", ss)
// 	if ss == nil {
// 		fmt.Println("nil")
// 	}

// 	ss = append(ss, 33)
// 	fmt.Printf("%#v\n", ss)

// 	v := ss[0]
// 	fmt.Printf("%#v\n", v)

// 	ss2 := make([]int, 3)
// 	fmt.Printf("%#v\n", ss2)

// 	ss2[1] = 33
// 	fmt.Printf("%#v\n", ss2)
// }

//////////////////////////////////////////////////////////
// workshop

func main() {
	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	var votes []float64

	votes = append(xs, ys...)

	fmt.Println("votes:", votes)
	fmt.Println("votes 5-8:", votes[5:9])

}
