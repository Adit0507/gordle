package gordle

import "fmt"

type Game struct{}

func New() *Game {
	g := &Game{}

	return g
}
func (g *Game) Play() {
	fmt.Println("Welcome to gordle")
	fmt.Printf("Enter a Guess... \n")
}