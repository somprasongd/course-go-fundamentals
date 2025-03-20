package main

import (
	"fmt"
	"runtime"
)

func normal() {
	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("Unknown OS: %s\n", os)
	}
}

func shortStatement() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("Unknown OS: %s\n", os)
	}
}

func switchTrue() {
	os := runtime.GOOS

	switch {
	case os == "darwin":
		fmt.Println("macOS")
	case os == "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("Unknown OS: %s\n", os)
	}
}

func main() {
	normal()
	shortStatement()
	switchTrue()
}
