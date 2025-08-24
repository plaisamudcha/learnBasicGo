package main

import (
	"fmt"
)

// func main() {
// 	num := 33
// 	if num%2 == 0 {
// 		fmt.Printf("this number(%d) is even", num)
// 	} else {
// 		fmt.Printf("this number(%d) is odd", num)
// 	}
// }

// func main() {
// 	lim := 225.0
// 	v := math.Pow(10, 3)
// 	if v < lim {
// 		fmt.Println("x power n is: ", v)
// 	} else {
// 		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
// 	}
// }

// func main() {
// 	ratings := 100
// 	if ratings < 5 {
// 		r := '😞'
// 		fmt.Printf("Disappoint %c\n", r)
// 	} else if ratings >= 5 && ratings < 7 {
// 		r := '😐'
// 		fmt.Printf("Normal %c\n", r)
// 	} else if ratings >= 7 && ratings < 10 {
// 		r := '🥰'
// 		fmt.Printf("Good %c\n", r)
// 	} else {
// 		r := '👻'
// 		fmt.Printf("%c%c%c%c%c\n", r, r, r, r, r)
// 	}
// }

// func main() {
// 	today := "Monday"

// 	switch today {
// 	case "Saturday":
// 		fmt.Println("today is Saturday")
// 		fallthrough
// 	case "Monday":
// 		fmt.Println("today is Monday")
// 		fallthrough
// 	case "Tuesday":
// 		fmt.Println("today is Tuesday")
// 	default:
// 		fmt.Printf("hello %v\n", today)
// 	}
// }

func main() {
	ratings := 8.4

	switch {
	case ratings < 5:
		fmt.Println("Disappointed")
	case ratings >= 5 && ratings < 7:
		fmt.Println("Normal")
	case ratings >= 7 && ratings < 10:
		fmt.Println("Good")
	default:
		fmt.Println("I don't know")
	}
}
