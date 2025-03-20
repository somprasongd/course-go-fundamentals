package main

import "fmt"

// Notification interface
type Notification interface {
	Send(message string)
}

// EmailNotification implements Notification
type EmailNotification struct{}

func (e *EmailNotification) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// SMSNotification implements Notification
type SMSNotification struct{}

func (s *SMSNotification) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// Factory function
func NewNotification(notificationType string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{}
	case "sms":
		return &SMSNotification{}
	default:
		panic("Invalid notification type")
	}
}

func main() {
	// ใช้งาน Factory
	email := NewNotification("email")
	email.Send("Hello via Email!")

	sms := NewNotification("sms")
	sms.Send("Hello via SMS!")
}
