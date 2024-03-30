package main

// import our package
// can't direct import like "fmt" since our package is not BUILT-IN function

import (
	"fmt"

	"github.com/DeQatana/cinema1/movie"
	"github.com/DeQatana/cinema1/ticket"
)

/*
package is folder, a box contains
with many elements which has same type or similar duties
*/

// package movie
// func reviewMovie(name string, rating float64) {
// 	fmt.Printf("I reviewed %s and it's rating is %.2f\n", name, rating)
// }

// func findMovieName(imdbID string) string {
// 	switch imdbID {
// 	case "tt4154796":
// 		return "Avengers: Endgame"
// 	case "tt1825683":
// 		return "Black Panther"
// 	}

// 	return "Not found"
// }

// package ticket
// func buyTicket(movie string) {
// 	fmt.Printf("I bought ticket to %s\n", movie)
// }

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4) // package.function
	ticket.BuyTicket(movieName)
}
