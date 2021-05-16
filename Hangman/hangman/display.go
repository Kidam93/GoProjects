package hangman

import "fmt"

func Draw(g *Game, letter string) {
	drawState(g)
	drawLetters(g.FoundLetters, letter)
	

}

func drawState(g *Game) {
	draw := ""
	turn := g.TurnLeft
	switch turn {
	case 0:
		draw = "I"
	case 1:
		draw = "II"
	case 2:
		draw = "III"
	case 3:
		draw = "IV"
	case 4:
		draw = "V"
	case 5:
		draw = "VI"
	case 6:
		draw = "VII"
	case 7:
		draw = "VIII"
	case 8:
		draw = "IX"
	default:
		draw = "X"
	}
	fmt.Println(draw)
}

func drawLetters(found []string, letter string) {
	for _, l := range(found) {
		fmt.Print(l)
	}
	fmt.Println("")
}