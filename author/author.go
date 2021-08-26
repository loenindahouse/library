package author

//Author struct for Author information
type Author struct {
	ID             int    `json:"id" db:"id"`
	Firstname      string `json:"first_name" db:"first_name"`
	Lastname       string `json:"last_name" db:"last_name"`
	Username       string `json:"user_name" db:"user_name"`
	Specialization string `json:"specialization" db:"specialization"`
	Booklist       string `json:"book_list" db:"book_list"`
}

//Author slice from authors
type Authors []Author
