package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds = 4 * time.Second

	fmt.Println("🏃🏼‍♂️ Printing immediately")
	fmt.Println("🤷‍♂️ Waiting for", seconds, "seconds...")
	time.Sleep(seconds)
	fmt.Println("✅ Done waiting")
}
