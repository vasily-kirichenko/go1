package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func run_fib() {
	for n := 0; n < 40; n++ {
		fmt.Printf("fib (%d) = %d\n", n, fib(n))
	}
}

/* Crypto hashes */

type digest func([]byte) [sha256.Size]byte

func bench(input *[]byte, name string, digest digest) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		digest(*input)
	}
	fmt.Printf("%s elapled %s", name, time.Since(start))
}

func run_bench() {
	input := make([]byte, 10000000)
	digests := []digest{sha256.Sum256}
	for _, d := range digests {
		bench(&input, "sha256", d)
	}
}

func main() {
	//start := time.Now()
	run_bench()
	//fmt.Printf("Elapsed %s", time.Since(start))
}
