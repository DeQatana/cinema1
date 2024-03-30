// the first file of package SHOULD have the same file name as package/folder

package movie

import "fmt"

func init() {
	fmt.Println("init: movie")
}

/*

func name in package :-
First letter of func name tell are they private or public

reviewMovie :- private, other file can't use

ReviewMovie :- public, allow other file to used
*/

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating is %.2f\n", name, rating)
}
