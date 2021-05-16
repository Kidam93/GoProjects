package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (l string, err error) {
	valid := false
	for !valid {
		fmt.Print("Saisir une lettre? ")
		l, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		l = strings.TrimSpace(l)

		if len(l) != 1{
			fmt.Printf("lettre invalide: %v, taille: %v", l, len(l))
			continue
		}
		valid = true
	}
	return l, err
}
