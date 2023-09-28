package main

import "fmt"

func QuadB(x, y int) {
	height := y
	width := x

	if height <= 0 || width <= 0 {
		fmt.Println("Error: Height or width of QuadB have a negative number or zero as input and is therefore invalid.")
	} else {

		for h := 1; h <= height; h++ {
			for w := 1; w <= width; w++ {
				if (h == 1 && w == 1) || (h == height && w == width) {
					fmt.Print("/")
				} else if (h == 1 && w == width) || (h == height && w == 1) {
					fmt.Print("\\")
				} else if h == 1 || h == height || w == 1 || w == width {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
	fmt.Println()
}
