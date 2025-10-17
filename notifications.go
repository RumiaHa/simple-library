package main

import (
	"fmt"
)

type Notifier interface{
	Notify(message string)
}

type EmailNotifier struct{
	EmailAddress string
}

type SMSNotifier struct{
	PhoneNumber int
}

func (e EmailNotifier) Notify(message string){
	fmt.Printf("Отправляю email на [%s]: '[%s]'.", e.EmailAddress, message)
}
func (e SMSNotifier) Notify(message string){
	fmt.Printf("Отправляю SMS на номер [%d]: '[%s]'.", e.PhoneNumber, message)
}