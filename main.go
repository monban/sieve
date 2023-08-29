package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	var max uint64
	var foundPrimes uint64
	p := message.NewPrinter(language.English)
	max = uint64(1_000_000_000)
	bs := NewBoolStore(max)

	fmt.Println()
	for i := uint64(2); i < max; i++ {
		if bs.Get(i) {
			continue
		}
		// fmt.Println(i, "is prime")
		foundPrimes++
		for j := i * 2; j < max; j += i {
			bs.Set(j)
		}
	}

	// fmt.Println(bs)
	p.Println("Found", foundPrimes, "primes less than", max)
}
