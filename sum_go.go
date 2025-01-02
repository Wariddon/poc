package main

import (
	"fmt"
	"time"
)

func main() {
	N := 10000000
	start := time.Now()

	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Go Time: %v seconds\n", time.Since(start).Seconds())
}
