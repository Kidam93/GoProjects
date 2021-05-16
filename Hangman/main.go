package main

import (
	"os"
	"fmt"
	"projects/hangman"
)

func main() {
	fmt.Println("HANGMAN WELCOME!")
	g := hangman.New("golang", 8)
	fmt.Printf("game struct %v", g)
	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}
	hangman.Draw(g, l)
	fmt.Println(l)
}