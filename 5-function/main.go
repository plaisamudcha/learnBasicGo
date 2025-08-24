package main

import "fmt"

// func add(x float64, y float64) (float64, string) {
// 	fmt.Println("hello this is add function", x, y)
// 	return x + y, "hello"
// }

// func swap(x string, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x int, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return //ไม่ต้องประกาศ x,y ก็ได้ ถ้าถูกกำหนดไว้แล้วที่ output แต่อันตราย
// }

// func add2(x, y float64) float64 {
// 	return x + y
// }

// func compute(fn func(float64, float64) float64) float64 {
// 	v := fn(42, 13)
// 	return v
// }

// func hypot(x, y float64) float64 {
// 	return math.Sqrt(x*x + y*y)
// }

// func main() {
// 	result, result2 := add(2.5, 10.5)
// 	fmt.Println(result, result2)
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// 	x, y := split(9)
// 	fmt.Println(x, y)
// 	r := compute(add2)
// 	fmt.Println("r: ", r)

// 	x1 := compute(hypot)
// 	fmt.Println("x1: ", x1)
// }

// func adder() (func() int, func() int) {
// 	sum := 0
// 	return func() int {
// 			sum++
// 			return sum
// 		}, func() int {
// 			return sum
// 		}
// }

// func main() {
// 	i, c := adder()
// 	fmt.Println(i())
// 	fmt.Println(c())
// 	fmt.Println(i())
// 	fmt.Print(c())
// }

////////////////////////////////////////////////////////////////////////
//workshop

func emote(x float64) string {
	if x < 5 {
		return "Disappointed"
	} else if x >= 5 && x < 7 {
		return "Normal"
	} else if x >= 7 && x < 10 {
		return "Good"
	} else {
		return "I don't know"
	}
}

func main() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}
