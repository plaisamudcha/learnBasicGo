package main

import "fmt"

// func main() {
// 	sum := 1
// 	for sum < 5 {
// 		sum += sum
// 		fmt.Println("sum: ", sum)
// 	}
// 	fmt.Println("sum: ", sum)
// }

// func main() {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += i
// 		fmt.Println("i: ", i, "sum: ", sum)
// 	}
// 	fmt.Println("sum: ", sum)
// }

// func main() {
// 	sum := 1
// 	for sum < 5 {
// 		sum += sum
// 		fmt.Println("sum:", sum)
// 	}
// 	fmt.Println("sum:", sum)
// }

// func main() {
// 	skills := [3]string{"js", "go", "python"}

// 	for i := 0; i < len(skills); i++ {
// 		fmt.Println(skills[i])
// 	}

// 	for i := range skills {
// 		fmt.Println(skills[i])
// 	}

// 	for i, val := range skills {
// 		fmt.Println("index:", i, "value:", val)
// 	}

// 	for _, val := range skills {
// 		fmt.Println("value:", val)
// 	}

// }

func main() {
	genre := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop: %#v\n", genre)

	for i := 0; i < len(genre); i++ {
		genre[i] = "Movie: " + genre[i]
	}

	fmt.Printf("after for loop: %#v\n", genre)

	for i, val := range genre {
		fmt.Println("index:", i, "value:", val)
	}
}
