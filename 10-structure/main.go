package main

import "fmt"

// type course struct {
// 	name       string
// 	instructor string
// 	price      float64
// }

// func main() {
// 	c := course{name: "Basic go", instructor: "Samudcha", price: 9999}
// 	c2 := course{"Basic Go", "Samudcha", 9999}
// 	c3 := course{instructor: "Samudcha"}
// 	c4 := course{}

// 	fmt.Println("name:", c.name)
// 	fmt.Println("instructor:", c.instructor)
// 	fmt.Println("price:", c.price)
// 	fmt.Printf("c: %#v\n", c)
// 	fmt.Printf("c2: %#v\n", c2)
// 	fmt.Printf("c3: %#v\n", c3)
// 	fmt.Printf("c4: %#v\n", c4)
// }

///////////////////////////////////////////////

// workshop

type movie struct {
	name        string
	year        int
	rating      float64
	genre       []string
	isSuperhero bool
}

func main() {
	var mvs []movie

	a := movie{
		name:        "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genre:       []string{"Action", "Drama"},
		isSuperhero: true,
	}

	b := movie{
		name:        "Avengers: Infinity War",
		year:        2018,
		rating:      8.4,
		genre:       []string{"Acion", "Sci-Fi"},
		isSuperhero: true,
	}

	mvs = append(mvs, a, b)
	for _, mv := range mvs {
		fmt.Printf("%#v\n", mv)
	}
}
