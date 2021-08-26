package author

import (
	"artemstudy/book"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db     *sqlx.DB
	logger hclog.Logger
}

//Create New Store constructor
func NewStore(db *sqlx.DB, logger hclog.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

//GetAll allows you to get all Authors
func (q *Store) GetAll() (Authors, error) {
	logger := q.logger.With("operation", "GetAuthors")
	var authors Authors
	logger.Debug("Start Query")
	rows, err := q.db.Queryx("SELECT * FROM authors")
	if err != nil {
		logger.Error("Failed to get Authors", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var author Author
		logger.Debug("Struct Scanning Started")
		err = rows.StructScan(&author)
		if err != nil {
			logger.Error("Failed to scan struct")
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

//GetByID allows you to get Author by ID
func (q *Store) GetByID(id string) (*Author, error) {
	logger := q.logger.With("operation", "GetAuthor")
	logger.Debug("Start Query")
	b, err := q.db.NamedQuery("SELECT * FROM authors Where id=:id", map[string]interface{}{"id": id})
	if err != nil {
		logger.Error("Failed to Get Author", "error", err)
		return nil, err
	}
	defer b.Close()
	if !b.Next() {
		return nil, nil
	}
	var author Author
	logger.Debug("Struct Scanning Started")
	err = b.StructScan(&author)
	if err != nil {
		logger.Error("Failed to Scan Struct")
		return nil, err
	}
	return &author, nil
}

//Create allows you to Create a New Author
func (q *Store) Create(author *Author) error {
	logger := q.logger.With("operation", "PostAuthor")
	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`INSERT INTO authors (first_name, last_name, user_name, specialization, book_list) VALUES (:first_name, :last_name, :user_name, :specialization, :book_list)`,
		map[string]interface{}{
			"first_name":     author.Firstname,
			"last_name":      author.Lastname,
			"user_name":      author.Username,
			"specialization": author.Specialization,
			"book_list":      author.Booklist,
		})

	if err != nil {
		logger.Error("Failed to Create Author", "error", err)
		return err
	}
	return nil
}

//Update allows you to Update Author's info
func (q *Store) Update(author *Author, id string) error {
	logger := q.logger.With("operation", "UpdateAuthor")
	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`UPDATE authors SET first_name =:first_name,
									last_name =:last_name,
									user_name =:user_name,
									specialization =:specialization,
									book_list =:book_list
								WHERE id =:id`,
		map[string]interface{}{
			"id":             id,
			"first_name":     author.Firstname,
			"last_name":      author.Lastname,
			"user_name":      author.Username,
			"specialization": author.Specialization,
			"book_list":      author.Booklist,
		})
	if err != nil {
		logger.Error("Failed to Update Author", "error", err)
		return err
	}
	return nil
}

//Delete allows you to Remove Some Author
func (q *Store) Delete(id string) error {
	logger := q.logger.With("operation", "DeleteAuthor")
	_, err := q.db.NamedExec(`DELETE FROM authors
								Where id =:id`, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logger.Error("Failed to Delete Author")
		return err
	}
	return nil
}

//GetAuthorsBooks allows you to get all Author's Books
func (q *Store) GetAuthorsBooks(id string) ([]book.Book, error) {
	logger := q.logger.With("operation", "Get Authors books", "id", id)
	var book []book.Book
	err := q.db.Select(&book, `SELECT * FROM books WHERE ID = $1`, id)
	if err != nil {
		logger.Error("Failed to Get Author`s Books", "error", err)
	}
	return book, nil

}
