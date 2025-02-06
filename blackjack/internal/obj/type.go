//go:generate go-enum -f=$GOFILE --names  --lower -t template.tmpl
package obj

import (
	"strings"
)

/*
Color definition:
ENUM(

	Diamond // ♦
	Hearth 	// ♥
	Spade	// ♠
	Club	// ♣

)
*/
type Color int

func (x Color) Sign() string {
	c := ColorCommentMap[x]
	return strings.TrimSpace(c)
}

/*
Value definition:
ENUM(

	One		// 1
	Two		// 2
	Three	// 3
	Four	// 4
	Five	// 5
	Six		// 6
	Seven	// 7
	Eight	// 8
	Nine	// 9
	Ten		// 10
	Jack	// J
	Queen	// Q
	King	// K
	As		// A

)
*/
type Value int

func (x Value) Sign() string {
	c := ValueCommentMap[x]
	return strings.TrimSpace(c)
}

func (x Value) Number() int {
	switch x {
	case ValueOne:
		return 1
	case ValueTwo:
		return 2
	case ValueThree:
		return 3
	case ValueFour:
		return 4
	case ValueFive:
		return 5
	case ValueSix:
		return 6
	case ValueSeven:
		return 7
	case ValueEight:
		return 8
	case ValueNine:
		return 9
	case ValueTen, ValueJack, ValueQueen, ValueKing:
		return 10
	case ValueAs:
		return 11
	default:
		return -1
	}
}
