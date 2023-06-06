package itp1_1

import (
	"fmt"
	"math"
)

func Itp1_1_b() {
	var x int
	fmt.Scan(&x)
	if 0 <= x && x <= 100 {
		ans := math.Pow(float64(x), 3)
		fmt.Println(int(ans))
	}
}
