package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibo(x int) int {
	if x < 2 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fibo(n)
	fmt.Printf("\rThe %d fibonacci number is %d\n", n, fibN)
}
