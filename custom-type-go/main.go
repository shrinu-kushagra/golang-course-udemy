package main

func main() {

	cards := deck{" Ace of Diamond. ", newType()}
	cards = append(cards, " Jack of Heart. ")
	
	cards.print()

}

func newType() string {
	return " Five of Spade."
}
