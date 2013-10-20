/**
 * Solutions to Project Euler problem set
 * http://projecteuler.net/problems
 */

package euler

import (
 	"time"
	// "fmt"
)

const (
	PROBLEM1_DEFAULT = 10
	PROBLEM2_DEFAULT = 10
)
	

func Problem1(maxnum int64) (int64, float64) {
	var i, sum int64
	start := time.Now()
	sum = 0
	for i = 0; i < maxnum; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum, time.Since(start).Seconds()
}
