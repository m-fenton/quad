package main

import "fmt"

func QuadA(x, y int) {
	height := y
	width := x

	if height <= 0 || width <= 0 {
		fmt.Println("Error: Height or width of QuadA have a negative number or zero as input and therefore invalid.")
	} else {

		for h := 1; h <= height; h++ {
			for w := 1; w <= width; w++ {
				if h >= 2 && h < height && w == 1 || h >= 2 && h < height && w == width {
					fmt.Print("|")
				} else if w == 1 || w == width {
					fmt.Print("o")
				} else if h == 1 || h == height {
					fmt.Print("-")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
	fmt.Println()
}
