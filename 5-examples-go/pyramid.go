package main

import (
	"fmt"
)

func pyramid() {
	var height int

	
	fmt.Print("Enter the height of the pyramid: ")
	_, err := fmt.Scan(&height)
	if err != nil || height <= 0 {
		fmt.Println("Please enter a valid positive integer for height.")
		return
	}

	
	for i := 1; i <= height; i++ {
		
		for j := i; j < height; j++ {
			fmt.Print(" ")
		}
		
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
