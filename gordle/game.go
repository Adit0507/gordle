package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

const solutionLength = 5

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesnt exist 🙄")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}

	return nil
}

func splitToUpperCase(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUpperCase(solution),
		maxAttempts: maxAttempts,
	}

	return g
}
func (g *Game) Play() {
	fmt.Println("Welcome to gordle")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🔥 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("You Lost! The solution was: %s", string(g.solution))
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
