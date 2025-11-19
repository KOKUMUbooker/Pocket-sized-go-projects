package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of gordle
type Game struct {
	reader *bufio.Reader
	solution []rune
	maxAttempts int 
}

// New returns a Game, which can be used to Play!
func New(playerInput io.Reader,solution string,maxAttempts int) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
		solution: splitToUppercaseCharacters(solution),
		maxAttempts:maxAttempts,
	}
	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!");

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask for a valid word
		guess := g.ask() 

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! 🎉 You found it in %d guess(es)! The word : %s ",currentAttempt,string(g.solution))
			return
		}
	}

	fmt.Printf("You've lost! 😔 The solution was: %s. \n", string(g.solution));
}

const solutionLength = 5

// errInvalidWordLength is returned when the guess has the wrong number of charactes
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the expected number of letters");

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	// propagate error to the caller by returning an error if guess length is not the expected one
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess),errInvalidWordLength);
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string int
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine() 
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess:");
			continue 
		}

		// Convert playerInput(an []byte) into a string first then afterwards 
		// into []rune slice
		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess);
		if err != nil { 
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution length of %d but povided %d ,try again\n",solutionLength,len(guess))
		} else {
			return guess
		}
	}

}