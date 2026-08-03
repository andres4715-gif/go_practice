package main

import (
	"fmt"
)

type Generic struct {
	Name string
}

func (s *Generic) something() {
	fmt.Println("The employee name is:", s.Name)
}

func main() {
	g := &Generic{Name: "Andres"}
	g.something()
}
