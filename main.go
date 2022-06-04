package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("I am water", "", "green", true)
	myFigure.Print()
}
