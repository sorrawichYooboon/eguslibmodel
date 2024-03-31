package main

import (
	"fmt"

	calculator "github.com/sorrawichYooboon/sgslibmodel/calculator"
	numutil "github.com/sorrawichYooboon/sgslibmodel/numutil"
)

func main() {
	a := calculator.CalSquareArea(5)
	fmt.Println(">>> Square Area: ", a)

	b := numutil.FindMax([]int{1, 2, 3, 4, 5})
	fmt.Println(">>> Max: ", b)
}
