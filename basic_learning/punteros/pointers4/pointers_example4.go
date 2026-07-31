package main

import (
	"fmt"
)

type Client struct {
	Name string
}

func (t *Client) GetName() string {
	return t.Name
}

func main() {
	p := &Client{Name: "Andres"}
	fmt.Println(p.GetName())
}