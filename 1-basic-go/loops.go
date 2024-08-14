package main

import "fmt"

func loops() {
	var n, sum = 10, 0

	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
	}

	fmt.Println("Sum =", sum)
}
