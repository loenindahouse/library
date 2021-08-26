package book

import (
	// "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db     *sqlx.DB
	logger hclog.Logger
}

//Create New Store Construction
func NewStore(db *sqlx.DB, logger hclog.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

//GetAll allows you to get all Books
func (q *Store) GetAll() (Books, error) {
	logger := q.logger.With("operation", "GetBooks")
	var books Books
	logger.Debug("Start Query")
	rows, err := q.db.Queryx("SELECT * FROM books")
	if err != nil {
		logger.Error("Failed to Get Books", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		logger.Debug("Struct Scanning Started")
		err = rows.StructScan(&book)
		if err != nil {
			logger.Error("Failed to scan struct", "error", err)
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

//GetByID allows you to Get By ID
func (q *Store) GetByID(id string) (*Book, error) {
	var book Book
	logger := q.logger.With("operation", "GetBook")
	logger.Debug("Start Query")
	b, err := q.db.NamedQuery("SELECT * FROM books Where id=:id", map[string]interface{}{"id": id})
	if err != nil {
		logger.Error("Failed to Get Book", "error", err)
		return &book, err
	}
	defer b.Close()

	if !b.Next() {
		return &book, nil
	}

	logger.Debug("Struct Scanning Started")

	err = b.StructScan(&book)
	if err != nil {
		logger.Error("Failed to Scan Struct", "error", err)
		return &book, err
	}
	return &book, nil
}

//Create allows you to Create New Book
func (q *Store) Create(book *Book) error {
	logger := q.logger.With("operation", "POstBook")
	_, err := q.db.NamedExec(`INSERT INTO books (title, genre, isbn) VALUES (:title, :genre, :isbn)`,
		map[string]interface{}{
			"title": book.Title,
			"genre": book.Genre,
			"isbn":  book.ISBN,
		})

	if err != nil {
		logger.Error("Failed to Create Book", "error", err)
		return err
	}
	return nil
}

//Update allows you to Update Book
func (q *Store) Update(book *Book, id string) error {
	logger := q.logger.With("operation", "UpdateBook")

	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`UPDATE books SET title =:title,
									genre =:genre,
									isbn =:isbn,
								WHERE id =:id`,
		map[string]interface{}{
			"id":        id,
			"title":     book.Title,
			"genre":     book.Genre,
			"isbn":      book.ISBN,
			"author_id": book.AuthorID,
		})
	if err != nil {
		logger.Error("Failed to Update Book")
		return err
	}
	return nil
}

//Delete allows you to Remove Some Book
func (q *Store) Delete(id string) error {
	logger := q.logger.With("operation", "DeleteBook")
	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`DELETE FROM books
								Where id =:id`, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logger.Error("Failed to Delete Book", "error", err)
		return err
	}
	return nil
}
