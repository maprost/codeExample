package internal

import (
	"strings"

	"github.com/maprost/codeExample/blackjack/internal/cmd"
)

func RunGame(screen cmd.Screen) {
	for {
		GameLoop(screen)

		screen.Printf("\nplay again? (y|n)")
		answer := screen.Input()
		if strings.ToLower(answer) != "y" {
			break
		}
	}
}

func GameLoop(screen cmd.Screen) {
	game := NewGame(screen)

	// do setup
	finished := game.Setup()

	// run rounds
	for !finished {
		finished = game.Round()
	}

	// check result
	game.CheckResult()
}
