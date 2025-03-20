# Go Fundamentals (2025)

Source code for Go Fundamentals

Full blog: [here](https://somprasongd.work/blog/go/go-fundamentals-2025)

---

ภาษา Go (หรือ Golang) เป็นภาษาโปรแกรมระดับสูงที่พัฒนาโดย Google โดยเน้นความเรียบง่าย ประสิทธิภาพสูง และรองรับการประมวลผลพร้อมกัน (Concurrency) ได้อย่างยอดเยี่ยม ด้วยความสามารถเหล่านี้ทำให้ Go ได้รับความนิยมอย่างรวดเร็วในหมู่นักพัฒนาซอฟต์แวร์ โดยเฉพาะอย่างยิ่งในระบบที่ต้องการความเร็วและความเสถียรสูง เช่น ระบบเครือข่าย บริการแบบกระจาย (Distributed Systems) และแอปพลิเคชันที่มีการประมวลผลพร้อมกัน

บทความนี้จะพาคุณไปรู้จักกับภาษา Go ตั้งแต่ขั้นพื้นฐานจนถึงแนวคิดขั้นสูง โดยครอบคลุมหัวข้อสำคัญ ได้แก่

- [เริ่มต้นใช้งาน](#1-เริ่มต้นใช้งาน)
- [พื้นฐานภาษา Go](#2-พื้นฐานภาษา-go)
- [ชนิดของข้อมูล](#3-ชนิดของข้อมูล)
- [ออกแบบโค้ดให้ยืดหยุ่นด้วย Interface](#4-ออกแบบโค้ดให้ยืดหยุ่นด้วย-interface)
- [ทำความรู้จักกับ Generic ในภาษา Go](#5-ทำความรู้จักกับ-generic-ในภาษา-go)
- [การทำงานแบบ Concurrency](#6-การทำงานแบบ-concurrency)
- [การจัดการข้อผิดพลาด (Error Handling)](<#7-การจัดการข้อผิดพลาด-(error-handling)>)
- [Advanced concepts ในภาษา Go](#8-advanced-concepts-ในภาษา-go)
- [Unit Testing](#9-unit-testing)

---

## 1. เริ่มต้นใช้งาน

เริ่มจากการติดตั้งภาษา Go และแนะนำการเขียนโปรแกรมเบื้องต้นก่อน

### หัวข้อที่ครอบคลุม

- เตรียมเครื่อง
- การเขียนโปรแกรม "Hello World"
- การรันโปรแกรม
- การใช้งาน Go package

---

### ▶️ เตรียมเครื่อง

- ติดตั้ง Go จาก[ที่นี่](https://go.dev/dl/) (ปัจจุบันเวอร์ชัน 1.24.1)
- ติดตั้ง VS Code จาก[ที่นี่](https://code.visualstudio.com/download) และติดตั้ง extensions เพิ่มเติมดังนี้
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) ต้องติดตั้ง
- [Error Lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens) แนะนำไว้เป็นตัวช่วยแสดงข้อผิดพลาด

### ▶️ การเขียนโปรแกรม "Hello World"

เริ่มต้นโปรเจกต์ใหม่ใน Go จำเป็นต้องสร้าง **Go module** เพื่อเปิดใช้งาน **dependency tracking**

1. **สร้าง Directory สำหรับโปรเจกต์**

```bash
mkdir hello
cd hello
```

2. **สร้าง Go Module**

แนะนำให้ตั้งชื่อ module ตามชื่อ repository หากต้องการแชร์โค้ดให้ผู้อื่นใช้งาน:

```bash
go mod init gobasic
```

3. **โครงสร้างพื้นฐานของโปรแกรม Go**

Go มีข้อกำหนดสำคัญ 2 ข้อสำหรับจุดเริ่มต้นของโปรแกรม:

- ต้องมี `package main`
- โปรแกรมเริ่มทำงานที่ฟังก์ชัน `main`

> 💡 ไฟล์นี้สามารถอยู่ที่ไหนและตั้งชื่ออะไรก็ได้ แต่โดยทั่วไปจะตั้งชื่อเป็น main.go

ตัวอย่างโค้ด "Hello World":

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello world!") // พิมพ์ข้อความออกทาง console
}
```

4. **การใช้งาน Template ในการแสดงผล**

Go รองรับการพิมพ์ข้อความแบบ format ด้วย `fmt.Printf()` ซึ่งมีรูปแบบดังนี้:

```go
package main

import "fmt"

func main() {
    fmt.Printf(`- แสดงข้อความใช้ %s
- แสดงตัวเลขใช้ %d
- แสดงค่าของตัวแปรใช้ %v
- แสดงชนิดของตัวแปรใช้ %T\n`,
        "ข้อความ",     // %s: แสดงข้อความ
        123,          // %d: แสดงตัวเลข
        "ใส่อะไรมาก็ได้", // %v: แสดงค่าของตัวแปร
        true)         // %T: แสดงชนิดของตัวแปร
}
```

**ผลลัพธ์ที่ได้:**

```go
- แสดงข้อความใช้ ข้อความ
- แสดงตัวเลขใช้ 123
- แสดงค่าของตัวแปรใช้ ใส่อะไรมาก็ได้
- แสดงชนิดของตัวแปรใช้ bool
```

### ▶️ การรันโปรแกรม

ใน Go สามารถรันโปรแกรมได้ 2 วิธีหลัก ๆ คือ **รันโดยตรง** และ **คอมไพล์เป็นไฟล์ binary**

1. **รันโปรแกรมโดยตรง**

ใช้คำสั่ง `go run` เพื่อตรวจสอบและรันโค้ดทันที:

```bash
go run main.go
```

  <aside>
💡

เหมาะสำหรับการทดสอบหรือพัฒนา เพราะไม่สร้างไฟล์ executable

  </aside>

2. **คอมไพล์เป็น Binary File**

หากต้องการนำโปรแกรมไปใช้งานหรือแจกจ่าย สามารถคอมไพล์ด้วย `go build`:

```bash
go build main.go
```

จากนั้นรันไฟล์ที่คอมไพล์แล้ว:

```bash
./main   # บน Linux/macOS
main.exe # บน Windows
```

  <aside>
💡

การใช้ `go build` จะสร้างไฟล์ executable ใน directory ปัจจุบัน

  </aside>

### ▶️ การใช้งาน Go package

ในโปรเจกต์จริง การเขียนโค้ดมักไม่ได้รวมอยู่ในไฟล์เดียว แต่จะมีการแยกโค้ดเป็นสัดส่วนเพื่อความเป็นระเบียบ ในภาษา Go สามารถทำได้โดยการสร้าง **package** ซึ่งต้องเริ่มต้นด้วยการสร้าง **Go Module** ก่อน

การสร้าง package สามารถทำได้โดยการแบ่งตาม **directory structure** ซึ่งมีขั้นตอนง่าย ๆ ดังนี้:

1. **สร้าง Directory สำหรับ Package**

ตั้งชื่อ directory ให้ตรงกับชื่อ package ที่ต้องการ เช่น หากต้องการสร้าง `package greet` ให้สร้าง directory ชื่อ `greet`:

```go
mkdir greet
```

2. **สร้างไฟล์แรกของ Package**

แนะนำให้ตั้งชื่อไฟล์แรกให้ตรงกับชื่อ package เพื่อความชัดเจน เช่น `greet.go`:

```bash
cd greet
touch greet.go
```

3. **เขียนโค้ดใน Package**

   ใน Go หากต้องการให้ **ตัวแปร (Variable)** หรือ **ฟังก์ชัน (Function)** สามารถเรียกใช้งานจาก package อื่นได้ ต้องตั้งชื่อขึ้นต้นด้วย **ตัวอักษรพิมพ์ใหญ่ (Exported)**:

```go
// greet/greet.go
package greet

import "fmt"

// ฟังก์ชันนี้สามารถถูกเรียกใช้จากภายนอกได้
func Hi() {
    fmt.Println("Hi 👋")
}
```

4. **การเรียกใช้งาน Package**

ในไฟล์ `main.go` สามารถ import และเรียกใช้ฟังก์ชันจาก package ได้ดังนี้:

```go
// main.go
package main

import "gobasic/greet"

func main() {
    greet.Hi() // แสดงผล: Hi 👋
}
```

  <aside>
💡

**หมายเหตุ:** ชื่อที่ใช้ใน `import` ต้องตรงกับ path ของ package ที่กำหนดไว้ใน go module เช่น `gobasic/greet` หากอยู่ใน module ชื่อ `gobasic`

  </aside>

---

## **2. พื้นฐานภาษา Go**

ทำความเข้าใจการประกาศตัวแปร ฟังก์ชัน และโครงสร้างควบคุมการทำงาน (Control Flow) พื้นฐานในภาษา Go

### หัวข้อที่ครอบคลุม

- การประกาศตัวแปร (Variable Declaration: `var`, `const`, `short declaration`)
- การทำงานของฟังก์ชัน (Function: definition, return values, named return)
- โครงสร้างควบคุมการทำงาน (Control Flow: `if`, `switch`, `for loop`)

---

### **🧱 การประกาศตัวแปร (Variable Declaration)**

ในภาษา Go สามารถประกาศตัวแปรได้หลายรูปแบบ ดังนี้:

1. **การประกาศแบบกำหนดชนิดข้อมูล (Explicit Type)**

ใช้ `var` ตามด้วยชื่อตัวแปรและชนิดของข้อมูล (Type)

```go
// ถ้าไม่กำหนดค่า จะได้ค่าเริ่มต้น (Zero value)
var i int     // 0
var s string  // ""
var ok bool   // false

// กำหนดค่าตอนประกาศ
var i int = 20
var s string = "hello"
var ok bool = true
```

2. **การประกาศแบบไม่กำหนดชนิดข้อมูล (Type Inference)**

เมื่อกำหนดค่าให้ตัวแปร Go จะอนุมานชนิดข้อมูลให้อัตโนมัติ

```go
// ถ้ากำหนดค่าแล้วละ type ได้
var i = 20
var s = "hello"
var ok = true
```

3. **การประกาศแบบ Short Declaration (`:=`)**

เป็นรูปแบบย่อสำหรับการประกาศตัวแปร ใช้ในกรณีที่อยู่ภายในฟังก์ชันเท่านั้น

```go
// เอา var ออก แล้วใช้ := แทน =
i := 20
s := "hello"
ok := true
```

> 💡 [Scope ของตัวแปรใน Go](/blog/go/variable-scope)
>
> - **Package Scope**: ตัวแปรที่ประกาศนอกฟังก์ชัน ใช้งานได้ใน package เดียวกัน
> - **Function Scope**: ตัวแปรที่ประกาศในฟังก์ชัน ใช้ได้เฉพาะในฟังก์ชันนั้น
> - **Block Scope**: ตัวแปรที่ถูกประกาศ ภายใน `{}` จะสามารถใช้งานได้เฉพาะในบล็อกนั้นเท่านั้น

---

### **🔒 ค่าคงที่ (Constant)**

การประกาศค่าคงที่ (Constant) ใช้ `const` แทน `var` และต้องกำหนดค่าตั้งแต่ประกาศ

```go
const i int = 20
const s string = "hello"
const ok bool = true
```

**การสร้าง `enum` ด้วย `iota`**

`iota` ใช้สร้างลำดับเลขอัตโนมัติ เริ่มจาก `0` และเพิ่มขึ้นทีละ `1`

```go
const (
    sunday = iota // 0
    monday        // 1
    tuesday       // 2
    wednesday     // 3
    thursday      // 4
    friday        // 5
    saturday      // 6
)
```

<aside>
💡

หากไม่ต้องการใช้ค่าบางตัว ให้ใช้ `_` แทน แต่ไม่สามารถข้ามลำดับได้

</aside>

```go
const (
    a = iota // 0
    _        // ข้าม 1
    b        // 2
)
```

---

### **⚙️ การทำงานของฟังก์ชัน (Function)**

ฟังก์ชันใน Go ใช้ keyword `func` ตามด้วยชื่อฟังก์ชัน พารามิเตอร์ (ถ้ามี) และค่าที่คืนกลับ (ถ้ามี)

1. ฟังก์ชันพื้นฐาน (Basic Function)

```go
func greeting() {
    fmt.Println("Hello")
}
```

2. ฟังก์ชันที่รับพารามิเตอร์ (Function with Parameters)

```go
func greeting(name string) {
    fmt.Println("Hello", name)
}
```

3. ฟังก์ชันที่คืนค่า (Function with Return Value)

```go
// สามารถประกาศชนิดของค่าที่คืนกลับได้หลังวงเล็บพารามิเตอร์
func sum(a, b int) int {
    return a + b
}
```

4. ฟังก์ชันที่คืนค่าหลายตัว (Multiple Return Values)

```go
func swap(a, b int) (int, int) {
    return b, a
}
```

5. การใช้ Named Return

สามารถตั้งชื่อค่าที่จะคืนกลับได้ โดยกำหนดไว้หลังวงเล็บพารามิเตอร์

```go
func swap(a, b int) (x int, y int) {
    x = b
    y = a
    return
}
```

> 💡 การใช้ Named Return ช่วยให้โค้ดอ่านง่ายขึ้น โดยเฉพาะเมื่อมีการคืนค่าหลายตัว

---

### **🕹️ Control Flow ในภาษา Go**

การควบคุมการทำงาน (Control Flow) ในภาษา Go มีคำสั่งหลักดังนี้:

1. `if-else`
2. `switch`

▶️ `if-else`

รูปแบบการใช้งาน `if-else` ใน Go คล้ายกับภาษาอื่น แต่ไม่ต้องใส่วงเล็บ `()` รอบเงื่อนไข:

- **รูปแบบทั่วไป**

  ```go
  func pow(x, n, lim float64) float64 {
      v := math.Pow(x, n)

      if v < lim {
          return v
      } else {
          fmt.Printf("%g >= %g\n", v, lim)
      }
      return lim
  }
  ```

- **แบบใช้ร่วมกับ Short Statement** (เหมาะสำหรับการประกาศตัวแปรที่ใช้เฉพาะในเงื่อนไข)

  ```go
  func pow(x, n, lim float64) float64 {
      if v := math.Pow(x, n); v < lim {
          return v
      }
      fmt.Printf("%g >= %g\n", v, lim)
      return lim
  }
  ```

  <aside>💡 ตัวแปรที่ประกาศใน Short Statement จะใช้ได้เฉพาะในบล็อก `if-else` เท่านั้น</aside>

▶️ `switch`

ใน Go `switch` มีความแตกต่างจากภาษาอื่น ดังนี้:

- ไม่ต้องใส่วงเล็บ `()`
- ไม่ต้องใช้ `break` เพราะ Go จะ `auto break` ให้อัตโนมัติ
- ใช้ `fallthrough` หากต้องการให้ทำงานต่อใน case ถัดไป

**ตัวอย่างการใช้งาน `switch`**

- **รูปแบบทั่วไป**

  ```go
  os := runtime.GOOS

  switch os {
  case "darwin":
      fmt.Println("macOS")
  case "linux":
      fmt.Println("Linux")
  default:
      fmt.Printf("Unknown OS: %s\n", os)
  }
  ```

- **แบบใช้ร่วมกับ Short Statement**

  ```go
  switch os := runtime.GOOS; os {
  case "darwin":
      fmt.Println("macOS")
  case "linux":
      fmt.Println("Linux")
  default:
      fmt.Printf("Unknown OS: %s\n", os)
  }
  ```

- **แบบไม่มีเงื่อนไข (Switch True)**

  ```go
  os := runtime.GOOS

  switch {
  case os == "darwin":
      fmt.Println("macOS")
  case os == "linux":
      fmt.Println("Linux")
  default:
      fmt.Printf("Unknown OS: %s\n", os)
  }
  ```

  <aside>💡 ใช้ `switch true` เพื่อเขียนเงื่อนไขที่ซับซ้อนหลายแบบ</aside>

---

### **🔁 Loop ในภาษา Go**

Go มีเพียง `for` เป็นคำสั่งวนลูป โดยสามารถเขียนได้หลายรูปแบบ:

- **รูปแบบมาตรฐาน** (เหมือน `for` ในภาษาอื่น ต่างตรงที่ไม่ต้องใส่ `()`)

  ```go
  sum := 0
  for i := 0; i < 10; i++ {
      sum += i
  }
  fmt.Println(sum)
  ```

- **แบบ `while loop`** (ใส่เฉพาะเงื่อนไข)

  ```go
  sum := 0
  i := 0
  for i < 10 {
      sum += i
      i++
  }
  fmt.Println(sum)
  ```

- **แบบ Infinite Loop** (ลูปไม่มีที่สิ้นสุด)

  ```go
  for {
      fmt.Println("Running...")
  }
  ```

- **แบบ `for range`** (ใช้กับ Array, Slice, Map, Channel)

  ```go
  nums := []int{1, 2, 3}

  for index, value := range nums {
      fmt.Println(index, value)
  }

  // ถ้าไม่ต้องการ index
  for _, value := range nums {
      fmt.Println(value)
  }

  // ถ้าต้องการเฉพาะ index
  for index := range nums {
      fmt.Println(index)
  }
  ```

---

### ⤵️ Defer

คำสั่ง `defer` ใช้เลื่อนการทำงานของฟังก์ชันออกไปจนกว่าฟังก์ชันปัจจุบันจะสิ้นสุดการทำงาน:

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")

    defer fmt.Println("This will run at the end")

    fmt.Println("Doing something")
}

// output
Start
Doing something
This will run at the end
```

⚠️ ข้อควรระวังเกี่ยวกับ `defer`

- ค่าของ Arguments จะถูกประมวลผลทันทีเมื่อ `defer` ถูกเรียก
- ลำดับการทำงานของ `defer` เป็นแบบ Stack (LIFO: Last In, First Out)

ตัวอย่างการทำงานแบบ Stack:

```go
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
counting
done
2
1
0

```

**ประโยชน์ของ `defer`:**

- ใช้สำหรับ Cleanup เช่น ปิดไฟล์หรือ Connection
- จัดการ Resource Cleanup ให้อยู่ในที่เดียว อ่านง่ายขึ้น

ตัวอย่างการปิดไฟล์ด้วย `defer`:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    fmt.Println("File opened successfully")
}
```

ในตัวอย่างนี้ `file.Close()` จะถูกเรียกใช้เมื่อฟังก์ชัน `main()` สิ้นสุดลง

---

## **3. ชนิดของข้อมูลใน Go**

Go เป็นภาษาที่มีระบบ type ที่ชัดเจนและเข้มงวด (strongly typed) ซึ่งช่วยให้เราทำงานกับข้อมูลได้อย่างมีประสิทธิภาพ

### หัวข้อที่ครอบคลุม

- ชนิดข้อมูลพื้นฐาน (Basic Types)
- การแปลงประเภท (Type Conversions)
- String (การทำงานเชิงลึกเกี่ยวกับข้อความ)
- Pointer (ตัวชี้หน่วยความจำ)
- Composite Types

---

### **🔢** ชนิดข้อมูลพื้นฐาน (Basic Types)

Go มีชนิดข้อมูลพื้นฐานที่ใช้งานบ่อยดังนี้:

| ชนิดข้อมูล | คำอธิบาย                 | Zero-Value | ตัวอย่าง                 |
| ---------- | ------------------------ | ---------- | ------------------------ |
| `int`      | จำนวนเต็ม                | `0`        | `int age = 30`           |
| `float64`  | จำนวนจริง (ทศนิยม)       | `0`        | `float64 price = 99.99`  |
| `string`   | ข้อความ                  | `""`       | `string name = "GoLang"` |
| `bool`     | ค่าความจริง (true/false) | `false`    | `bool isActive = true`   |

ตัวอย่างการใช้งาน:

```go
package main

import "fmt"

func main() {
    var age int = 30
    var price float64 = 99.99
    var name string = "GoLang"
    var isActive bool = true

    fmt.Println("Age:", age)
    fmt.Println("Price:", price)
    fmt.Println("Name:", name)
    fmt.Println("Is Active:", isActive)
}
```

---

### **🔄** การแปลงประเภท (Type Conversions)

Go ไม่รองรับ implicit type conversion (การแปลงชนิดข้อมูลอัตโนมัติ) แต่เราสามารถทำ explicit conversion ได้โดยการใช้ฟังก์ชันแปลงชนิดข้อมูลโดยตรง

ตัวอย่างการแปลงชนิดข้อมูล:

```go
package main

import "fmt"

func main() {
    var age int = 30
    var price float64 = float64(age) // แปลง int เป็น float64

    fmt.Println("Age as float64:", price)

    var num float64 = 42.5
    var intNum int = int(num) // แปลง float64 เป็น int (ปัดเศษทิ้ง)

    fmt.Println("Float to int:", intNum)
}
```

<aside>
💡

การแปลงชนิดข้อมูลอาจทำให้ข้อมูลสูญหาย เช่น การแปลง `float64` เป็น `int` จะตัดทศนิยมออก

</aside>

---

### **📜** String (การทำงานเชิงลึกเกี่ยวกับข้อความ)

ใน Go ชนิดข้อมูล `string`

- เป็นลำดับของ byte (ค่าที่อ่านได้คือ `uint8` หรือ ASCII)
- เป็น immutable (เปลี่ยนแปลงค่าโดยตรงไม่ได้)
- เก็บข้อมูลแบบ UTF-8 ทำให้รองรับหลายภาษา

```go
str := "Hello, Go!"
fmt.Println(str[0]) // 72 (ASCII ของ 'H')
fmt.Println(string(str[0])) // H
```

**❇️ การนับความยาวตัวอักษรในภาษา Go**

1. `len()` - การนับความยาวแบบ Byte

ฟังก์ชัน `len()` ที่มาพร้อมกับ Go จะคืนค่าจำนวนไบต์ (bytes) ใน `string` โดย `string` ใน Go จะประกอบไปด้วยลำดับของ `byte` และ `len()` จะนับจำนวนของ `byte` ที่ใช้ทั้งหมด

ตัวอย่างเช่น หากสตริงนั้นมีตัวอักษรหลาย `byte` (เช่น อักษรในรูปแบบ UTF-8) `len()` จะนับจำนวน `byte` ที่ใช้ ไม่ใช่จำนวนตัวอักษร (rune) ซึ่ง "`rune`" คือ ตัวอักษร Unicode

**โค้ดตัวอย่าง:**

```go
package main

import (
"fmt"
)

func main() {
str := "Hello, สวัสดี"
fmt.Println("ความยาวในหน่วย byte:", len(str)) // ผลลัพธ์: 17 (เนื่องจากอักษรภาษาไทยใช้หลาย byte ใน UTF-8)
}
```

2. `utf8.RuneCountInString()` - การนับจำนวนตัวอักษร (rune)

เมื่อทำงานกับ `string` ที่มีตัวอักษรหลาย `byte` (เช่น อักษรที่ไม่ใช่ ASCII) อาจจะต้องการนับจำนวนตัวอักษร (rune) แทนการนับจำนวน `byte` ในกรณีนี้ เราสามารถใช้ฟังก์ชัน `utf8.RuneCountInString()` จากแพ็คเกจ `utf8` ซึ่งจะนับจำนวน `rune` (หรือจำนวนตัวอักษร Unicode) ในสตริงนั้น

**โค้ดตัวอย่าง:**

```go
package main

import (
"fmt"
"unicode/utf8"
)

func main() {
str := "Hello, สวัสดี"
fmt.Println("จำนวน rune (ตัวอักษร):", utf8.RuneCountInString(str)) // ผลลัพธ์: 13
}

```

ความแตกต่างนี้เกิดจากการที่ `string` ในภาษา Go ใช้การเข้ารหัสแบบ UTF-8 ซึ่งเป็นการเข้ารหัสที่มีขนาดตัวแปร (variable-width encoding) UTF-8 ใช้ `byte` เดียวสำหรับอักษร ASCII แต่จะใช้หลาย `byte` สำหรับอักษรในภาษาที่ไม่ใช่ ASCII เช่น ภาษาไทย จีน อิโมจิ ฯลฯ

- ใช้ `len()` เมื่อต้องการรู้ความยาวในหน่วย `byte` ของ `string` ซึ่งสำคัญสำหรับการจัดการหน่วยความจำ หรือการทำงานกับการเข้ารหัสในระดับต่ำ
- ใช้ `utf8.RuneCountInString()` เมื่อคุณต้องการนับจำนวนตัวอักษร เช่น ในกรณีที่ต้องการประมวลผลหรือแสดงผล `string` ที่มีตัวอักษรหลาย`byte` โดยเฉพาะเมื่อทำงานกับการแสดงผลในภาษาต่าง ๆ หรือการรองรับหลายภาษา

**❇️ การวนลูปใน String**

ใช้ `for range` และค่า value จะเป็น `rune` ออกมาเลย

**โค้ดตัวอย่าง:**

```go
package main

import "fmt"

func main() {
 message := "Hello, สวัสดี"

 for i, char := range message {
  fmt.Printf("Index: %d, Code: %d, Char: %c\n", i, char, char)
 }
}
```

**ผลลัพท์ที่ได้:**

```go
Index: 0, Code: 72, Char: H
Index: 1, Code: 101, Char: e
Index: 2, Code: 108, Char: l
Index: 3, Code: 108, Char: l
Index: 4, Code: 111, Char: o
Index: 5, Code: 44, Char: ,
Index: 6, Code: 32, Char:
Index: 7, Code: 3626, Char: ส
Index: 10, Code: 3623, Char: ว
Index: 13, Code: 3633, Char: ั
Index: 16, Code: 3626, Char: ส
Index: 19, Code: 3604, Char: ด
Index: 22, Code: 3637, Char: ี
```

❇️ **Package ที่เกี่ยวกับ String**

1. `strings` – ใช้จัดการกับข้อความ

```go
package main

import (
"fmt"
"strings"
)

func main() {
// ตรวจสอบว่ามีคำนี้ในข้อความหรือไม่ case sensitive
result1 := strings.Contains("Hello Gopher", "go")
fmt.Println(result1) // false

// นับคำที่ต้องการหาว่ามีกี่คำในข้อความ
result2 := strings.Count("สวัสดีชาวโก", "ดี")
fmt.Println(result2) // 1

// ตรวจสอบคำขึ้นต้น
result3 := strings.HasPrefix("สวัสดีชาวโก", "สวั")
fmt.Println(result3) // true

// ตรวจสอบคำลงท้าย
result4 := strings.HasSuffix("สวัสดีชาวโก", "โก")
fmt.Println(result4) // true

// ต่อข้อความจาก []string
result5 := strings.Join([]string{"สวัสดี", "ชาวโก"}, "_")
fmt.Println(result5) // สวัสดี_ชาวโก

// แปลงข้อความเป็นตัวพิมพ์ใหญ่ทั้งหมด
result6 := strings.ToUpper("Hello Gopher")
fmt.Println(result6) // HELLO GOPHER

// แปลงข้อความเป็นตัวพิมพ์เล็กทั้งหมด
result7 := strings.ToLower("Hello Gopher")
fmt.Println(result7) // hello gopher
}
```

2. `strconv` – แปลงข้อความเป็นตัวเลข และกลับกัน

```go
package main

import (
"fmt"
"strconv"
)

func main() {
// แปลงเป็น float ต้องระบุขนาดเสมอ (32/64)
f, _ := strconv.ParseFloat("3.1415", 64)
fmt.Println(f) // 3.14

// แปลงเป็น int ต้องระบุเลขฐานของข้อความที่จะแปลงด้วย
i, _ := strconv.ParseInt("-42", 10, 64)
fmt.Println(i) // -42

// แปลงเป็น int ต้องระบุเลขฐานของข้อความที่จะแปลงด้วย
u, _ := strconv.ParseUint("42", 10, 64)
fmt.Println(u) // 42

// แปลงเป็น bool
b, _ := strconv.ParseBool("true")
fmt.Println(b) // true

// แปลงข้อความเป็นจัวเลข
i2, _ := strconv.Atoi("-42")
fmt.Println(i2) // -42

// แปลงตัวเลขเป็นข้อความ
s := strconv.Itoa(-42)
fmt.Println(s) // "-42"
}
```

---

### **📌** Pointer (ตัวชี้หน่วยความจำ)

Pointer คือ ตัวแปรที่เก็บที่อยู่ของหน่วยความจำ (memory address) การใช้ pointer ช่วยให้เราสามารถเปลี่ยนค่าที่อยู่ในหน่วยความจำได้โดยตรง มีค่า zero value คือ `nil`

การใช้งาน Pointer ในช่วงแรกๆ อาจทำให้หลายๆ คนงง แต่มีวิธีจำง่ายๆ คือ

- การประกาศใช้ `*T` (`*`หน้าชนิดของข้อมูล)
- ใช้ operator `&` หน้าชื่อตัวแปร เมื่อต้องการเอาค่า memory address ออกมา
- ใช้ operator `*` หน้าชื่อตัวแปร เพื่อเข้าถึงค่า หรือแก้ไขค่าใน address นั้นๆ

  ```go
  package main

  import (
    "fmt"
  )

  func main() {
    name := "Go"

   var s *string // สร้าง pointer ของ string

   s = &name // ดึงค่า address ออกมา

   fmt.Println(*s) // ผลลัพท์: Go เนื่องจาก *s เข้าถึงค่าใน address ที่อ้างอิงถึง

   *s = "Golang"     // แก้ไขค่าใน address ที่อ้างอิงถึง
   fmt.Println(name) // Golang
  }

  ```

❇️ **Pass by Value vs Pass by Reference**

การส่งค่าผ่านฟังก์ชันใน Go นั้น จะเป็นการส่งแบบ **Pass by Value** เสมอ ดังนั้นเมื่อส่งค่าผ่านฟังก์ชันมันจะทำสำเนาตัวแปรใหม่ออกมา ทำให้ตัวแปรที่ส่งไป และตัวแปรในฟังก์ชันคือคนละตัวกัน

หากต้องการให้ตัวแปรที่ส่งไป สามารถถูกแก้ไขจากในฟังก์ชันได้ ต้องส่งแบบ **Pass by Reference** หรือก็คือส่งเป็น pointer แทน

**โค้ดตัวอย่าง:**

```go
package main

import (
 "fmt"
)

func inc(n int) {
 n = n + 1
}

func incPointer(n *int) {
 *n = *n + 1
}

func main() {
 i := 1

 inc(i)
 fmt.Println(i) // ผลลัพท์: 1

 incPointer(&i)
 fmt.Println(i) // ผลลัพท์: 2
}
```

---

### 📦 Composite Types

**Composite Types** ในภาษา Go คือประเภทข้อมูลที่สามารถประกอบด้วยข้อมูลหลายๆ ค่า (หรือหลายๆ ฟิลด์) เพื่อให้สามารถเก็บข้อมูลที่มีหลายลักษณะในตัวเดียวได้ ซึ่งประเภทข้อมูลที่ถือเป็น composite types ใน Go ได้แก่:

- **Array**
- **Slice**
- **Struct**
- **Map**

### 🗃️ Array

Array คือชุดข้อมูลที่มีขนาดคงที่ (fixed size) และทุกๆ สมาชิกใน array จะต้องเป็นประเภทข้อมูลเดียวกัน และมีค่า zero value ตามประเภทข้อมูลนั้น ๆ

```go
package main

import "fmt"

func main() {
 var nums [4]int = [4]int{1, 2, 3} // กำหนด array ขนาด 4

 fmt.Println("Array:", nums) // ผลลัพท์: Array: [1 2 3 0]
}
```

- การเข้าถึงและแก้ไขค่า

  ```go
  fmt.Println(nums[0]) // ผลลัพท์: 1
  fmt.Println(nums[1]) // ผลลัพท์: 2
  fmt.Println(nums[2]) // ผลลัพท์: 3
  fmt.Println(nums[3]) // ผลลัพท์: 0

  nums[3] = 4                 // แก้ไขค่าในตำแหน่งที่ 3
  fmt.Println("Array:", nums) // ผลลัพท์: Array: [1 2 3 4]
  ```

- การวนลูปด้วย `for range`

  ```go
  for i, v := range nums {
      fmt.Println(i, v) // i คือ index, v คือ value
  }
  ```

### 📏 Slice

Slice คือประเภทข้อมูลที่คล้ายกับ array แต่มีความยืดหยุ่นมากกว่า เพราะสามารถขยายขนาดได้ตามต้องการ ซึ่ง slice จะอ้างอิงไปยัง array ที่อยู่ภายใน มีค่า zero value คือ `nil`

```go
// จะมีค่าเป็น zero value คือ nil ยังใช้งานไม่ได้
var nums []int

// ต้อง allocate memory ให้ก่อนถึงจะใช้งานได้
nums = make([]int, 4) // จะต้องระบุขนาดเริ่มต้นให้ก่อน ทุกตำแหน่งจะได้ค่า zero value ของ type นั้นๆ
// []int{0, 0, 0, 0}

// หรือสร้างแบบว่างๆ
nums := []int{}

// หรือ จะใส่ค่าไปตั้งแต่ประกาศเลยก็ได้
nums := []int{1, 2, 3}
```

- `append()` - เพิ่มค่าลงใน slice

  ```go
  nums := []int{1, 2, 3}

  nums = append(nums, 4)

  fmt.Printf("%#v\n", nums) // []int{1, 2, 3, 4, }
  ```

- `len()` กับ `cap()` - ดูขนาด และดูพื้นที่สำหรับเก็บข้อมูลใน slice ตามลำดับ

  ```go
  y := len(nums)
  z := cap(nums)

  fmt.Printf("%#v, len=%v\n, cap=%v", nums, y, z) // []int{1, 2, 3, 4}, len=4, cap=6
  ```

- `[index_เริ่มต้น:index_ที่ต้องการ+1]` - สำหรับการ slice ข้อมูล

  ```go
  //          0   1   2   3   4   5   6   7   8
  x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

  // เอาตั้งแต่ 0 จนถึงตัวสุดท้าย
  y := x[0:]

  // หรือเขียนแบบนี้ก็ได้
  y := x[:]

  // ถ้าต้องการ 30, 40, 50
  y :=x[2:5] // ค่า 30 คือ index ที่่ 2 ส่วน 50 คือ index ที่ 4 ดังนั้นต้อง + 1 = 5

  // ถ้าต้องการต้องแต่เริ่มต้น จนถึงตำแหน่งที่ต้องการ สามารถละ index เริ่มต้นได้
  y :=x[:5]
  ```

  <aside>
    💡 เปลี่ยน array เป็น slice ง่ายๆ ด้วย `slice := array[:]` แต่‼️เป็นการแชร์ array กัน
    ถ้าแก้ไขที่ตำแหน่งเดียวกันจะเปลี่ยนทั้งคู่ เพราะ slice จะเป็น memory address ของ array
  </aside>

- การลบข้อมูลออกจาก slice
  จะใช้วิธี slice ข้อมูลออกแบบ 2 ก้อน แล้วนำมาต่อกันใหม่ เช่น

  ```go
  words := []string{"A", "B", "C", "D", "E"}

  // ถ้าต้องการจะลบ "C" ออกไป ซึ่งก็คือตำแหน่งที่ 2 -> words[:2]
  // จะต้องได้ {"A", "B"} + {"D", "E"}
  // แต่ไม่สามารถใส่ words[3:] ลงไปตรงๆ ได้
  // จะต้องใช้ spread operator แทนแบบนี้
  words = append(words[:2], words[3:]...)

  fmt.Println(words) // ผลลัพธ์: [A B D E]
  ```

- การวนลูปด้วย `for range`

  ```go
  nums := []int{1, 2, 3}
  for i, v := range nums {
      fmt.Println(i, v) // i คือ index, v คือ value
  }
  ```

### 📚 Map

Map คือประเภทข้อมูลที่เก็บข้อมูลในรูปแบบของ key-value pairs ซึ่งสามารถเข้าถึงข้อมูลได้อย่างรวดเร็วโดยการใช้ key และมีค่า zero value คือ `nil`

```go
// ถ้าประกาศขึ้นมาลอยๆ จะมีค่าเป็น zero value คือ nil
var m map[string]string

m := make(map[string]string)
// หรือ
m := map[string]string{}

// หรือ จะใส่ค่าไปตั้งแต่ประกาศเลยก็ได้
m := map[string]string{
  "a": "apple",
  "b": "banana", // ต้องปิดด้วย , เสมอ
}

// เข้าถึงค่าโดยใช้ key
fmt.Println(m["a"]) // apple
```

- การเข้าถึงค่าโดยใช้ key ถ้าที่ไม่มีอยู่จริงจะได้ zero value กลับมา ทำให้เกิดปัญหาว่า key นั้นมีจริง และมีค่าตามนั้น หรือ key นั่นไม่มี ดังนั้นจะต้องตรวจสอบก่อน

  ```go
  fmt.Printf("%v\n", m["c"])
  // ""

  v, ok := m["c"]
  if ok {
    // แสดงค่าของ key "c"
  }
  ```

- การแก้ไขข้อมูล

  ```go
  m["b"] = "berry"

  fmt.Printf("%#v\n", m)

  // map[string]string{"a":"apple", "b":"berry"}
  ```

- การเพิ่มข้อมูล

  ```go
  m["c"] = "cranberry"

  fmt.Printf("%#v\n")

  // map[string]string{"a":"apple", "b":"banana", "c":"cranberry"}
  ```

- การลบข้อมูล

  ```go
  delete(m, "b")

  fmt.Printf("%#v\n", m)

  // map[string]string{"a":"apple", "c":"cranberry"}
  ```

- การวนลูปด้วย `for range`

  ```go
  m := map[string]string{
    "a": "apple",
    "b": "banana",
  }

  for k, v := range m {
   fmt.Println(k, v) // จะได้ key แทน index
  }
  ```

### 🏗️ Struct

Struct คือประเภทข้อมูลที่สามารถเก็บข้อมูลที่มีหลายประเภท (heterogeneous data) ซึ่งประกอบด้วยฟิลด์ต่างๆ ที่แต่ละฟิลด์สามารถเป็นชนิดข้อมูลที่แตกต่างกันได้

```go
package main

import "fmt"

type Rectangle struct {
 Width  float64
 Height float64
}

func main() {
 rec := Rectangle{} // จะมีค่าเป็น empty -> {}

 // หรือจะกำหนดค่าไปเลยก็ได้

 rec = Rectangle{
  Width:  10,
  Height: 20,
 }

 // การอ่านค่า
 fmt.Println(rec.Width) // 10

 // การกำหนด/แก้ไขค่า
 rec.Width = 30
 fmt.Println(rec.Width) // 30
}
```

❇️ Method

ใน Go นั้น method ของ `struct` เป็นฟังก์ชันที่ถูกผูกไว้กับ `struct` หรือประเภทข้อมูลที่กำหนด โดยที่ `method` จะมีรูปแบบที่แตกต่างจากฟังก์ชันทั่วไป ซึ่งจะมีการกำหนด **receiver** (ตัวรับค่า) ที่เป็นตัวแทนของ `struct` ที่ฟังก์ชันนั้นจะทำงานด้วย

1. **ฟังก์ชันธรรมดา (Normal Function)**

ก่อนที่เราจะเข้าใจ method ของ `struct` ให้เรามาดูฟังก์ชันธรรมดากันก่อน ฟังก์ชันธรรมดาคือฟังก์ชันที่ไม่มี receiver ซึ่งหมายความว่า ฟังก์ชันเหล่านี้ไม่ได้ผูกกับตัวแปรใดๆ โดยตรง

ตัวอย่าง:

```go
package main

import "fmt"

type Rectangle struct {
Width  float64
Height float64
}

// normal fucntion
func Area(rec Rectangle) float64 {
return rec.Width * rec.Height
}

func main() {
rec := Rectangle{
  Width:  10,
  Height: 20,
}

// fucntion style
fmt.Println(Area(rec))
}
```

2. **Method ของ `struct` (Methods with Receiver)**

   เมื่อเราต้องการสร้าง method สำหรับ `struct` ใน Go, เราจะกำหนด **receiver** ในการประกาศ method นั้นๆ ซึ่ง receiver จะเป็นตัวแปรที่ใช้ในการอ้างอิงข้อมูลของ `struct` ที่ method จะทำงานด้วย โดย receiver จะเป็นส่วนสำคัญในการเชื่อมโยงระหว่าง method และ `struct`

การประกาศ method จะมีรูปแบบดังนี้:

```go
func (receiverName Type) methodName(parameters) returnType {
    // ฟังก์ชันที่ใช้ receiver
}
```

ตัวอย่าง:

```go
package main

import "fmt"

type Rectangle struct {
Width  float64
Height float64
}

// recevier function
func (rec Rectangle) Area() float64 {
return rec.Width * rec.Height
}

func main() {
rec := Rectangle{
  Width:  10,
  Height: 20,
}

// method style
fmt.Println(rec.Area())
}
```

3. **Receiver Type: Value vs Pointer Receiver**

method receiver สามารถใช้ได้ทั้งในรูปแบบของ **value receiver** หรือ **pointer receiver** ซึ่งการเลือกใช้ขึ้นอยู่กับความต้องการของการจัดการกับข้อมูลใน `struct`

**▶️ Value Receiver:**

เมื่อใช้ value receiver, Go จะทำการคัดลอกค่าของ `struct` นั้นๆ มาใช้ใน method ซึ่งหมายความว่า method จะไม่สามารถแก้ไขค่าภายใน `struct` ได้

ตัวอย่าง:

```go
package main

import "fmt"

type Rectangle struct {
Width  float64
Height float64
}

// Value receiver
func (r Rectangle) setWidth(w float64) {
    r.width = w // การเปลี่ยนแปลงนี้จะไม่ส่งผลกับตัวแปร rect ที่ถูกเรียก
}

func main() {
    rect := Rectangle{width: 10, height: 20}
    rect.setWidth(20)
    fmt.Println("Width after set:", rect.width) // ผลลัพธ์: 10
}
```

**▶️ Pointer Receiver:**

เมื่อใช้ pointer receiver, method จะรับพอยน์เตอร์ของ `struct` ซึ่งทำให้ method สามารถแก้ไขข้อมูลภายใน `struct` ได้

ตัวอย่าง:

```go
package main

import "fmt"

type Rectangle struct {
Width  float64
Height float64
}

// Pointer receiver
func (r *Rectangle) setWidth(w float64) {
    r.width = w // การเปลี่ยนแปลงนี้จะส่งผลกับตัวแปร rect ที่ถูกเรียก
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    rect.setWidth(20)
    fmt.Println("Width after set:", rect.width) // ผลลัพธ์: 20
}
```

การเลือกใช้ Value Receiver หรือ Pointer Receiver

- **Value Receiver:** ใช้เมื่อคุณไม่ต้องการให้ method แก้ไขข้อมูลของ `struct` หรือเมื่อ `struct` เป็นค่าที่เล็ก (เช่น `int`, `string`, หรือ `float`)
- **Pointer Receiver:** ใช้เมื่อคุณต้องการให้ method แก้ไขข้อมูลใน `struct` หรือเมื่อ `struct` เป็นค่าที่ใหญ่

❇️ **Type Embedding**

ใน Go, **Type Embedding** (หรือที่เรียกว่า Composition) เป็นแนวทางการออกแบบที่ทำให้ `struct` สามารถฝัง (embed) `struct` อื่นได้ ซึ่งทำให้ `struct` ใหม่สามารถใช้ฟังก์ชันและพฤติกรรมจาก `struct` ที่ถูกฝังได้โดยไม่ต้องประกาศฟังก์ชันใหม่

การใช้ Type Embedding ช่วยให้ Go มีลักษณะคล้ายกับการสืบทอด (inheritance) ในภาษาอื่นๆ อย่างเช่น C++ หรือ Java แต่ใน Go, การใช้ embedding จะเป็นการ **composition** ซึ่งช่วยให้สามารถสร้างโค้ดที่มีความยืดหยุ่นและง่ายต่อการบำรุงรักษา

```go
package main

import "fmt"

type Shape struct {
 Name string
}

type Rectangle struct {
 Width  float64
 Height float64
 Shape  // Embedding Shape
}

func main() {
 rec := Rectangle{
  Width:  10,
  Height: 20,
  Shape:  Shape{Name: "Rectangle"},
 }

 fmt.Println(rec.Shape.Name) // ผลลัพธ์: Rectangle
 // คุณสมบัติการ promoted fields
 fmt.Println(rec.Name)       // ผลลัพธ์: Rectangle
}
```

<aside>
💡

หากมี method ชื่อเดียวกันใน `struct` method จาก `struct` ที่ประกาศล่าสุดจะถูกใช้

</aside>

---

## 4. ออกแบบโค้ดให้ยืดหยุ่นด้วย **Interface**

การใช้งาน **Interface** ใน Go เป็นเครื่องมือที่ช่วยให้โค้ดของเรายืดหยุ่นและสามารถทำงานได้กับหลายประเภทข้อมูล **interface** เป็นชนิดข้อมูลที่สามารถเก็บค่าได้หลากหลายประเภทโดยไม่ต้องระบุชนิดข้อมูลเฉพาะล่วงหน้า

### หัวข้อที่ครอบคลุม

- **Empty Interface (`interface{}`)**
- **Interface ที่ไม่เป็น Empty (Interface with Methods)**
- **Type Assertion (การยืนยันประเภทข้อมูล)**
- **Nil Interface (กรณีที่ interface มีค่าเป็น nil)**

---

### **❇️ Empty Interface (`interface{}`)**

**Empty interface** หรือ `interface{}` คือ interface ที่ไม่กำหนด method ใดๆ เลย หมายความว่า `interface{}` สามารถรับค่าเป็นได้ทุกประเภท ซึ่งทำให้มันเป็นเครื่องมือที่มีประโยชน์เมื่อเราต้องการรับค่าใดๆ ที่ไม่รู้ประเภทล่วงหน้า

**ตัวอย่าง:**

```go
package main

import "fmt"

// ฟังก์ชันที่ใช้รับค่าเป็น Empty Interface (interface{})
func printAnything(value interface{}) {
    fmt.Println(value)
}

func main() {
    printAnything("Hello, World!")   // รับค่า string
    printAnything(123)               // รับค่า int
    printAnything(true)              // รับค่า bool
}
```

**ผลลัพธ์:**

```
Hello, World!
123
true
```

---

### **❇️ Interface ที่ไม่เป็น Empty (Interface with Methods)**

นอกเหนือจาก Empty interface, Go ยังอนุญาตให้เรากำหนด **interface** ที่มี method เพื่อให้ชนิดข้อมูลที่ implement interface นั้นๆ ต้องมี method ที่เรากำหนดไว้ใน interface เช่น เราสามารถสร้าง `interface` ที่บังคับให้ต้องมีฟังก์ชัน `Speak` ดังนี้:

**ตัวอย่าง:**

```go
package main

import "fmt"

// สร้าง interface ที่กำหนด method Speak
type Speaker interface {
    Speak() string
}

// สร้าง struct ที่ implement interface Speaker
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

// ฟังก์ชันที่รับ argument เป็น interface Speaker
func introduce(speaker Speaker) {
    fmt.Println(speaker.Speak())
}

func main() {
    person := Person{Name: "Alice"}
    introduce(person) // จะใช้ method Speak ที่อยู่ใน struct Person
}
```

**ผลลัพธ์:**

```
Hello, my name is Alice
```

ในตัวอย่างนี้:

- `Speaker` เป็น interface ที่มี method `Speak()` ซึ่งทุกๆ ชนิดที่ implement interface นี้จะต้องมี method `Speak()` ที่ทำงานบางอย่าง
- `Person` struct มี method `Speak()` ที่ทำให้ `Person` เป็นไปตามข้อกำหนดของ `Speaker` interface
- ฟังก์ชัน `introduce` รับ parameter ที่เป็น `Speaker` ซึ่งสามารถรับค่าทุกประเภทที่มี method `Speak`

---

### **❇️ Type Assertion**

**Type assertion** เป็นการตรวจสอบและแปลงประเภทของค่าใน interface ว่าเป็นชนิดใด เราสามารถทำได้โดยใช้รูปแบบ `value, ok := iface.(Type)` ซึ่งจะพยายามแปลงค่าใน `iface` ไปเป็นประเภท `Type` ถ้าไม่สามารถแปลงได้จะได้รับค่าผลลัพธ์เป็น `false` ในตัวแปร `ok`

**ตัวอย่าง:**

```go
package main

import "fmt"

func checkType(value interface{}) {
    if str, ok := value.(string); ok {
        fmt.Println("This is a string:", str)
    } else if num, ok := value.(int); ok {
        fmt.Println("This is an integer:", num)
    } else {
        fmt.Println("Unknown type")
    }
}

func main() {
    checkType("Hello, Go!")  // String
    checkType(42)            // Integer
    checkType(3.14)          // Unknown type
}
```

**ผลลัพธ์:**

```
This is a string: Hello, Go!
This is an integer: 42
Unknown type
```

ในตัวอย่างนี้:

- เราทำการ **type assertion** เพื่อเช็คประเภทของค่าที่รับเข้ามาใน `checkType`
- เราใช้ `value.(string)` เพื่อดูว่าค่าที่รับมาเป็นประเภท `string` หรือไม่
- ถ้าไม่ใช่ `string` เราก็เช็คต่อว่าเป็น `int` หรือไม่
- ถ้าไม่ตรงกับประเภทใดๆ ที่เราตรวจสอบ จะเข้าสู่สภาวะ "Unknown type"

---

### **❇️ Nil Interface (กรณีที่ interface มีค่าเป็น nil)**

เมื่อ `interface{}` ไม่มีค่า (เช่น `nil`), การใช้ type assertion กับ `nil` ก็สามารถเกิดขึ้นได้

**ตัวอย่าง:**

```go
package main

import "fmt"

func checkNil(value interface{}) {
    if value == nil {
        fmt.Println("The interface is nil")
    } else {
        fmt.Println("The interface is not nil")
    }
}

func main() {
    var i interface{}
    checkNil(i)  // ค่าเป็น nil

    i = 10
    checkNil(i)  // ค่าไม่เป็น nil
}
```

**ผลลัพธ์:**

```
The interface is nil
The interface is not nil
```

---

## 5. ทำความรู้จักกับ Generic ในภาษา Go

Generics เป็นความสามารถที่เพิ่มเข้ามาใน Go ตั้งแต่เวอร์ชัน 1.18 ช่วยให้สามารถสร้างฟังก์ชันและโครงสร้างข้อมูลที่รองรับชนิดข้อมูลหลายแบบ (type-agnostic) โดยไม่ต้องทำซ้ำโค้ด

### หัวข้อที่ครอบคลุม

- ก่อนหน้าที่จะมี Generics
- Generic Function
- Generic Type
- การใช้ Constraint เพื่อจำกัดประเภท
- ข้อดีของการใช้ Generics

---

### ❇️ ก่อนหน้าที่จะมี Generics

ก่อนหน้าที่จะมี Generics การสร้างฟังก์ชันที่รองรับหลายชนิดข้อมูลต้องใช้ `interface{}` ซึ่งขาดความปลอดภัยในการตรวจสอบประเภท (type-safety) ในขณะคอมไพล์

ตัวอย่างการใช้ `interface{}` แบบเดิม:

```go
package main

import "fmt"

func PrintValue(value interface{}) {
    fmt.Println(value)
}

func main() {
    PrintValue(42)       // int
    PrintValue("Hello")  // string
}
```

ปัญหาของการใช้ `interface{}`

1. **ไม่มี Type-safety**: ต้องแปลงประเภทเองเมื่อใช้งาน
2. **ประสิทธิภาพต่ำกว่า**: มี Overhead จากการแปลงประเภท

### ❇️ Generic Function

สามารถสร้างฟังก์ชันที่รองรับหลายประเภทข้อมูลด้วย Type Parameter โดยใช้เครื่องหมาย `[]` กำหนดชื่อประเภท

ตัวอย่าง: ฟังก์ชัน `Sum` ที่รวมค่าของ slice:

```go
package main

import "fmt"

// T คือ Type Parameter ที่รองรับ int และ float64
func Sum[T int | float64](nums []T) T {
    var total T
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(Sum([]int{1, 2, 3}))       // 6
    fmt.Println(Sum([]float64{1.5, 2.5})) // 4.0
}
```

**คำอธิบาย**:

- `T` เป็นตัวแทนของประเภท (type parameter)
- `T int | float64` กำหนดข้อจำกัด (constraint) ว่า `T` ต้องเป็น `int` หรือ `float64`

### ❇️ Generic Type

สามารถสร้างโครงสร้างข้อมูล (struct) ที่รองรับหลายประเภทได้

ตัวอย่าง: กล่องเก็บข้อมูลแบบ Generic:

```go
package main

import "fmt"

type Box[T any] struct {
    Value T
}

func main() {
    intBox := Box[int]{Value: 100}
    strBox := Box[string]{Value: "Hello"}

    fmt.Println(intBox.Value) // 100
    fmt.Println(strBox.Value) // Hello
}
```

**คำอธิบาย**:

- `Box[T any]` คือโครงสร้างที่รองรับทุกประเภท (`any` คือ type constraint ที่รองรับทุกชนิดข้อมูล)

### ❇️ การใช้ Constraint เพื่อจำกัดประเภท

สามารถกำหนดข้อจำกัดของประเภทด้วย `interface`

ตัวอย่าง: ฟังก์ชัน `Max` ที่คืนค่ามากที่สุด:

```go
package main

import "fmt"

type Ordered interface {
    int | float64 | string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Max(5, 10))         // 10
    fmt.Println(Max(3.5, 2.8))     // 3.5
    fmt.Println(Max("Go", "Golang")) // Golang
}
```

**คำอธิบาย**:

- `Ordered` เป็น `interface` ที่กำหนดข้อจำกัดให้รองรับ `int`, `float64`, `string`
- ฟังก์ชัน `Max` ใช้ `T Ordered` เพื่อจำกัดเฉพาะประเภทที่เปรียบเทียบได้

### ❇️  ข้อดีของการใช้ Generics เพื่อเพิ่มความยืดหยุ่น

1. **ลดการทำซ้ำโค้ด (Code Reusability)**: สามารถสร้างฟังก์ชันและโครงสร้างข้อมูลที่ทำงานได้กับหลายประเภท
2. **เพิ่มความปลอดภัย (Type Safety)**: ประเภทข้อมูลถูกตรวจสอบในขณะคอมไพล์
3. **ประสิทธิภาพที่ดีขึ้น**: ลดการใช้ `interface{}` ที่มี Overhead สูง
4. **รองรับหลายประเภท (Type Agnostic)**: สามารถประยุกต์ใช้กับโครงสร้างข้อมูลทั่วไป เช่น Stack, Queue, หรือ Map

---

## 6. การทำงานแบบ Concurrency

**Concurrency** ใน Go ช่วยให้สามารถประมวลผลหลายๆ อย่างพร้อมกันได้โดยไม่ต้องรอให้กระบวนการหนึ่งเสร็จสิ้นก่อนถึงจะเริ่มกระบวนการถัดไป โดยมีเครื่องมือสำคัญ ได้แก่:

- **Goroutine**: ใช้สำหรับรันฟังก์ชันแบบขนาน
- **Channel**: ใช้สำหรับสื่อสารระหว่าง Goroutine
- **WaitGroup**: ใช้รอให้ Goroutine ทั้งหมดเสร็จสิ้น
- **Mutex**: ใช้ป้องกันการเข้าถึงทรัพยากรพร้อมกัน (ป้องกัน **Race Condition**)
- **Select Statement**: ใช้จัดการหลาย Channel พร้อมกัน

---

### ❇️ **1. Goroutine**

**Goroutine** คือหน่วยประมวลผลที่ Go ใช้สำหรับทำงานแบบขนาน มีน้ำหนักเบากว่า Thread และจัดการโดย Go runtime

**ตัวอย่าง:** การสร้าง Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 5; i++ {
        fmt.Println(msg, i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go printMessage("Hello from Goroutine") // สร้าง Goroutine
    printMessage("Hello from Main")
}
```

**ผลลัพธ์ (อาจเปลี่ยนแปลงได้ในแต่ละครั้ง):**

```
Hello from Main 0
Hello from Goroutine 0
Hello from Main 1
Hello from Goroutine 1
...
```

> 📝 หมายเหตุ:
>
> ถ้าไม่มี `time.Sleep` หรือ `WaitGroup` Goroutine อาจไม่ทันทำงาน เพราะ `main()` จบก่อน

---

### ❇️ **2. Channel**

**Channel** ใช้สำหรับส่งข้อมูลระหว่าง Goroutine

**ตัวอย่าง:** การสื่อสารระหว่าง Goroutine ด้วย Channel

```go
package main

import "fmt"

func sendData(ch chan string) {
    ch <- "Hello from Goroutine" // ส่งข้อมูลผ่าน Channel
}

func main() {
    ch := make(chan string) // สร้าง Channel
    go sendData(ch)         // เรียก Goroutine
    msg := <-ch             // รอรับข้อมูลจาก Channel
    fmt.Println(msg)
}
```

**ผลลัพธ์:**

```
Hello from Goroutine
```

> 📝 หมายเหตุ:

- Channel จะบล็อก (รอ) จนกว่ามีการรับ-ส่งข้อมูล
- `make(chan string)` ใช้สร้าง Channel ที่รับข้อมูลประเภท `string`

---

### ❇️ **3. Buffered Channel**

Buffered Channel สามารถเก็บข้อมูลได้หลายค่าพร้อมกัน โดยไม่บล็อกทันที

**ตัวอย่าง:** Buffered Channel

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3) // Buffered Channel ที่รองรับ 3 ค่า

    ch <- 1
    ch <- 2
    ch <- 3

    fmt.Println(<-ch) // ดึงค่าออก
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

```

**ผลลัพธ์:**

```
1
2
3

```

> 📝 หมายเหตุ:
>
> ถ้าใส่ข้อมูลเกินความจุ Channel จะเกิด Deadlock

---

### ❇️ **4. WaitGroup**

`WaitGroup` ใช้รอให้ Goroutine ทั้งหมดทำงานเสร็จก่อนที่ `main()` จะสิ้นสุด

**ตัวอย่าง:** ใช้ WaitGroup เพื่อรอ Goroutine

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // บอก WaitGroup ว่า Goroutine เสร็จแล้ว
    fmt.Printf("Worker %d started\n", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1) // เพิ่มจำนวน Goroutine
        go worker(i, &wg)
    }
    wg.Wait() // รอจน Goroutine ทั้งหมดเสร็จ
    fmt.Println("All workers done!")
}
```

**ผลลัพธ์:**

```
Worker 1 started
Worker 2 started
Worker 3 started
All workers done!

```

---

### ❇️ **5. Race Condition**

**Race Condition** เกิดขึ้นเมื่อมีหลาย Goroutine แก้ไขข้อมูลพร้อมกัน ทำให้ข้อมูลผิดพลาด

**ตัวอย่าง:** ปัญหา Race Condition

```go
package main

import (
    "fmt"
    "time"
)

var count = 0

func increment() {
    for i := 0; i < 1000; i++ {
        count++
    }
}

func main() {
    for i := 0; i < 10; i++ {
        go increment()
    }
    time.Sleep(2 * time.Second) // รอให้ Goroutine ทำงานเสร็จ
    fmt.Println("Final count:", count)
}
```

**ผลลัพธ์ (คาดเดาไม่ได้):**

```
Final count: 7931
```

> 📝 หมายเหตุ:
>
> ค่า `count` ควรเป็น `10000` แต่เนื่องจาก Race Condition ทำให้ค่าผิดพลาด

---

### ❇️ **6. แก้ Race Condition ด้วย Mutex**

**Mutex** ใช้ล็อกการเข้าถึงทรัพยากรร่วมกัน

**ตัวอย่าง:** ใช้ Mutex เพื่อป้องกัน Race Condition

```go
package main

import (
    "fmt"
    "sync"
)

var (
    count = 0
    mu    sync.Mutex
)

func increment() {
    for i := 0; i < 1000; i++ {
        mu.Lock()   // ล็อกก่อนเข้าถึงข้อมูล
        count++
        mu.Unlock() // ปลดล็อกเมื่อเสร็จ
    }
}

func main() {
    for i := 0; i < 10; i++ {
        go increment()
    }

    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        wg.Wait()
        fmt.Println("Final count:", count)
    }()
    wg.Done()
}
```

**ผลลัพธ์:**

```
Final count: 10000
```

---

### ❇️ **7. Select Statement**

ใช้ `select` เพื่อรอหลาย Channel พร้อมกัน

**ตัวอย่าง:** ใช้ Select เพื่อรับค่าจากหลาย Channel

```go
package main

import (
    "fmt"
    "time"
)

func sendData(ch chan string, msg string, delay time.Duration) {
    time.Sleep(delay)
    ch <- msg
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go sendData(ch1, "From Channel 1", 2*time.Second)
    go sendData(ch2, "From Channel 2", 1*time.Second)

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

**ผลลัพธ์:**

```
From Channel 2

```

---

## 7. การจัดการข้อผิดพลาด (Error Handling)

การจัดการข้อผิดพลาด (Error Handling)

- Error interface
- Panic & Recover

### **🆘 Error interface**

เนื่องจากภาษา Go นั้นจะไม่มี `try-catch` ทำให้ต้องจัดการ `error` ณ จุดที่เกิด `error` เลย

ซึ่งวิธีการจัดการ Error จะใช้วิธีตรวจว่ามีค่า error หรือไม่ ซึ่งถ้ามีจะต้องมีค่า`!=nil` เพราะ Error นั้นเป็น interface ตัวหนึ่ง ทำให้มี zero value คือ `nil`

```go
n, err := strconv.Atoi("5s")
if err !=nil {
  // อาจจะ log หรือส่งออกไปให้ที่อื่นจัดการต่อ
  return err
}
```

และเราสามารถสร้าง Error ขึ้นมาได้จาก package `errors`

```go
func findIndex(s []int, num int) (int, error) {
 for i, v := range s {
  if v == num {
   return i, nil
  }
 }
 return 0, errors.New("number not found")
}
```

และเนื่องจาก Error เป็น interface ที่หน้าตาแบบนี้

```go
type error interface {
  Error() string
}
```

ดังนั้น struct อะไรก็ได้ที่มี method `Error() string` ก็เป็น Erorr ได้เหมือนกัน

### **💣 Panic**

นอกจาก error แล้วใน Go ยังมี `panic()` ที่เอาไว้จัดการ error

สิ่งที่ `panic()` แตกต่างจาก error คือ เมื่อเกิด panic ขึ้นแล้ว โปรแกรมจะหยุดการทำงานลงทันที

ดังนั้น จะใช้ panic ก็ต่อเมื่อไม่ต้องการให้โปรแกรมทำงานต่อแล้ว เช่น ตอนเริ่มต้นโปรแกรม ถ้าโปรแกรมไม่สามารถต่อกับระบบฐานข้อมูลได้ ก็จะให้หยุดการทำงานไปเลย

```go
func main() {
  fmt.Println(1)
  fmt.Println(2)
  panic("Fail")
  fmt.Println(3)
  fmt.Println(4)
  fmt.Println(5)
}

// Output
1
2
panic: Fail
```

### 🪤 ดักจับ Panic ด้วย Recover

เมื่อมีการเรียกฟังก์ชันซ่อนๆ กัน แล้วเกิด panic ขึ้น ระบบจะดูว่าในฟังก์ชันที่เรียกฟังก์ชันทีเกิด panic นั้น ได้มีการจัดการ panic หรือไม่ ถ้าไม่มีก็จะส่ง panic ต่อขึ้นไปเรื่อยๆ จนถึง main() ถ้าไม่การการจัดการก็จะจบการทำงานของโปรแกรมไป

ซึ่งเราสามารถดักจับ panic ได้ด้วยการใช้ `recover()`

ตัวอย่าง:

```go
func a() {
  b()
}

func b() {
  panic("Panic in b")
}

func main() {
  a()
  fmt.Println("Completed")
}

// Output
panic: Panic in b
```

แบบนี้จะโปรแกรมจะหยุดการทำงานเพราะเกิด panic แต่ถ้าเพิ่ม recover เข้าไปใน `a()` โปรแกรมก็จะทำงานได้ต่อจนจบ ดังนี้:

```go
func a() {
  defer func() {
    if r:= recover(); r != nil {
      fmt.Println(r, "Recover in a")
    }
  }()
  b()
}

func b() {
  panic("Panic in b")
}

func main() {
  a()
  fmt.Println("Completed")
}

// Output
Panic in b Recover in a
Completed
```

---

## 8. Advanced concepts ในภาษา Go

### หัวข้อที่ครอบคลุม

- การใช้งานฟังก์ชันขั้นสูง
- การใช้ Context เพื่อจัดการงานที่ซับซ้อน
- การทำงานกับ JSON (encoding/json package)

---

### 💎 การใช้งานฟังก์ชันขั้นสูง

การใช้งานฟังก์ชันขั้นสูงใน Go มีความสำคัญเพราะช่วยให้โค้ดยืดหยุ่นและมีประสิทธิภาพมากขึ้น โดยสามารถจัดการกับพฤติกรรมต่างๆ ของฟังก์ชันได้อย่างยืดหยุ่น เช่น การใช้ค่าหลายค่าในฟังก์ชันเดียวกัน หรือการส่งฟังก์ชันเป็นอาร์กิวเมนต์ในฟังก์ชันอื่นๆ

1. **Variadic Function**

ฟังก์ชันแบบ Variadic ช่วยให้สามารถรับจำนวนอาร์กิวเมนต์ที่ไม่จำกัดในฟังก์ชันเดียวกันได้ โดยใช้ `...`

ซึ่งเป็นประโยชน์เมื่อเราต้องการทำงานกับชุดข้อมูลที่มีขนาดไม่แน่นอน เช่น การรวมค่าหลายค่าเข้าด้วยกัน

ตัวอย่าง:

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))        // Output: 6
    fmt.Println(sum(10, 20, 30, 40)) // Output: 100
}
```

2. **First Class Function**

ฟังก์ชันใน Go เป็น First-Class Citizen หมายความว่า ฟังก์ชันสามารถถูกส่งผ่านเป็นอาร์กิวเมนต์, คืนค่าจากฟังก์ชัน, หรือเก็บไว้ในตัวแปรได้ ซึ่งทำให้ฟังก์ชันเป็นเครื่องมือที่ทรงพลังในการออกแบบโปรแกรมที่มีความยืดหยุ่น

ตัวอย่าง:

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    fn := add             // เก็บฟังก์ชันในตัวแปร
    fmt.Println(fn(3, 4)) // เรียกใช้ฟังก์ชันผ่านตัวแปร
}
```

3. **Higher Order Function**

Higher-Order Function คือฟังก์ชันที่สามารถรับฟังก์ชันอื่นเป็นอาร์กิวเมนต์ หรือส่งฟังก์ชันอื่นกลับเป็นผลลัพธ์ ฟังก์ชันประเภทนี้ช่วยให้เราสามารถสร้างลอจิกที่ยืดหยุ่นและนำกลับมาใช้ใหม่ได้ง่าย

ตัวอย่าง:

```go
package main

import "fmt"

func apply(fn func(int) int, value int) int {
    return fn(value)
}

func square(x int) int {
    return x * x
}

func main() {
    result := apply(square, 4)
    fmt.Println(result) // Output: 16
}

```

4. **Clossure Function**

Closure คือฟังก์ชันที่สามารถ "จดจำ" สถานะภายนอกที่มันถูกสร้างขึ้น แม้ว่าฟังก์ชันนั้นจะถูกส่งไปที่อื่นหรือเรียกใช้ในเวลาภายหลัง ก็คือการที่ฟังก์ชันสามารถเข้าถึงตัวแปรที่อยู่ในขอบเขตของมันได้แม้จะไม่มีการเข้าถึงจากภายนอก

ตัวอย่าง:

```go
package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    next := counter()
    fmt.Println(next())
    fmt.Println(next())
}
```

---

### **⏱️ การใช้ Context ใน Go**

การใช้ `context` ใน Go เป็นเครื่องมือสำคัญที่ช่วยจัดการเกี่ยวกับ **Timeout**, **Deadline**, **Cancellation**, และการส่งค่าในฟังก์ชันต่างๆ โดยเฉพาะในโปรแกรมที่ต้องจัดการการทำงานแบบหลายเธรด (concurrent) หรือการทำงานที่ต้องรอผลลัพธ์ในเวลาจำกัด เช่น การทำงานกับ HTTP request หรือการประมวลผลที่ต้องทำงานร่วมกันหลายๆ ฟังก์ชัน

1. **Timeout และ Deadline**

- **Timeout** คือการกำหนดระยะเวลาที่ฟังก์ชันหรือกระบวนการสามารถทำงานได้ โดยหลังจากหมดระยะเวลา จะยกเลิกการทำงานโดยอัตโนมัติ
- **Deadline** คือการกำหนดเวลาที่ต้องเสร็จสิ้นการทำงาน (โดยมีเวลาที่กำหนดในอนาคต)

ในการใช้งาน `context` เพื่อจัดการ Timeout หรือ Deadline จะใช้ `context.WithTimeout` หรือ `context.WithDeadline` ซึ่งจะส่งคืน `context` ที่มีการกำหนดเวลาจำกัด (Timeout/Deadline) และสามารถใช้ในฟังก์ชันต่างๆ เช่น HTTP request หรือการทำงานที่ต้องรอผลลัพธ์จาก external service

**ตัวอย่าง:**

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func longRunningTask(ctx context.Context) {
    select {
    case <-time.After(2 * time.Second): // จำลองงานที่ใช้เวลานาน
        fmt.Println("Task completed")
    case <-ctx.Done(): // เมื่อเวลา timeout หรือมีการยกเลิก
        fmt.Println("Task cancelled:", ctx.Err())
    }
}

func main() {
    // ใช้ context.WithTimeout เพื่อกำหนดเวลาสำหรับการทำงาน
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel() // ให้แน่ใจว่า cancel ถูกเรียกหลังจากการใช้งาน

    longRunningTask(ctx) // เรียกใช้งานฟังก์ชันที่ต้องรอผล
}
```

ในตัวอย่างนี้ ฟังก์ชัน `longRunningTask` จะถูกยกเลิกทันทีที่ `context` หมดเวลา (หลังจาก 1 วินาที) ถึงแม้ว่าจะยังทำงานไม่เสร็จ

2. **Cancellation**

การยกเลิกการทำงาน (Cancellation) จะใช้เมื่อคุณต้องการหยุดกระบวนการหรือการทำงานในกรณีที่ไม่จำเป็นต้องทำงานต่อ เช่น เมื่อผู้ใช้ปิดหน้าต่าง หรือเมื่อได้รับคำสั่งให้หยุดกระบวนการ

การยกเลิกสามารถทำได้โดยการใช้ `context.WithCancel` ซึ่งจะส่งคืน `context` ที่สามารถใช้ในการตรวจสอบการยกเลิกและฟังก์ชัน `cancel` ที่จะใช้เพื่อยกเลิกการทำงานนั้น

**ตัวอย่าง:**

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func longRunningTask(ctx context.Context) {
    select {
    case <-time.After(5 * time.Second): // งานที่ใช้เวลานาน
        fmt.Println("Task completed")
    case <-ctx.Done(): // ถ้ามีการยกเลิก
        fmt.Println("Task cancelled:", ctx.Err())
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go longRunningTask(ctx) // เรียกใช้งานงานที่ต้องใช้เวลานาน

    time.Sleep(2 * time.Second) // จำลองการทำงานในบางช่วงเวลา
    cancel() // ยกเลิกการทำงานหลังจาก 2 วินาที
}
```

ในตัวอย่างนี้ งานจะถูกยกเลิกเมื่อ `cancel()` ถูกเรียกใน `main` หลังจากผ่านไป 2 วินาที

3. การส่งค่า (Value Passing) ผ่าน Context

`context` ยังสามารถใช้ในการส่งข้อมูลระหว่างฟังก์ชันต่างๆ ได้ เช่น การส่งค่าที่ต้องการใช้ในหลายๆ ฟังก์ชันหรือหลายๆ งาน โดยใช้ `context.WithValue` ซึ่งจะส่งคืน `context` ที่มีข้อมูลบางอย่างที่เกี่ยวข้องกับกระบวนการนั้น

**ตัวอย่าง:**

```go
package main

import (
    "context"
    "fmt"
)

func process(ctx context.Context) {
    if value := ctx.Value("key"); value != nil {
        fmt.Println("Received value:", value)
    } else {
        fmt.Println("No value found")
    }
}

func main() {
    ctx := context.WithValue(context.Background(), "key", "myValue")
    process(ctx) // ส่งค่า "myValue" ผ่าน context
}
```

ในตัวอย่างนี้ ข้อมูล (`"myValue"`) จะถูกส่งผ่าน `context` และสามารถนำไปใช้ในฟังก์ชันที่ต้องการได้

---

### 🔃 การทำงานกับ JSON

การทำงานกับ **JSON** ใน Go ใช้แพ็กเกจ `encoding/json` ที่มาพร้อมกับ Go ซึ่งมีฟังก์ชันสำหรับการแปลงข้อมูลระหว่าง **Go struct** และ **JSON** format (ทั้งการแปลงจาก struct ไปเป็น JSON และจาก JSON ไปเป็น struct) การทำงานกับ JSON ใน Go สามารถทำได้ง่ายๆ โดยใช้ฟังก์ชัน `json.Marshal` และ `json.Unmarshal`

1. **การแปลงจาก struct เป็น JSON (Marshal)**

การแปลงจาก `struct` ไปเป็น JSON สามารถทำได้โดยใช้ฟังก์ชัน `json.Marshal` ซึ่งจะรับ `struct` เป็นอาร์กิวเมนต์และคืนค่าเป็น JSON ในรูปแบบของ **byte slice** (เช่น `[]byte`)

**ตัวอย่าง:**

```go
package main

import (
    "encoding/json"
    "fmt"
)

// สร้าง struct ที่จะใช้ในการแปลง
type Person struct {
    Name    string `json:"name"`    // ใช้ tag "json" เพื่อกำหนดชื่อฟิลด์ใน JSON
    Age     int    `json:"age"`
    Address string `json:"address,omitempty"` // กรณีที่เป็นค่าว่างจะไม่แสดงใน JSON
}

func main() {
    person := Person{
        Name:    "John Doe",
        Age:     30,
        Address: "123 Main St",
    }

    // แปลง struct ไปเป็น JSON
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshaling:", err)
        return
    }

    // แสดงผล JSON ในรูปแบบ string
    fmt.Println(string(jsonData))
}
```

**ผลลัพธ์:**

```json
{ "name": "John Doe", "age": 30, "address": "123 Main St" }
```

ในตัวอย่างนี้:

- เราสร้าง `struct` ชื่อ `Person` โดยมี `Name`, `Age`, และ `Address` เป็นฟิลด์
- เราใช้ `json.Marshal` เพื่อแปลง `struct` เป็น JSON
- ผลลัพธ์ที่ได้คือ JSON ที่แสดงค่าของฟิลด์ใน `struct` ซึ่ง `omitempty` ใน tag `json` ใช้เพื่อไม่ให้ฟิลด์ที่มีค่าเป็นค่าว่าง (เช่น `""` หรือ `0`) ปรากฏใน JSON

2. **การแปลงจาก JSON เป็น struct (Unmarshal)**

การแปลงจาก JSON ไปเป็น `struct` ใช้ฟังก์ชัน `json.Unmarshal` ซึ่งรับ JSON ในรูปแบบของ **byte slice** (เช่น `[]byte`) และแปลงมันเป็น `struct` ที่กำหนด

**ตัวอย่าง:**

```go
package main

import (
    "encoding/json"
    "fmt"
)

// สร้าง struct ที่ใช้ในการแปลง JSON
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Address string `json:"address,omitempty"`
}

func main() {
    jsonData := `{"name":"John Doe","age":30,"address":"123 Main St"}`

    var person Person

    // แปลง JSON ไปเป็น struct
    err := json.Unmarshal([]byte(jsonData), &person)
    if err != nil {
        fmt.Println("Error unmarshaling:", err)
        return
    }

    // แสดงผล struct
    fmt.Println(person)
}
```

**ผลลัพธ์:**

```
{John Doe 30 123 Main St}
```

ในตัวอย่างนี้:

- เรามี JSON ที่เป็น string ซึ่งเก็บข้อมูลของ `name`, `age`, และ `address`
- เราใช้ `json.Unmarshal` เพื่อแปลง JSON เป็น `struct` และใส่ค่าผลลัพธ์ลงในตัวแปร `person`
- `json.Unmarshal` ต้องการพอยน์เตอร์ไปที่ตัวแปร `person` เพื่อทำการอัพเดทข้อมูลจาก JSON

3. **การจัดการกับ JSON ที่มี Key ไม่ตรงกับฟิลด์ใน struct**

ในบางกรณี JSON ที่ได้รับอาจมีคีย์ที่ไม่ตรงกับฟิลด์ใน `struct` หรือไม่ต้องการให้แปลงบางฟิลด์ใน `struct` ไปเป็น JSON เราสามารถใช้ **tag** ในการควบคุมการแปลงได้

- ใช้ `json:"fieldname"` เพื่อกำหนดชื่อฟิลด์ใน JSON
- ใช้ `json:"-"` เพื่อไม่ให้ฟิลด์นั้นถูกแปลงเป็น JSON หรือไม่ให้รับค่าเมื่อแปลงจาก JSON

**ตัวอย่าง:**

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Address string `json:"-"` // ไม่ให้แปลง Address ไปเป็น JSON
}

func main() {
    person := Person{
        Name:    "John Doe",
        Age:     30,
        Address: "123 Main St",
    }

    jsonData, _ := json.Marshal(person)
    fmt.Println(string(jsonData)) // {"name":"John Doe","age":30}
}
```

**ผลลัพธ์:**

```json
{ "name": "John Doe", "age": 30 }
```

---

## **9. Unit Testing ในภาษา Go**

Unit Testing เป็นส่วนสำคัญของการพัฒนาซอฟต์แวร์ เพื่อให้แน่ใจว่าฟังก์ชันต่าง ๆ ทำงานถูกต้องตามที่คาดหวัง ใน Go เราสามารถเขียน Unit Test ได้โดยใช้แพ็กเกจ `testing` ซึ่งเป็นส่วนหนึ่งของ Go standard library

### หัวข้อที่ครอบคลุม

- การเขียน Unit test
- การรัน Unit test
- การทำ Table-Driven Test

---

### 🔖 การเขียน Unit test

โดยทั่วไป การเขียน Unit Test ใน Go จะมีรูปแบบดังนี้:

- ไฟล์หลัก: `main.go`
- ไฟล์ทดสอบ: `main_test.go` (ชื่อไฟล์ต้องลงท้ายด้วย `_test.go`)

ตัวอย่างการเขียน Unit Test

- สมมติว่าคุณมีฟังก์ชันคำนวณผลรวมในไฟล์ `main.go`:

  ```go
  // main.go
  package main

  func Add(a, b int) int {
      return a + b
  }
  ```

- เราสามารถเขียน Unit Test สำหรับฟังก์ชันนี้ได้ดังนี้:

  ```go
  // main_test.go
  package main

  import "testing"

  func TestAdd(t *testing.T) {
      result := Add(2, 3)
      expected := 5

      if result != expected {
          t.Errorf("Add(2, 3) = %d; ต้องการ %d", result, expected)
      }
  }
  ```

### 🪄 รัน Unit Test

ใช้คำสั่ง:

```bash
go test
```

หากการทดสอบผ่าน จะเห็นข้อความสรุปว่าการทดสอบสำเร็จ หากมีข้อผิดพลาดจะมีข้อความแจ้งเตือน

### ✅ การทำ Table-Driven Test

เป็นแนวทางที่นิยมในการเขียน Unit Test เพื่อทดสอบหลาย ๆ กรณีพร้อมกัน:

```go
func TestAddCases(t *testing.T) {
    cases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Positive numbers", 2, 3, 5},
        {"Negative numbers", -1, -1, -2},
        {"Zero", 0, 0, 0},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            result := Add(c.a, c.b)
            if result != c.expected {
                t.Errorf("Add(%d, %d) = %d; ต้องการ %d", c.a, c.b, result, c.expected)
            }
        })
    }
}
```

---

## สรุป

ภาษา Go เป็นภาษาที่มีประสิทธิภาพสูง ใช้งานง่าย และเหมาะสำหรับการพัฒนาแอปพลิเคชันที่ต้องการความเร็วและรองรับการประมวลผลพร้อมกัน บทความนี้ได้ครอบคลุมเนื้อหาตั้งแต่การเริ่มต้นใช้งาน พื้นฐานของภาษา ไปจนถึงแนวคิดขั้นสูง เช่น Interface, Generics, Concurrency และการจัดการข้อผิดพลาด

เมื่อคุณเข้าใจหลักการพื้นฐานและแนวคิดสำคัญเหล่านี้แล้ว คุณจะสามารถนำความรู้ไปประยุกต์ใช้ในการพัฒนาแอปพลิเคชันจริงได้อย่างมั่นใจ และสามารถขยายขอบเขตความสามารถของคุณไปสู่การสร้างระบบที่มีความซับซ้อนมากยิ่งขึ้น

การเรียนรู้ภาษา Go เป็นการลงทุนที่คุ้มค่า ไม่ว่าคุณจะเป็นนักพัฒนามือใหม่หรือมีประสบการณ์แล้ว การมีความเชี่ยวชาญในภาษา Go จะช่วยให้คุณก้าวทันเทคโนโลยีและสามารถพัฒนาโซลูชันที่มีประสิทธิภาพในโลกของซอฟต์แวร์ยุคปัจจุบัน
