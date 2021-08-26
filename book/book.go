package book

//Book struct for Book information
type Book struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Genre    string `json:"genre" db:"genre"`
	ISBN     string `json:"isbn" db:"isbn"`
	AuthorID string `json:"author_id" db:"author_id"`
}

//Book slice from books
type Books []Book
