package reader

//Reader struct for Reader information
type Reader struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Booklist  string `json:"book_list" db:"book_list"`
}

//Reader slice from Readers
type Readers []Reader
