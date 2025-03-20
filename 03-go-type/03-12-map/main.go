package main

import "fmt"

func main() {
	// ถ้าประกาศขึ้นมาลอยๆ จะมีค่าเป็น zero value คือ nil
	var m map[string]string

	m = make(map[string]string)
	// หรือ
	m = map[string]string{}

	// หรือ จะใส่ค่าไปตั้งแต่ประกาศเลยก็ได้
	m = map[string]string{
		"a": "apple",
		"b": "banana", // ต้องปิดด้วย , เสมอ
	}

	// เข้าถึงค่าโดยใช้ key
	fmt.Println(m["a"]) // apple

	// การตรวจสอบ key
	fmt.Printf("%v\n", m["c"]) // ""

	v, ok := m["c"]
	if ok {
		fmt.Println("ค่าของ key c:", v)
	}

	// การแก้ไขข้อมูล
	m["b"] = "berry"

	fmt.Printf("%#v\n", m) // map[string]string{"a":"apple", "b":"berry"}

	// การเพิ่มข้อมูล
	m["c"] = "cranberry"

	fmt.Printf("%#v\n") // map[string]string{"a":"apple", "b":"banana", "c":"cranberry"}

	// การลบข้อมูล
	delete(m, "b")

	fmt.Printf("%#v\n", m) // map[string]string{"a":"apple", "c":"cranberry"}

	// การวนลูปด้วย for range
	m = map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range m {
		fmt.Println(k, v) // จะได้ key แทน index
	}
}
