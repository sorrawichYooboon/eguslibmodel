package main

import (
	"fmt"

	calculator "github.com/sorrawichYooboon/sgslibmodel/calculator"
	numutil "github.com/sorrawichYooboon/sgslibmodel/numutil"
	personinfo "github.com/sorrawichYooboon/sgslibmodel/personinfo"
)

func main() {
	squareArea := calculator.CalSquareArea(5)
	fmt.Println(">>> Square Area: ", squareArea)

	numMax := numutil.FindMax([]int{1, 2, 3, 4, 5})
	fmt.Println(">>> Max: ", numMax)

	firstStudent := personinfo.NewStudent("Sorrawich Yooboon", 25, 86)
	fmt.Println(">>> Student: ", firstStudent.GetStudentInfo())
}
