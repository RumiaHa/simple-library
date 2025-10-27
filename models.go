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
	Year     int
	IsIssued bool
	ReaderId *int
}

type Library struct {
	Books        []*Book
	Readers      []*Reader
	lastBookID   int
	LastReaderID int
}

func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.LastReaderID++

	newReader := &Reader{
		ID:        lib.LastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)
	return newReader
}

func (lib *Library) AddBook(title, author string, year int) (*Book, error) {

	for _, b :=  range lib.Books{
		if b.Title == title && b.Author == author {
			return nil, fmt.Errorf("книга '%s' автора '%s' уже существует", title, author)
		}
	}


	lib.lastBookID++

	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}



	lib.Books = append(lib.Books, newBook)
	return newBook, nil
}

func (lib *Library) FindBookByID(id int) (*Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	if book.IsIssued {
		return fmt.Errorf("книга %s уже выдана", book.Title)
	}

	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен", reader.FirstName, reader.LastName)
	}

	book.IsIssued = true
	book.ReaderId = &reader.ID

	return nil
}


func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	status := ""

	if r.IsActive {
		status = "Активен"
	} else {
		status = "Не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь:%s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string {
	status := "В библиотеке"
	if b.IsIssued && b.ReaderId != nil {
		status = fmt.Sprintf("На руках у читателя с ID: %d", *b.ReaderId)
	}
	return fmt.Sprintf(`"%s (%s, %d), Статус: %s"`, b.Title, b.Author, b.Year, status)
}

func (r *Reader) AssignBook(b *Book) error {

	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !r.IsActive {
		return fmt.Errorf("читатель %s %s не активен", r.FirstName, r.LastName)
	}
	
	b.IsIssued = true
	b.ReaderId = &r.ID
	return nil
}

func (b *Book) IssuesBook(r *Reader) error {

	if b.IsIssued {
		return fmt.Errorf("книга %s уже кому-то выдана", b.Title)
	}
	if !r.IsActive {
		return fmt.Errorf("читатель %s %s не активен и не может получить книгу", r.FirstName, r.LastName)
	}

	b.IsIssued = true
	b.ReaderId = &r.ID
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}

	b.IsIssued = false
	b.ReaderId = nil
	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}