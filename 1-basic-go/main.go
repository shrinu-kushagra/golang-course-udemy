package main

import "fmt"

func main() {

	fmt.Println("Hello world")
	card := newType()
	fmt.Println("Message: ", card)

	printState()
    loops()
	switchcase()
	recursion()
}

func newType() string {
	return "It's a function.:)"
}
