package main

import "fmt"

func main() {
	if true {
		var x = 100
		fmt.Println(x) // ✅ ใช้งานได้ภายใน if-block
	}
	// fmt.Println(x) ❌ Error: x ใช้ได้เฉพาะใน if-block เท่านั้น

	for i := 1; i <= 3; i++ {
		fmt.Println(i) // ✅ ใช้งานได้ภายใน for-loop
	}
	// fmt.Println(i) ❌ Error: i ใช้ได้เฉพาะใน loop เท่านั้น
}
