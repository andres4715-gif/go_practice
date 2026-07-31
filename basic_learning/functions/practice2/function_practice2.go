package main

import (
	"fmt"
	"myProject/basic_learning/functions/practice2/op_mat"
)

func main() {
	totalAddition := op_mat.Add(30, 90)
	totalSubtraction := op_mat.Subtraction(90, 30)
	totalTimes := op_mat.Times(10, 2)
	totalDivided := op_mat.DividedBy(10, 2)
	total := op_mat.TotalAddition(10, 20, 30, 5, 2, 2, 1)

	// Print
	fmt.Println("Total addition is: ", totalAddition)
	fmt.Println("Total subtraction is: ", totalSubtraction)
	fmt.Println("Total times is: ", totalTimes)
	fmt.Println("Total divided is: ", totalDivided)
	fmt.Println("Total full addition: ", total)
}
