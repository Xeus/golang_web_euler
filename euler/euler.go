/**
 * Solutions to Project Euler problem set
 * http://projecteuler.net/problems
 */

package euler

import (
 	"time"
 	"errors"
)

const (
	PROBLEM1_DEFAULT = 10
	PROBLEM2_DEFAULT = 5
	PROBLEM3_DEFAULT = 600851475143
)

// sum of numbers less than 1000 divisible by 3 or 5
func Problem1(maxNum int64) (int64, float64) {
	var i, sum int64
	start := time.Now()
	sum = 0
	for i = 0; i < maxNum; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum, time.Since(start).Seconds()
}

// sum of fibonacci #s <= 4mil that are divisible by 2
func Problem2(maxNum int64) (int64, float64) {
	var sum, sequence, oldSequence int64
	start := time.Now()

	// start conditions
	oldSequence = 0
	sequence = 1
	sum = 0

	for sequence <= maxNum {
		if sequence % 2 == 0 {
			sum += sequence
		}
		sequence, oldSequence = oldSequence, sequence + oldSequence
	}
	return sum, time.Since(start).Seconds()
}

// sum of fibonacci #s <= 4mil that are divisible by 2
func Problem2Alt(maxNum int64) (int64, float64) {
	var sum, sequence, oldSequence, tmp int64
	start := time.Now()

	// start conditions
	oldSequence = 1
	sequence = 2
	sum = 0

	for sequence <= maxNum {
		tmp = sequence  // need a tmp to save sequence while it's changed
		if sequence % 2 == 0 {  // key part of question
			sum += sequence
		}
		sequence += oldSequence
		oldSequence = tmp
	}
	return sum, time.Since(start).Seconds()
}

// largest prime factor of number 600851475143
func Problem3(maxNum int64) (int64, float64, error) {
	var i int64
	start := time.Now()
	if maxNum < 0 {
		return maxNum, time.Since(start).Seconds(), errors.New("negative number")
	}
	for i=2; i * i <= maxNum; i++ {
		for maxNum % i == 0 {
			maxNum = maxNum / i
		}
	}
	return maxNum, time.Since(start).Seconds(), nil
}
