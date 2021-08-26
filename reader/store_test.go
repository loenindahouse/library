package reader

import (
	"github.com/DATA-DOG/go-sqlmock"

	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"testing"
)

func TestGetAllReaders(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	reader := Reader{}
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{reader.FirstName, reader.LastName, reader.Booklist})
	mock.ExpectQuery("SELECT (.*) FROM reader").
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
func TestGetReader(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Logger",
		Level:      hclog.Debug,
		JSONFormat: false,
	})
	reader := Reader{}
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{reader.FirstName, reader.LastName, reader.Booklist})
	mock.ExpectQuery("SELECT (.*) FROM reader").
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
	reader := Reader{}
	reader.FirstName = "fawfawf"
	reader.LastName = "fawfawf"
	reader.Booklist = "1,2"

	mock.ExpectExec("INSERT INTO reader").
		WithArgs(reader.FirstName, reader.LastName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	storeReader := NewStore(sqlxDB, logger)
	if err = storeReader.Create(&reader); err != nil {
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
	reader := Reader{}
	reader.FirstName = "fawfawf"
	reader.LastName = "fawfawf"
	reader.Booklist = "1"

	mock.ExpectExec("UPDATE reader").
		WithArgs(reader.FirstName, reader.LastName, "1", reader.Booklist).
		WillReturnResult(sqlmock.NewResult(1, 1))
	storeAuthor := NewStore(sqlxDB, logger)

	if err = storeAuthor.Update(&reader, "1"); err != nil {
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
	mock.ExpectExec("DELETE FROM reader").
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
