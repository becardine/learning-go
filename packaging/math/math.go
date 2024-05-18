package math

type Math struct {
	a int
	b int
}

func NewMath(a, b int) Math {
	return Math{a, b}
}

// when a method/struct starts with a lowercase letter, it is private
// when a method/struct starts with an uppercase letter, it is public
func (m *Math) Add() int {
	return m.a + m.b
}
