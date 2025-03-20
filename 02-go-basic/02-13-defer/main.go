package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("This will run at the end")

	fmt.Println("Doing something")
}

// ผลลัพธ์
// Start
// Doing something
// This will run at the end
