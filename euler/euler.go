/**
 * Solutions to Project Euler problem set
 * http://projecteuler.net/problems
 */

package euler

import (
	// "fmt"
)

const (
	PROBLEM1_DEFAULT = 10
	PROBLEM2_DEFAULT = 10
)
	

func Problem1(maxnum int) (int) {
	sum := 0
	for i := 0; i < maxnum; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}
