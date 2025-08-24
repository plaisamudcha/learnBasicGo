package main

import "fmt"

// func main() {
// 	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true
// 	var r = '😀'

// 	fmt.Printf("msg: %s\n", msg)       // string
// 	fmt.Printf("age: %d\n", age)       // decimal
// 	fmt.Printf("price: %.2f\n", price) // float
// 	fmt.Printf("check: %t\n", check)   // boolean
// 	fmt.Printf("r: %#v\n", r)
// 	fmt.Printf("r: %c\n", r) // character

// 	fmt.Printf("all: %#v,%#v,%#v,%#v\n", msg, age, price, check) // %#v ใช้ได้ทุก type
// 	fmt.Printf("type: %T -- msg: %#v\n", msg, msg)
// 	fmt.Printf("type: %T -- msg: %#v\n", age, age)
// 	fmt.Printf("type: %T -- msg: %#v\n", price, price)
// 	fmt.Printf("type: %T -- msg: %#v\n", check, check)
// 	fmt.Printf("type: %T -- msg: %c\n", r, r)
// }

func main() {
	movie := `เรื่อง: Avenger: Endgame
ปั: 2019
เรตติ้ง: 8.4
ประเภท: Sci-Fi
ซุปเปอร์ฮีโร่: true
เรื่องโปรด: '12346'
	`
	fmt.Printf("%s\n", movie)
}
