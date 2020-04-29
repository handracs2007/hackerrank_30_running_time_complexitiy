// https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	var limit = int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var numOfCases, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < numOfCases; i++ {
		scanner.Scan()
		var number, _ = strconv.Atoi(scanner.Text())

		if isPrime(number) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}
