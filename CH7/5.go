package CH7

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			continue
		} else if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}
