package main

import "fmt"

// func show(val any) {
// 	// i, ok := val.(int)
// 	// if ok {
// 	// 	i += 2
// 	// 	fmt.Println(i)
// 	// }
// 	// j, k := val.(string)
// 	// if k {
// 	// 	j += "hello"
// 	// 	fmt.Println(j)
// 	// }
// 	// i, _ := val.(int)
// 	// i += 2
// 	// fmt.Println(i)

// 	switch v := val.(type) {
// 	case int:
// 		i := v + 2
// 		fmt.Println("int:", i)
// 	case string:
// 		i := v + "hello"
// 		fmt.Println("string:", i)
// 	default:
// 		fmt.Println("not handle type")
// 	}
// }

// func main() {
// 	//var v interface{}
// 	var v any
// 	v = "gopher"
// 	v = 36
// 	v = 36.2
// 	show(v)
// }

// type promotioner interface {
// 	discount() int
// }

// type infoer interface {
// 	info()
// }

// type presenter interface {
// 	promotioner
// 	infoer
// }

// func summary(val presenter) {
// 	fmt.Println("discount price:", val.discount())
// 	val.info()
// }

// func sale(val promotioner) {
// 	fmt.Printf("sale: %#v\n", val.discount())
// }

// func show(val infoer) {
// 	val.info()
// }

// type course struct{}

// func (c course) discount() int {
// 	return 599
// }

// func (c course) info() {
// 	fmt.Println("info:", c)
// }

// func main() {
// 	v := course{}
// 	sale(v)
// 	show(v)
// 	summary(v)
// }

func vote(v voter, rating float64) {
	v.addVote(rating)
}

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

type voter interface {
	addVote(float64)
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

	vote(eg, 8)
	fmt.Println("votes:", eg.votes)
}
