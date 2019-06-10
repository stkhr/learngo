// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	type placeholder [5]string

	var(
		zero = placeholder{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		}

		one = placeholder{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		}

		two = placeholder{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		}

		three = placeholder{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		}

		four = placeholder{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		}

		five = placeholder{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		}

		six = placeholder{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		}

		seven = placeholder{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		}

		eight = placeholder{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		}

		nine = placeholder{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		}
	)

	degits := [10][5]string{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	//fmt.Println(len(degits[0]))
	//
	//for i := range degits {
	//	for j := range degits[i] {
	//		fmt.Println(degits[i][j])
	//	}
	//}

	for i := range degits[0] {
		for j := range degits {
			fmt.Print(degits[j][i], "  ")
		}
		fmt.Println()
	}

}
