package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Game struct {
	reader *bufio.Reader
}

const solutionLength = 5

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesnt exist 🙄")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}

	return nil
}

func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}
func (g *Game) Play() {
	fmt.Println("Welcome to gordle")
	gues := g.ask()

	fmt.Printf("Your guess is %s\n", string(gues))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d character guess: \n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read")
			continue
		}

		guess := []rune(string(playerInput))

		// verification
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s\n", err.Error())
		} else {
			return guess
		}
	}
}
