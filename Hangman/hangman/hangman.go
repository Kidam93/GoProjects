package hangman

import (
	"strings"
)


type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnLeft     int
}

func New(word string, turn int) *Game {
	letters := strings.Split(strings.ToUpper(word), " ")
	found := make([]string, len(word))
	for i := 0; i < len(word); i++{
		found[i] = "_"
	}

	g := &Game{
		State: "",
		Letters: letters,
		FoundLetters: found,
		UsedLetters: []string{},
		TurnLeft: turn,
	}
	return g
}