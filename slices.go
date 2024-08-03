package main

import "fmt"

func main() {

	cards := []string{" Ace of Diamond. ", newType()}
	fmt.Println("Message: ", cards)

}

func newType() string {
	return " Five of Spade."
}
