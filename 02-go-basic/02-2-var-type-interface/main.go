package main

func typeInterface() (int, string, bool) {
	// ถ้ากำหนดค่าแล้วละ type ได้
	var i = 20
	var s = "hello"
	var ok = true
	return i, s, ok
}

func main() {
	typeInterface()
}
