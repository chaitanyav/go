/* August 23, 2014
 * Author: NagaChaitanya Vellanki
 * Concurrent Sieve of Eratosthenes
 */

package main

import (
	"fmt"
	"math"
	"sync"
)

const N int64 = 1000

var primes = make([]bool, N)

func computePrimes(num int64, wg *sync.WaitGroup) {
	for i := 2 * num; i <= N; i += num {
		if i%num == 0 && primes[i-1] {
			primes[i-1] = false
		}
	}
	wg.Done()
}

func main() {
	sqrtN := int64(math.Ceil(math.Sqrt(float64(N))))

	for index := range primes {
		primes[index] = true
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go computePrimes(2, &wg)

	for i := int64(3); i <= sqrtN; i += 2 {
		wg.Add(1)
		go computePrimes(i, &wg)
	}

	wg.Wait()
	fmt.Println("Done computing primes less than or equal to", N, " :-)")

	for index := range primes {
		if index > 0 && primes[index] {
			fmt.Println(index + 1)
		}
	}

}
