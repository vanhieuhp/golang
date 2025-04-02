package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("EmailNotifier send: ", message)
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
	fmt.Println("SMSNotifier send: ", message)
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotificationService(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t string) Notifier {
	switch t {
	case "email":
		return EmailNotifier{}
	case "sms":
		return SMSNotifier{}
	default:
		return SMSNotifier{}
	}

}

func main() {
	s := Service{
		notifier: CreateNotifier("email"),
	}

	s.SendNotificationService("Hello World")
}
