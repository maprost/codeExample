package obj

type Card struct {
	Value Value
	Color Color
}

func (x Card) Number() int {
	return x.Value.Number()
}
func (x Card) ColorSign() string {
	return x.Color.Sign()
}
func (x Card) ValueSign() string {
	return x.Value.Sign()
}
