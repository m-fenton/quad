package main

import "fmt"

func QuadE(x, y int) {
	height := y
	width := x

	if height <= 0 || width <= 0 {
		fmt.Println("Error: Height or width of QuadE have a negative number or zero as input and is therefore invalid.")
	} else {

		for h := 1; h <= height; h++ {
			for w := 1; w <= width; w++ {
				if h >= 2 && h < height && w == 1 || h >= 2 && h < height && w == width {
					fmt.Print("B")
				} else if w == 1 && h == 1 {
					fmt.Print("A")
				} else if w == 1 && h == height {
					fmt.Print("C")
				} else if w == h && w == 1 {
					fmt.Print("A")
				} else if w == width && h == 1 {
					fmt.Print("C")
				} else if w == width && h == height {
					fmt.Print("A")
				} else if w == 1 || h == width {
					fmt.Print("C")
				} else if h == 1 || h == height {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
	fmt.Println()
}
