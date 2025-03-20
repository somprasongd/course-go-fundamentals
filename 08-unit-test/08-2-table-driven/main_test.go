// main_test.go
package main

import "testing"

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
				t.Errorf("got %q, wanted %q", result, c.expected)
			}
		})
	}
}

// go test -v gobasic
// === RUN   TestAddCases
// === RUN   TestAddCases/Positive_numbers
// === RUN   TestAddCases/Negative_numbers
// === RUN   TestAddCases/Zero
// --- PASS: TestAddCases (0.00s)
//     --- PASS: TestAddCases/Positive_numbers (0.00s)
//     --- PASS: TestAddCases/Negative_numbers (0.00s)
//     --- PASS: TestAddCases/Zero (0.00s)
// PASS
// ok      gobasic 0.436s
