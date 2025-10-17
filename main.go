package main

import (
	"fmt"
)


func main(){
	user1 := Reader{
		ID: 1,
		FirstName: "Agunda",
		LastName: "Kokoit`i",
		IsActive: true,
	}
	
reader2 := Reader{
		ID:        2,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}




	book1 := Book{
		Title: "Voina i mir",
		Author: "Lev Tolstoy",
		Year: 1869,
	}

book1.IssuesBook(&user1)
fmt.Println(book1)
book1.IssuesBook(&reader2)


	user1.Deactivate()
	fmt.Println(user1)
	fmt.Println("---")
	book1.IssuesBook(&user1)

	email1 := EmailNotifier{
		EmailAddress: "nikededa2@gmail.com",
	}
	phone1 := SMSNotifier{
		PhoneNumber: 7988888888,
	}
	notifier := []Notifier{email1,phone1}

	for _, notifitacions := range notifier{
		notifitacions.Notify("\n Ваша книга просрочена!")
	}

}