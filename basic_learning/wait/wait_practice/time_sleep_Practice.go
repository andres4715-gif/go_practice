package main

import (
	"fmt"
	"time"
)

func main() {
	const nonRoutableIPPolicyTimeout = 1 * time.Minute
	time.Sleep(nonRoutableIPPolicyTimeout)
	fmt.Println("⏰ We are waiting for", nonRoutableIPPolicyTimeout)
}
