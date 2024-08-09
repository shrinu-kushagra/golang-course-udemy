package main

import "fmt"

func main() {

	cards := []string{" Ace of Diamond. ", newType()}
	fmt.Println("Message: ", cards)
	for index, card := range cards {
		fmt.Println(index, card)
	}

}

func newType() string {
	return " Five of Spade."
}
