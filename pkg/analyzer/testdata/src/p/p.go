package p

func returnMap() map[int]int { return nil }

func rangeOverMap() {
	var m map[int]struct{}
	for range m { // want "range iteration over map"

	}

	for a := range m { // want "range iteration over map"
		_ = a
	}

	for range returnMap() { // want "range iteration over map"
	}
}

type M1 = map[int]int

func returnM1() M1 { return nil }

func rangeOverMapAlias() {
	var m M1
	for range m { // want "range iteration over map"

	}

	for a := range m { // want "range iteration over map"
		_ = a
	}

	for range returnM1() { // want "range iteration over map"
	}
}

type M2 map[int]int

func returnM2() M2 { return nil }

func rangeOverMapTypeDef() {
	var m M2
	for range m { // want "range iteration over map"

	}

	for a := range m { // want "range iteration over map"
		_ = a
	}

	for range returnM2() { // want "range iteration over map"
	}
}

func rangeOverArray() {
	var a []int
	for range a {

	}

	for a := range a {
		_ = a
	}

}
