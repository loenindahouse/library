package main

import (
	"artemstudy/author"
	"artemstudy/book"
	"artemstudy/database"
	"artemstudy/reader"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var port int
	var dbStr string
	var jsonFormat bool

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "dbStr",
			EnvVar:      "LIBRARY_DBCONN",
			Usage:       "String for db connection",
			Destination: &dbStr,
			Required:    true,
		},
		&cli.IntFlag{
			Name:        "port",
			EnvVar:      "LIBRARY_PORT",
			Destination: &port,
			Required:    true,
		},
		&cli.BoolFlag{
			Name:        "json_log",
			EnvVar:      "LIBRARY_JSON",
			Destination: &jsonFormat,
			Usage:       "check for JSON Format",
		},
	}

	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "library-rest-api",
		Level:      hclog.Debug,
		JSONFormat: false,
	})

	app.Action = func(c *cli.Context) error {
		logger.Debug("Service started")
		//Run DB connection
		db, err := database.Connect(dbStr, logger)
		if err != nil {
			logger.Error("Failed to connect to db", "error", err)
			return err
		}
		logger.Debug("Start Migration")
		//Run DB Migration
		err = database.Migrate(dbStr, logger)
		if err != nil {
			logger.Error("Failed to migrate db", "error", err)
			return err
		}
		//Create new Router
		e := echo.New()

		bookStore := book.NewStore(db, logger.Named("bookStore"))
		bookService := book.NewService(bookStore, logger.Named("bookService"))
		b := e.Group("/books")
		bookService.BindHandlers(b)

		authorStore := author.NewStore(db, logger.Named("authorStore"))
		authorService := author.NewService(authorStore, logger.Named("authorService"))
		a := e.Group("/authors")
		authorService.BindHandlers(a)

		readerStore := reader.NewStore(db, logger.Named("readerStore"))
		readerService := reader.NewService(readerStore, logger.Named("readerService"))
		r := e.Group("/readers")
		readerService.BindHandlers(r)

		//Start server
		logger.Debug("Server is listening")
		err = e.Start(fmt.Sprintf(":%d", port))
		if err != nil {
			logger.Error("Server Can't start")
			return err
		}

		return nil
	}
	//Run Action function with Environment
	err := app.Run(os.Args)
	if err != nil {
		logger.Error("Failed to run", "error", err)
		os.Exit(1)
	}
}
