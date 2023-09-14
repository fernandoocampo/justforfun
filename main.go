package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("I am water", "", "green", true)
	myFigure.Print()

	homdir, err := os.UserHomeDir()
	fmt.Println("homedir", homdir, "err", err)
}
