package main

import (
	"fmt"
	"math/big"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	loop := 6000

	wg.Add(loop)

	mf := func(e int64) {
		defer wg.Done()
		//get a cadicate of 2^p-1
		two := big.NewInt(2)
		one := big.NewInt(1)
		p := big.NewInt(e)

		n := big.NewInt(0)
		n.Exp(two, p, nil)
		n.Sub(n, one)

		//fmt.Println(n)
		if n.ProbablyPrime(4) {
			fmt.Printf("\033[100D\033[K2^%d - 1\n", e)
		} else {
			fmt.Printf("\033[100D\033[K2^%d - 1 false", e)
		}

	}

	for i := 0; i < loop; i++ {
		mf(int64(i))
	}
	wg.Wait()
	fmt.Println("done")
}
