package internal_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maprost/codeExample/blackjack/internal"
	"github.com/maprost/testbox/must"
	"github.com/maprost/testbox/should"
)

type fakeScreen struct {
	inputReturn string
	output      []string
}

func (x *fakeScreen) Printf(msg string, args ...interface{}) {
	x.output = append(x.output, fmt.Sprintf(msg, args...))
}

func (x *fakeScreen) Input() string {
	return x.inputReturn
}

func (x *fakeScreen) Reset() {
	// do nothing here
}

func TestGame(t *testing.T) {
	checkResult := func(t *testing.T, s *fakeScreen, substr string) {
		t.Helper()
		found := false
		for _, o := range s.output {
			if strings.Contains(o, substr) {
				found = true
			}
		}
		should.BeTrue(t, found)
	}
	checkWinResult := func(t *testing.T, s *fakeScreen) {
		t.Helper()
		checkResult(t, s, "You win")
	}
	checkLooseResult := func(t *testing.T, s *fakeScreen) {
		t.Helper()
		checkResult(t, s, "You loose")
	}

	t.Run("check setup - only two cards", func(t *testing.T) {
		screen := &fakeScreen{
			inputReturn: "n",
		}
		game := internal.NewGame(screen)
		game.BankValue = 1 // bank will not win

		finished := game.Setup()
		must.BeTrue(t, finished)
		must.HaveLength(t, game.CardsOfPlayer, 2)

		game.CheckResult()
		if game.CardsOfPlayer.Number() > 21 {
			checkLooseResult(t, screen)
		} else {
			checkWinResult(t, screen)
		}
	})

	t.Run("check setup - more cards", func(t *testing.T) {
		screen := &fakeScreen{
			inputReturn: "y",
		}
		game := internal.NewGame(screen)
		game.BankValue = 1 // bank will not win

		finished := game.Setup()

		must.HaveLength(t, game.CardsOfPlayer, 2)
		game.CheckResult()

		if game.CardsOfPlayer.Number() < 22 {
			must.BeFalse(t, finished)
			checkWinResult(t, screen)
		} else {
			must.BeTrue(t, finished)
			checkLooseResult(t, screen)
		}
	})

	t.Run("check rounds - take as many cards as possible", func(t *testing.T) {
		screen := &fakeScreen{
			inputReturn: "y",
		}
		game := internal.NewGame(screen)
		finished := game.Setup()

		for !finished {
			finished = game.Round()
		}

		must.BeTrue(t, finished)
		must.BeTrue(t, game.CardsOfPlayer.Number() > 21)

		cardsOfPlayerMinusOne := game.CardsOfPlayer[:len(game.CardsOfPlayer)-1]
		must.BeTrue(t, cardsOfPlayerMinusOne.Number() < 22)

		game.CheckResult()
		checkLooseResult(t, screen)
	})
}
