package main

import (
	"github.com/maprost/codeExample/blackjack/internal"
	"github.com/maprost/codeExample/blackjack/internal/cmd"
)

func main() {
	screen := cmd.NewScreen()

	screen.Reset()
	screen.Printf("Welcome to Blackjack:\n")
	screen.Printf("1: play\n")
	screen.Printf("2: exit\n")
	screen.Printf("enter: ")
	answer := screen.Input()
	if answer == "1" {
		internal.RunGame(screen)
	}
}
