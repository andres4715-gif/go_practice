package main

import (
	"fmt"
	"time"
)

func main() {
	const nonRoutableIPPolicyTimeout = 1 * time.Minute
	fmt.Println("waiting for", nonRoutableIPPolicyTimeout, "to happen")
	time.Sleep(nonRoutableIPPolicyTimeout)
	fmt.Println("Waiting for ", nonRoutableIPPolicyTimeout, "minutes")
}
