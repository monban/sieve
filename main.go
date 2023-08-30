package main

import (
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)
	max := 1_000_000_000
	for n := 10; n <= max; n *= 10 {
		start := time.Now()
		primes := pi(uint64(n))
		p.Printf("Found %d primes less than %d in %s.\n", primes, n, time.Since(start))
	}
}

// pi(n) returns the number of primes less than n
func pi(n uint64) uint64 {
	var count uint64
	bs := NewBoolStore(n)

	for i := uint64(2); i < n; i++ {
		if bs.Get(i) {
			continue
		}
		count++
		for j := i * i; j < n; j += i {
			bs.Set(j)
		}
	}
	return count
}

// findPrimes returns a slice containing all primes less than n
func findPrimes(n uint64) []uint64 {
	var foundPrimes []uint64
	bs := NewBoolStore(n)

	for i := uint64(2); i < n; i++ {
		if bs.Get(i) {
			continue
		}
		foundPrimes = append(foundPrimes, i)
		for j := i * i; j < n; j += i {
			bs.Set(j)
		}
	}
	return foundPrimes
}
