package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)

	if v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	fmt.Printf("%g >= %g\n", v, lim) // error: ตัวแปรที่ประกาศใน Short Statement จะใช้ได้เฉพาะในบล็อก if-else เท่านั้น
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
}
