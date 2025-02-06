package internal

import (
	"math/rand"
	"strings"

	"github.com/maprost/codeExample/blackjack/internal/cmd"
	"github.com/maprost/codeExample/blackjack/internal/obj"
)

type Game struct {
	screen        cmd.Screen
	BankValue     int
	Stack         obj.Stack
	CardsOfPlayer obj.Stack
}

func NewGame(screen cmd.Screen) Game {
	return Game{
		screen:    screen,
		BankValue: rand.Intn(4) + 17,
		Stack:     obj.NewShuffleStack(),
	}
}

// Setup give the player two cards, he can now finish or get more cards
func (x *Game) Setup() (finished bool) {
	x.screen.Reset()
	if x.takeCard() {
		return true
	}
	if x.takeCard() {
		return true
	}
	return x.showAndAsk()
}

// Round give the player a new card every round
func (x *Game) Round() (finished bool) {
	x.screen.Reset()
	if x.takeCard() {
		return true
	}
	return x.showAndAsk()
}

func (x *Game) CheckResult() {
	x.screen.Reset()

	cardsValue := x.CardsOfPlayer.Number()
	x.screen.Printf("You have %d cards:\n%s\n with a value of %d.\n", len(x.CardsOfPlayer), x.CardsOfPlayer.String(), cardsValue)
	if cardsValue > 21 {
		x.screen.Printf("You have a value of more than 21. You loose.\n")
		return
	}

	x.screen.Printf("The bank has the value: %d\n", x.BankValue)
	if cardsValue > x.BankValue {
		x.screen.Printf("Your value is higher. You win.\n")
		return
	}
	if cardsValue == x.BankValue {
		x.screen.Printf("Your value is equal. You win.\n")
		return
	}

	x.screen.Printf("Your value is lower. You loose.\n")
}

func (x *Game) showAndAsk() (finished bool) {
	cardsValue := x.CardsOfPlayer.Number()
	x.screen.Printf("You have %d cards:\n%s\n with a value of %d.\n", len(x.CardsOfPlayer), x.CardsOfPlayer.String(), cardsValue)

	if cardsValue > 21 {
		return true
	}

	x.screen.Printf("Do you want a new card (y|n):")
	answer := x.screen.Input()
	if strings.TrimSpace(strings.ToLower(answer)) == "n" {
		return true
	}
	return false
}

func (x *Game) takeCard() (finished bool) {
	c, ok := x.Stack.FirstCard()
	if !ok {
		x.screen.Printf("no more cards left, game finished")
		return true
	}
	x.CardsOfPlayer = append(x.CardsOfPlayer, c)
	return false
}
