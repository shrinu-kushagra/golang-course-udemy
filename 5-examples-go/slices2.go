package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))   //here len func returns length of myslice
	fmt.Printf("capacity = %d\n", cap(myslice)) //here cap func returns capacity of myslice
}
