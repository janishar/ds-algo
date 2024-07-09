package main

import (
	"fmt"
	"time"

	"github.com/janishar/ds-algo/sample/algo"
)

func main() {
	start := time.Now()

	fib := algo.FibItr(47)
	// fib := algo.FibRecurse(47)

	elapsed := time.Since(start)
	fmt.Println(fib)
	fmt.Println(elapsed)
}
