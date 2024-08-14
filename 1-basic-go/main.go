package main

import "fmt"

func main() {

	fmt.Println("Hello world")
	card := newType()
	fmt.Println("Message: ", card)
	fmt.Println("After adding on github.")
	fmt.Println(newType())

	printState()
    loops()
	switchcase()
}

func newType() string {
	return "It's a function.:)"
}
