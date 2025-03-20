package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
		// เท่ากับแบบนี้
		// defer fmt.Println(0)
		// defer fmt.Println(1)
		// defer fmt.Println(2)
	}

	fmt.Println("done")
}

// output
// counting
// done
// 2
// 1
// 0
