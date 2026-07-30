package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/durationpb"
)

func MiFunction() {
	duracion := durationpb.New(10000000000) // 10 seconds in nanoseconds
	fmt.Println(duracion.AsDuration())
}

func main() {
	fmt.Println("Printing from main")
	MiFunction()
}
