package book

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	book := Book{}
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{book.Title, book.ISBN, book.Genre})
	mock.ExpectQuery("SELECT (.*) FROM books").
		WillReturnRows(rows)
	storeAuthor := NewStore(sqlxDB, logger)

	if _, err = storeAuthor.GetAll(); err != nil {
		t.Errorf("error was not expected while deleting Author: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	defer sqlxDB.Close()
}
func TestGetBook(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	book := Book{}
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{book.Title, book.ISBN, book.Genre})
	mock.ExpectQuery("SELECT (.*) FROM books").
		WithArgs("1").
		WillReturnRows(rows)
	storeAuthor := NewStore(sqlxDB, logger)

	if _, err = storeAuthor.GetByID("1"); err != nil {
		t.Errorf("error was not expected while deleting Author: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	defer sqlxDB.Close()
}
func TestCreateReader(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	book := Book{}
	book.Title = "fawfawf"
	book.Genre = "fawfawf"
	book.ISBN = "fawfawf"

	mock.ExpectExec("INSERT INTO books").
		WithArgs(book.Title, book.Genre, book.ISBN).
		WillReturnResult(sqlmock.NewResult(1, 1))

	storeReader := NewStore(sqlxDB, logger)
	if err = storeReader.Create(&book); err != nil {
		t.Errorf("error was not expected while inserting stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	defer sqlxDB.Close()
}

func TestUpdateStore(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	book := Book{}
	book.Title = "fawfawf"
	book.Genre = "fawfawf"
	book.AuthorID = "1"

	mock.ExpectExec("UPDATE books").
		WithArgs(book.Title, book.Genre, "", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	storeAuthor := NewStore(sqlxDB, logger)

	if err = storeAuthor.Update(&book, "1"); err != nil {
		t.Errorf("error was not expected while inserting stats: %s", err)

	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	defer sqlxDB.Close()
}

func TestDeleteStore(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	mock.ExpectExec("DELETE FROM books").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	storeAuthor := NewStore(sqlxDB, logger)

	if err = storeAuthor.Delete("1"); err != nil {
		t.Errorf("error was not expected while deleting Author: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	defer sqlxDB.Close()
}
