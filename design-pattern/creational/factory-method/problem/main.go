package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
}

func (EmailNotifier) Send(message string) {
	fmt.Println("EmailNotifier Send: ", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Println("SMSNotifier Send: ", message)
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotificationService(message string) {
	s.notifier.Send(message)
}

func main() {
	s := Service{
		// I don't want my users init a new notifier like this.
		// They should call to something to produce a notifier with its specific type
		// CreateNotifier(type) Notifier
		notifier: SMSNotifier{},
	}

	s.SendNotificationService("Hello World")
}
