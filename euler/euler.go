package euler

import (
	"fmt"
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

func main() {
	fmt.Println(Problem1(10))
}
