package main

import (
	"net"
	"fmt"
)

func main() {
	ipString := "10.0.0.5"
	fmt.Println("String IP:", ipString)
	ip := net.ParseIP(ipString)

	if ip == nil {
		fmt.Println("invalid IP")
		return
	}

	fmt.Println("Parser IP :", ip)
	fmt.Println("¿The ip is IPv4?", ip.To4() != nil)
}