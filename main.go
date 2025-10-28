package main

import (
    "fmt"
	//"github.com/RumiaHa/simple-library.git/domain"
    "github.com/RumiaHa/simple-library.git/library"
    //"github.com/RumiaHa/simple-library.git/notifications"

)

func main() {
   myLibrary := library.New()

   myLibrary.AddBook("Гарри Поттер","Джоанн Роулинг", 1999)
   myLibrary.AddReader("Sofia","Tedeeva")
   myLibrary.IssueBookToReader(1,1)

   fmt.Printf("Книга %s, авторства %s, %d года находится у %s %s", myLibrary.Books[0].Title,myLibrary.Books[0].Author,myLibrary.Books[0].Year, myLibrary.Readers[0].FirstName,myLibrary.Readers[0].LastName,)

}