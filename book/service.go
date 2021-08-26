package book

import (
	"github.com/hashicorp/go-hclog"
	"net/http"
	// "io"
	// "strconv"

	"github.com/labstack/echo/v4"
)

type Service struct {
	store  *Store
	logger hclog.Logger
}

//Create New Service constuctions
func NewService(store *Store, logger hclog.Logger) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}

//Binds a Book`s Handlers to a Router
func (s *Service) BindHandlers(b *echo.Group) {
	//Books routing
	b.GET("", s.getAll)
	b.GET("/:id", s.getByID)
	b.POST("", s.create)
	b.PUT("/:id", s.update)
	b.DELETE("/:id", s.delete)
}

//getAllBooks allows you to Get All Books
func (s *Service) getAll(c echo.Context) error {
	logger := s.logger.With("handler", "GetAllBooks")
	logger.Debug("Get Books from DB")
	reqBody, err := s.store.GetAll()
	if err != nil {
		logger.Error("Failed to get books", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to get books")
	}

	return c.JSON(http.StatusOK, reqBody)
}

//getBook allows you to Get a Book by ID
func (s *Service) getByID(c echo.Context) error {
	logger := s.logger.With("handler", "GetBook")
	id := c.Param("id")
	logger.Debug("Get parameter", id)
	book, err := s.store.GetByID(id)
	if err != nil {
		logger.Error("Incorrect parameter", "error", err)
		return c.String(http.StatusBadRequest, "400 Bad Request")
	}
	if book == nil {
		logger.Error("Not Found", "error", err)
		return c.String(http.StatusNotFound, "Not Found")
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, book)
}

//postBook allows you to Create a new Book
func (s *Service) create(c echo.Context) (err error) {
	logger := s.logger.With("handler", "PostBook")
	logger.Debug("Book Deserialization")
	u := new(Book)
	if err = c.Bind(u); err != nil {
		logger.Error("Failed to Deserialize", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Deserialize")
	}
	logger.Debug("Creating Book")
	b := s.store.Create(u)
	return c.JSON(http.StatusCreated, b)
}

//updateBook allows you to Update Book`s information
func (s *Service) update(c echo.Context) (err error) {
	logger := s.logger.With("handler", "UpdateBook")
	id := c.Param("id")
	logger.Debug("Get parameter", id)
	u := new(Book)
	if err = c.Bind(u); err != nil {
		logger.Error("Failed to Deserialize Author", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Deserialize Author")
	}
	b := s.store.Update(u, id)
	return c.JSON(http.StatusOK, b)

}

//deleteBook allows you to Remove some Book
func (s *Service) delete(c echo.Context) (err error) {
	logger := s.logger.With("handler", "DeleteBook")
	logger.Debug("Get Parameter")
	id := c.Param("id")
	b := s.store.Delete(id)
	return c.JSON(http.StatusOK, b)
}
