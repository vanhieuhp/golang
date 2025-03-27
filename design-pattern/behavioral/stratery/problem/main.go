package main

import "fmt"

type NotificationService struct {
	notificationType string
}

func (s NotificationService) SendNotification(message string) {
	if s.notificationType == "email" {
		fmt.Println("Sending email to " + message)
	} else if s.notificationType == "slack" {
		fmt.Println("Sending slack to " + message)
	}
}

func main() {
	s := NotificationService{"email"}
	s.SendNotification("hello world")
}
