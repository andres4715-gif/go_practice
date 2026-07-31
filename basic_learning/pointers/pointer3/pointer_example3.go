package main

import "fmt"

func changingPointerValue(pointer *int) {
	*pointer = 100
}

func main() {
	myVar := 10
	fmt.Println("The original myVar value is:", myVar)
	changingPointerValue(&myVar)
	fmt.Println("The new myVar value is: ", myVar)
}
