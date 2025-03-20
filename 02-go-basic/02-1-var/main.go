package main

func zeroValue() (int, string, bool) {
	// ถ้าไม่กำหนดค่า จะได้ค่าเริ่มต้น (Zero value)
	var i int    // 0
	var s string // ""
	var ok bool  // false
	return i, s, ok
}

func initialize() (int, string, bool) {
	// กำหนดค่าตอนประกาศ
	var i int = 20
	var s string = "hello"
	var ok bool = true
	return i, s, ok
}

func main() {
	zeroValue()
	initialize()
}
