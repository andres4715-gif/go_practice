package structExample

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) SayHello(friendName string) string {
	fmt.Println(p.Name, "Say hello to", friendName)
	return friendName
}
