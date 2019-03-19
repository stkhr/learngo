// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// EXAMPLE #1

	name := "Nikola"
	fmt.Println(name)

	// name already exists in this block
	// name := "Marie"

	// just assigns new values to name
	// and declares the new variable age with a value of 66
	name, age := "Marie", 66
	fmt.Println(name, age)

	fmt.Println(name)
	name, on := "hoge", true
	fmt.Println(name, on)
	// EXAMPLE #2

	// name = "Albert"
	// birth := 1879

	// redeclaration below equals to the statements just above
	//
	// `name` is an existing variable
	//   Go just assigns "Albert" to the name variable
	//
	// `birth` is a new variable
	//   Go declares it and assigns it a value of `1879`
	name, birth := "Albert", 1879

	fmt.Println(name, birth)
}
