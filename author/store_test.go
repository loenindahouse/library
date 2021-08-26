package author

import (
	"artemstudy/book"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"testing"
)

func TestGetAllAuthors(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	author := Author{}
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{author.Firstname, author.Lastname, author.Username, author.Specialization, author.Booklist})
	mock.ExpectQuery("SELECT (.*) FROM authors").
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
func TestGetAuthor(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	author := Author{}
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{author.Firstname, author.Lastname, author.Username, author.Specialization, author.Booklist})
	mock.ExpectQuery("SELECT (.*) FROM authors").
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
func TestCreateAuthor(t *testing.T) {
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
	author := Author{}
	author.Firstname = "fawfawf"
	author.Lastname = "fawfawf"
	author.Username = "wfafaw"
	author.Specialization = "fwafaw"
	author.Booklist = "1,2"
	mock.ExpectExec("INSERT INTO authors").
		WithArgs(author.Firstname, author.Lastname, author.Username, author.Specialization, author.Booklist).
		WillReturnResult(sqlmock.NewResult(1, 1))

	storeAuthor := NewStore(sqlxDB, logger)
	if err = storeAuthor.Create(&author); err != nil {
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
	author := Author{}
	author.Firstname = "fawfawf"
	author.Lastname = "fawfawf"
	author.Username = "wfafaw"
	author.Specialization = "fawfawf"
	author.Booklist = "1,2"
	mock.ExpectExec("UPDATE authors").
		WithArgs(author.Firstname, author.Lastname, author.Username, author.Specialization, author.Booklist, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	storeAuthor := NewStore(sqlxDB, logger)

	if err = storeAuthor.Update(&author, "1"); err != nil {
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
	mock.ExpectExec("DELETE FROM authors").
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

func TestGetAuthorsBooks(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})

	books := book.Book{}
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Error("failed", "error", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{books.Title, books.ISBN, books.Genre, books.AuthorID})
	mock.ExpectQuery("SELECT (.*) FROM books WHERE").WithArgs("1").WillReturnRows(rows)
	storeAuthor := NewStore(sqlxDB, logger)
	if _, err = storeAuthor.GetAuthorsBooks("1"); err != nil {
		logger.Error("error was not expected while geting Author`s books:", "error", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	defer sqlxDB.Close()

}
