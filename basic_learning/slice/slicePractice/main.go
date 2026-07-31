package main

import (
	"fmt"
	"myProject/basic_learning/slice/slicePractice/sliceInitialPractice"
)

func main() {
	mySliceWithData := []int{1, 2, 3, 4, 5} // Example slice with data
	result, err := sliceInitialPractice.ReturningSliceWithData(mySliceWithData)

	if err != nil {
		fmt.Println("We are getting an error:", err)
	} else {
		fmt.Println("The final result is: ", result)
	}

	resultDel, errDel := sliceInitialPractice.ReturningSliceWithDeleteData(mySliceWithData)
	if errDel != nil {
		fmt.Println("Delete error:", errDel)
	} else {
		fmt.Println("The final Delete result: ", resultDel)
	}
}
