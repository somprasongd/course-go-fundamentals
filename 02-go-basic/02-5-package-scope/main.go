package main

import "fmt"

// ตัวแปรระดับแพ็กเกจ
var GlobalVar = "I am visible everywhere in this package"
var localVar = "I am visible only within this package"

func main() {
	fmt.Println(GlobalVar) // ✅ ใช้งานได้
	fmt.Println(localVar)  // ✅ ใช้งานได้
}
