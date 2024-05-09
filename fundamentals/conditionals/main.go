package main

func main() {
	// conditionals - if, else, switch, select
	// && - and
	// || - or
	a := 10
	b := 20
	if a > b {
		println("a é maior que b")
		// não tem else if em go
	} else {
		println("b é maior que a")
	}

	// switch
	switch a {
	case 10:
		println("a é 10")
	case 20:
		println("a é 20")
	default:
		println("a não é 10 nem 20")
	}
}
