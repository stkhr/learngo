// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func main() {
	a, b := 10, 5.5
	a = a - int(b)
	fmt.Println(float64(a) + b)
}
