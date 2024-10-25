package main

import "fmt"

type figures int

const(
	square figures = iota
	circle
	triangle
)

var myFigure figures = square

func Area(figure figures) (func(float64) float64, bool) {
	var ok bool = false

	switch {
	case figure == square:
		ok = true
	case figure == circle:
		ok = true
	case figure == triangle:
		ok = true
	}

	return func(x float64) float64, bool {
		if figure == square {
			return x * x, false
		} else {
			fmt.Println("Ошибочка")
		}
	}, ok
}

ar, ok := Area(myFigure)

if != ok {
	fmt.Println(`Ошибочка`)
	return
}

myArea := ar(x)
