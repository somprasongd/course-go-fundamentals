// main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	// Arrange
	input1, input2 := 4, 6
	want := 10

	// Act
	got := Add(input1, input2)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// go test -v gobasic
// ผลลัพธ์:
// === RUN   TestAdd
// --- PASS: TestAdd (0.00s)
// PASS
// ok
