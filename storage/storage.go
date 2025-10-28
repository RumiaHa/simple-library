package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/RumiaHa/simple-library.git/domain"
)

type Storable interface {
	Save() error
	Load() error
}

func SaveBookToCSV(filename string, books []domain.Book) error {
	file, err := os.Create(filename)
	if err != nil{
		return  fmt.Errorf("не удалось создать файл %s: %w", filename,err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	headers := []string{"ID", "Название", "Автор"}
	if err := writer.Write(headers); err != nil{
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	for _, book := range books {
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать список книгу с ID %d: &w", book.ID,err)
		}
	}
	return nil
}

func LoadBookFromCSV(filename string) ([]domain.Book, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w",filename,err)
	}
	defer file.Close()

	reader :=csv.NewReader(file)

	if _, err := reader.Read(), err != nil {
		return nil, fmt.Errorf("не удалось прочитать заголовок из файла: %w",err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		continue
	}


}
