package main

import "fmt"

// func changePrice(p int) {
// 	p = p - 599
// 	fmt.Println("change", p, &p)
// }

// func main() {
// 	var price int = 9999
// 	var addr *int = &price

// 	changePrice(price)
// 	fmt.Println("after", price, addr)
// }

////////////////////////////////////
// workshop

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperhero bool
}

func (m *movie) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}

func main() {
	eg := &movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperhero: true,
	}

	eg.addVote(8)
	fmt.Println("votes:", eg.votes)
}
