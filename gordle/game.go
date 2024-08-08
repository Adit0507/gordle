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
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

// verifies every character of the guess against the solution
func computeFeedback(guess, solution []rune) feedback {
	res := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		return res
	}

	// checks for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			res[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// loooks for letters in the wrong position
	for posInGuess, character := range guess {
		if res[posInGuess] != absentCharacter {
			// character has been already marked, ignore it
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// letter of the soln. has already been assinged, skip to the next letter of the soln.
				continue
			}

			if character == target {
				res[posInGuess] = wrongPosition
				used[posInSolution] = true

				// skip to next letter of the guess
				break
			}
		}
	}

	return res
}

func splitToUpperCase(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func New(reader io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0{
		return nil, ErrCorpusIsEmpty
	}
	
	g := &Game{
		reader:      bufio.NewReader(reader),
		solution:    []rune(strings.ToUpper(pickWords(corpus))),
		maxAttempts: maxAttempts,
	}

	return g, nil
}

func (g *Game) Play() {
	fmt.Println("Welcome to gordle")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		fb := computeFeedback(guess, g.solution)

		fmt.Println(fb.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🔥 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("You Lost! The solution was: %s", string(g.solution))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d character guess: \n", len(g.solution))

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
