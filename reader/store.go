package reader

import (
	"artemstudy/book"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Store struct {
	db     *sqlx.DB
	logger hclog.Logger
}

//Create New Store Constructor
func NewStore(db *sqlx.DB, logger hclog.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

//GetAll allows you to get all Readers
func (q *Store) GetAll() (Readers, error) {

	var readers Readers
	logger := q.logger.With("operation", "GetReaders")
	logger.Debug("Start Query")
	rows, err := q.db.Queryx("SELECT * FROM reader")

	if err != nil {
		logger.Error("Failed to Get Readers", " error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reader Reader
		logger.Debug("Struct Scanning started")
		err = rows.StructScan(&reader)
		if err != nil {
			logger.Error("Failed to scan struct")
			return nil, err
		}
		readers = append(readers, reader)
	}
	return readers, nil
}

//GetByID allows you to get Reader by ID
func (q *Store) GetByID(id string) (*Reader, error) {
	var reader Reader
	logger := q.logger.With("operation", "GetReader")
	logger.Debug("Start Query")
	b, err := q.db.NamedQuery("SELECT * FROM reader Where id=:id",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		logger.Error("Failed to Get Reader")
		return nil, err
	}
	defer b.Close()

	if !b.Next() {
		return &reader, nil
	}
	logger.Debug("Struct Scanner Started")
	err = b.StructScan(&reader)
	if err != nil {
		logger.Error("Failed to Scan Struct", "error", err)
		return &reader, err
	}
	return &reader, nil
}

//Create allows you to Create New Reader
func (q *Store) Create(reader *Reader) error {
	logger := q.logger.With("operation", "CreateReader")
	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`INSERT INTO reader (first_name, last_name) VALUES (:first_name, :last_name)`,
		map[string]interface{}{
			"first_name": reader.FirstName,
			"last_name":  reader.LastName,
		})

	if err != nil {
		logger.Error("Failed to Create Reader", "error", err)
		return err
	}
	return nil
}

//Update allows you to Update Readers's info
func (q *Store) Update(reader *Reader, id string) error {
	logger := q.logger.With("operation", "UpdateReader")
	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`UPDATE reader SET first_name =:first_name,
									last_name =:last_name,
									book_list =:book_list,
									WHERE id =:id`,
		map[string]interface{}{
			"id":         id,
			"first_name": reader.FirstName,
			"last_name":  reader.LastName,
			"book_list":  reader.Booklist,
		})
	if err != nil {
		logger.Error("Failed to Update Reader", "error", err)
		return err
	}
	return nil
}

//Delete allows you to Remove some Reader
func (q *Store) Delete(id string) error {
	logger := q.logger.With("operation", "DeleteReader")
	logger.Debug("Start Query")
	_, err := q.db.NamedExec(`DELETE FROM reader Where id =:id`,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		logger.Debug("Not Found")
		return err
	}
	return nil
}

//GetReadersBook allows you to get all of the reader's books
func (q *Store) GetReadersBook(id string) ([]book.Book, error) {
	logger := q.logger.With("operation", "Get readers books", "id", id)
	reader := Reader{}
	err := q.db.Get(&reader, "SELECT * FROM reader WHERE ID = $1", id)
	if err != nil {
		logger.Error("Failed to get Reader for Reader's books", "error", err)
		return nil, err
	}
	bid := strings.Split(reader.Booklist, ",")
	books := []book.Book{}
	for _, i := range bid {
		book := book.Book{}
		err := q.db.Get(&book, "SELECT * FROM books WHERE ID = $1", i)
		if err != nil {
			logger.Error("Failed to Get Reader's books", "error", err)
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
