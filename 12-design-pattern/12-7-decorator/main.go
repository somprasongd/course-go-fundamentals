package main

import "fmt"

// Notifier interface
type Notifier interface {
	Send(message string)
}

// BasicNotifier (Concrete Component)
type BasicNotifier struct{}

func (n *BasicNotifier) Send(message string) {
	// จำลองการทำงาน
	fmt.Println("Sending message")
}

// LoggingDecorator (Decorator)
type LoggingDecorator struct {
	notifier Notifier
}

func NewLoggingDecorator(n Notifier) Notifier {
	return &LoggingDecorator{notifier: n}
}

func (l *LoggingDecorator) Send(message string) {
	fmt.Println("[Log] Message:", message)
	l.notifier.Send(message)
}

// EncryptionDecorator (Decorator)
type EncryptionDecorator struct {
	notifier Notifier
}

func NewEncryptionDecorator(n Notifier) Notifier {
	return &EncryptionDecorator{notifier: n}
}

func (e *EncryptionDecorator) Send(message string) {
	encryptedMessage := "Encrypted(" + message + ")"
	e.notifier.Send(encryptedMessage)
}

func main() {
	notifier := &BasicNotifier{}

	// เพิ่ม Logging
	logNotifier := NewLoggingDecorator(notifier)

	// เพิ่ม Encryption
	secureNotifier := NewEncryptionDecorator(logNotifier)

	secureNotifier.Send("Hello, World!")
}
