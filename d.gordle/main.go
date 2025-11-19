/*
Gordle - word guessing game
Requirements
- Write a program that picks a random word in a list
- Read the guesses of the player from the standard input
- Give feedback on whether the characters are correctly placed or not
- The player wins if they find, or loses after the maximum number of
  unsuccessful attempts
*/

package main

import (
	"gordle/gordle"
	"os"
)
const maxAttempts = 6;

func main(){
	solution := "hello"
	
	g := gordle.New(os.Stdin, solution, maxAttempts)
	g.Play();
}