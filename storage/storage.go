package main

type Storable interface {
	Save() error 
	Load() error 
}

//Лучший кандидат на реализацию этого интерфейса - Library, потому что эта структура захватывает и Book и Reader