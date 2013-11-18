/**
 * Solutions to Project Euler problem set
 * http://projecteuler.net/problems
 */

package euler

import (
 	"fmt"
 	"strconv"
 	"time"
 	"errors"
 	"math"
)

func ProblemDefaults() (map[int]int64) {
	return map[int]int64 {
		1: 10,
		2: 5,
		3: 600851475143,
		4: 999,
	}
}

func Problem1(maxNum int64) (string, int64, float64) {
	var i, sum int64
	start := time.Now()
	sum = 0
	for i = 0; i < maxNum; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return "sum of numbers less than 1000 divisible by 3 or 5", sum, time.Since(start).Seconds()
}

func Problem2(maxNum int64) (string, int64, float64) {
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
	return "sum of fibonacci #s <= 4mil that are divisible by 2", sum, time.Since(start).Seconds()
}

func Problem2Alt(maxNum int64) (string, int64, float64) {
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
	return "sum of fibonacci #s <= 4mil that are divisible by 2", sum, time.Since(start).Seconds()
}

func Problem3(maxNum int64) (string, int64, float64, error) {
	var i int64
	desc := "largest prime factor of number 600851475143"
	start := time.Now()
	if maxNum < 0 {
		return desc, maxNum, time.Since(start).Seconds(), errors.New("negative number")
	}
	for i=2; i * i <= maxNum; i++ {
		for maxNum % i == 0 {
			maxNum = maxNum / i
		}
	}
	return desc, maxNum, time.Since(start).Seconds(), nil
}

func isPalindrome(product int64) (bool) {
	stringify := strconv.FormatInt(product, 10)
	strLen := len(stringify)
	halfWord := int(math.Floor(float64(strLen / 2)))
	for j := 0; j < halfWord; j++ {
		if (stringify[j] != stringify[strLen - j - 1]) {
			return false
		}
	}
	return true
}

// once the first palindrome is found (probably as the second term i),
// then we can assume no other product components will be lower than
// that number so we can reduce the number of tuples needed to be checked
func findPalindrome(high int64, low int64) (int64, int64) {
	var product, firstLow int64
	var highestProduct int64 = 0
	firstLow = 0
	for h := high; h >= low; h-- {
		for i := high; i >= low; i-- {
			product = h * i

			// don't bother if product is lower than the highest palindrome found
			if product < highestProduct {
				break
			}

			if (isPalindrome(product) == true && product > highestProduct) {
				highestProduct = product
				if (h < i) {
					low = h
					if firstLow == 0 {
						firstLow = h
					}
				} else {
					low = i
					if firstLow == 0 {
						firstLow = i
					}
				}
			}
		}
	}
	return highestProduct, firstLow
}

func Problem4(maxNum int64) (string, int64, float64, string, error) {
	var desc string = "largest palindrome of product of 2 3-digit numbers"
	start := time.Now()

	// error checking
	if maxNum < 0 {
		return desc, maxNum, time.Since(start).Seconds(), "", errors.New("negative number")
	} else if maxNum <= 10 {
		return desc, maxNum, time.Since(start).Seconds(), "", errors.New("number is too low")
	} else if maxNum > 9999999 {
		return desc, maxNum, time.Since(start).Seconds(), "", errors.New("number will take too long to compute")
	}

	var highestProduct int64 = 0
	var high, low, i int64
	var step float64
	var firstLow int64 = 0
	step = 0.1
	high = maxNum
	low = maxNum - int64(math.Floor(float64(maxNum) * step))
	for i = 1; high > firstLow; i++ {
		highestProduct, firstLow = findPalindrome(high, low)
		if (highestProduct != 0) {
			extraInfo := fmt.Sprintf("high: %d, low: %d", high, low)
			return desc, highestProduct, time.Since(start).Seconds(), extraInfo, nil
		}
		high = low
		low = int64(math.Floor(float64(maxNum) * float64(i) * step))
	}
	return desc, maxNum, time.Since(start).Seconds(), "", errors.New("no palindrome found")
}
