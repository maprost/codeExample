package obj

import (
	"fmt"
	"math/rand"
)

type Stack []Card

func NewStack() Stack {
	stack := make(Stack, 0, 52)
	for _, color := range AllColors {
		for _, value := range AllValues {
			stack = append(stack, Card{
				Value: value,
				Color: color,
			})
		}
	}
	return stack
}

func NewShuffleStack() Stack {
	stack := NewStack()
	stack.Shuffle()
	return stack
}

func (x *Stack) Shuffle() {
	rand.Shuffle(len(*x), func(i, j int) {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	})
}

func (x *Stack) FirstCard() (Card, bool) {
	if len(*x) == 0 {
		return Card{}, false
	}
	c := (*x)[0]
	*x = (*x)[1:]
	return c, true
}

func (x *Stack) Sep(sep string) string {
	var txt string
	for _, c := range *x {
		if txt != "" {
			txt += sep
		}
		txt += c.ValueSign() + c.ColorSign()
	}
	return txt
}

func (x *Stack) Number() int {
	var n int
	for _, c := range *x {
		n += c.Number()
	}
	return n
}

/*

    ____________________________
   |  ________________________  |
   | |                        | |
   | |                        | |
   | |                        | |
   | |                        | |
   | |                        | |
   | |                        | |
   | |                        | |
   | |                        | |
   | |________________________| |
   |____________________________|

*/

func (x *Stack) String() string {
	var txt string
	const cardSep = "  "

	add := func(s string) {
		var tmp string
		for _ = range *x {
			if tmp != "" {
				tmp += cardSep
			}
			tmp += s
		}
		txt += tmp + "\n"
	}

	addFn := func(fn func(c Card) string) {
		var tmp string
		for _, c := range *x {
			if tmp != "" {
				tmp += cardSep
			}
			tmp += fn(c)
		}
		txt += tmp + "\n"
	}

	singleColorLine := func(c Card) string {
		sign := c.ColorSign()
		return fmt.Sprintf("║      %s      ║", sign)
	}
	doubleColorLine := func(c Card) string {
		sign := c.ColorSign()
		return fmt.Sprintf("║    %s   %s    ║", sign, sign)
	}

	valueLine := func(c Card) string {
		leftSign := c.ValueSign()
		rightSign := leftSign
		if len(leftSign) == 1 {
			leftSign += " "
			rightSign = " " + rightSign
		}
		return fmt.Sprintf("║ %s       %s ║", leftSign, rightSign)
	}

	add("╔═════════════╗")
	addFn(valueLine)
	add("║             ║")
	addFn(singleColorLine)
	addFn(doubleColorLine)
	addFn(singleColorLine)
	add("║             ║")
	addFn(valueLine)
	add("╚═════════════╝")
	return txt
}
