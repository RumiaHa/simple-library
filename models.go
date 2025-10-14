package main

import "fmt"

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}


type Book struct {
	ID       int
	Title    string
	Author   string
	Year int
	IsIssued bool

}


//Выводит в консоль информацию о читателе
func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName,r.LastName,r.ID)
}

//Dectivate делает читателя не активным
func (r* Reader) Deactivate(){
	r.IsActive = false
}

func (r Reader) String() string{
	status := ""

	if r.IsActive{
		status = "Активен"
	} else {
		status = "Не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь:%",r.FirstName,r.LastName,r.ID,status)
}

func (b Book) String() string {
	return fmt.Sprintf(`"%s (%s, %d)"`, b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook() (r Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана", b.Title)
		return
	}   
	  if !r.IsActive {
		fmt.Printf("Читатель %s %s не активен и не      может получить книгу.", r.FirstName, r.LastName)
		return
	}

	b.IsIssued = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
	return
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s уже в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}
