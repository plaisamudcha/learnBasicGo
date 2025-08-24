package main

import "fmt"

// type course struct {
// 	name, instructor string
// 	price            int
// }

// func (c course) discount(d int) int {
// 	p := c.price - d
// 	fmt.Println("discount:", p)
// 	return p
// }

// func (c course) info() {
// 	fmt.Println("name:", c.name)
// 	fmt.Println("instructor:", c.instructor)
// 	fmt.Println("price:", c.price)
// }

// func main() {
// 	c := course{"Basic Go", "Samudcha", 9999}
// 	fmt.Printf("%#v\n", c)

// 	d := c.discount(500)
// 	fmt.Println("discount price:", d)

// 	c.info()
// }

//////////////////////////////////////////////////////
// workshop

type movie struct {
	title       string
	year        int
	rating      float64
	genres      []string
	isSuperhero bool
}

func (m movie) info() {
	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.rating)
	fmt.Println("Genres:")
	for _, g := range m.genres {
		fmt.Printf("\t%s\n", g)
	}
}

func main() {
	aw := movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperhero: true,
	}

	aw.info()
}
