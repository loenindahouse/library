package author

import (
	"github.com/hashicorp/go-hclog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service struct {
	store  *Store
	logger hclog.Logger
}

//Create New Service constructor
func NewService(store *Store, logger hclog.Logger) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}

//Binds a Author`s Handlers to a Router
func (s *Service) BindHandlers(a *echo.Group) {
	//Authors routing
	a.GET("", s.getAll)
	a.GET("/:id", s.getByID)
	a.POST("", s.create)
	a.PUT("/:id", s.update)
	a.DELETE("/:id", s.delete)
	a.GET("/:id/books", s.getAuthorsBooks)
}

//getAllAuthors allows you to Get All Authors
func (s *Service) getAll(c echo.Context) error {
	logger := s.logger.With("handler", "GetAuthors")
	logger.Debug("Get authors from db")
	reqBody, err := s.store.GetAll()
	if err != nil {
		logger.Error("Failed to get authors", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to get authors")
	}

	return c.JSON(http.StatusOK, reqBody)
}

//getAuthorByID allows you to Get Author by ID
func (s *Service) getByID(c echo.Context) error {
	logger := s.logger.With("handler", "GetAuthor")
	logger.Debug("Get parameter from db")
	id := c.Param("id")
	logger.Debug("Check if such author exists")
	author, err := s.store.GetByID(id)
	if err != nil {
		logger.Error("Incorrect parameter", " error", err)
		return c.String(http.StatusBadRequest, "400 Bad Request")
	}
	if author == nil {

		return c.String(http.StatusNotFound, "404 Not Found")
	}

	c.Response().Header().Set("Content-Type", "application/json")
	logger.Debug("Serialization of author")
	return c.JSON(http.StatusOK, author)
}

//postAuthor allows you to Create a new Author
func (s *Service) create(c echo.Context) (err error) {
	logger := s.logger.With("handler", "PostAuthor")
	logger.Debug("Author Deserialization")
	u := new(Author)
	if err = c.Bind(u); err != nil {
		logger.Error("Failed to Deserialize", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Deserialize")
	}
	logger.Debug("Creating Author")
	b := s.store.Create(u)
	return c.JSON(http.StatusCreated, b)
}

//updateAuthor allow you to Update Author
func (s *Service) update(c echo.Context) (err error) {
	logger := s.logger.With("handler", "UpdateAuthor")
	id := c.Param("id")
	logger.Debug("Get Author's parameter", id)
	u := new(Author)
	if err = c.Bind(u); err != nil {
		logger.Error("Failed to Deserialize Author", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Deserialize Author")
	}
	logger.Debug("Check if such Author exists")
	b := s.store.Update(u, id)
	logger.Debug("Updating Author")
	return c.JSON(http.StatusOK, b)
}

//deleteAuthor allow you to Remove some Author
func (s *Service) delete(c echo.Context) (err error) {
	logger := s.logger.With("handler", "DeleteAuthor")
	logger.Debug("Get Parameter")
	id := c.Param("id")
	logger.Debug("Check if such Author exists", id)
	b := s.store.Delete(id)
	logger.Debug("Removing an Author")
	return c.JSON(http.StatusOK, b)
}

//getAuthorsBooks allows you to Get Author`s Books
func (s *Service) getAuthorsBooks(c echo.Context) (err error) {
	logger := s.logger.With("handler", "GetAuthorsBooks")
	id := c.Param("id")
	logger.Debug("Get parameter", id)
	res, err := s.store.GetByID(id)
	if err != nil {
		logger.Error("Failed to Get Author", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Get Author")
	}
	if res == nil {
		logger.Error("404 Not Found", "error", err)
		return c.String(http.StatusNotFound, "404 Not found")
	}
	logger.Debug("Get Author`s Books")
	books, err := s.store.GetAuthorsBooks(id)
	if err != nil {
		logger.Error("Failed to Get Author`s Books", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Get Author`s Books")
	}

	return c.JSON(http.StatusCreated, books)
}
