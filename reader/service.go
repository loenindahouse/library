package reader

import (
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
	store  *Store
	logger hclog.Logger
}

//Create New Service Construction
func NewService(store *Store, logger hclog.Logger) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}

//Binds a Reader`s Handlers to a Router
func (s *Service) BindHandlers(r *echo.Group) {
	//Readers routing
	r.GET("", s.getAll)
	r.GET("/:id", s.getByID)
	r.POST("", s.create)
	r.PUT("/:id", s.update)
	r.DELETE("/:id", s.delete)
	r.GET("/:id/books", s.getReadersBook)
}

//getAllReaders allows you to get all Readers
func (s *Service) getAll(c echo.Context) error {

	logger := s.logger.With("handler", "GetReaders")
	logger.Debug("Get all readers from db")

	reqBody, err := s.store.GetAll()
	if err != nil {
		logger.Error("Failed to get readers", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to get readers")
	}
	logger.Debug("Serialization of readers")
	return c.JSON(http.StatusOK, reqBody)
}

//getReaderByID allows you to get reader by id
func (s *Service) getByID(c echo.Context) error {
	logger := s.logger.With("handler", "GetReader")
	id := c.Param("id")
	logger.Debug("Get parameter from db", id)
	reader, err := s.store.GetByID(id)
	if err != nil {
		logger.Error("Failed to get reader", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to get reader")
	}

	logger.Debug("Check if such reader exists")
	if reader == nil {
		logger.Error("Reader Not Found", "error", err)
		return c.String(http.StatusNotFound, "404 Not Found")
	}

	c.Response().Header().Set("Content-Type", "application/json")

	logger.Debug("Serialization of reader")
	return c.JSON(http.StatusOK, reader)
}

//postReader allows you to create a new Reader
func (s *Service) create(c echo.Context) (err error) {
	logger := s.logger.With("handler", "PostReaders")
	logger.Debug("Reader Deserialization")
	u := new(Reader)
	if err = c.Bind(u); err != nil {
		logger.Error("Failed to Deserialize", "error", err)

		return c.String(http.StatusInternalServerError, "Failed to Deserialize")
	}
	logger.Debug("Creating reader")
	b := s.store.Create(u)
	return c.JSON(http.StatusCreated, b)
}

//updateReader allows you to Update Reader`s information
func (s *Service) update(c echo.Context) (err error) {
	logger := s.logger.With("handler", "UpdateReaders")
	id := c.Param("id")
	logger.Debug("Get parameter", id)
	u := new(Reader)
	if err = c.Bind(u); err != nil {
		logger.Error("Failed to Deserialize", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to Deserialize")
	}
	logger.Debug("Update readers")
	b := s.store.Update(u, id)
	return c.JSON(http.StatusOK, b)
}

//deleteReader allows you to remove some Reader
func (s *Service) delete(c echo.Context) (err error) {
	logger := s.logger.With("handler", "DeleteReaders")
	id := c.Param("id")
	logger.Debug("Check if such reader exists", id)
	b := s.store.Delete(id)
	logger.Debug("Removing a Reader")
	return c.JSON(http.StatusOK, b)
}

//getReadersBook allows you to Get Reader`s books
func (s *Service) getReadersBook(c echo.Context) (err error) {
	logger := s.logger.With("handler", "GetReadersBook")
	id := c.Param("id")
	logger.Debug("get parameter", id)
	res, err := s.store.GetByID(id)
	if err != nil {
		logger.Error("Failed to get Reader", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to get Reader")
	}
	if res == nil {
		logger.Error("Not Found", "error", err)
		return c.String(http.StatusNotFound, "404 Not found")
	}

	logger.Debug("Check if such Reader exists")
	books, err := s.store.GetReadersBook(id)
	if err != nil {
		logger.Error("Failed to get Reader`s Books", "error", err)
		return c.String(http.StatusInternalServerError, "Failed to get Reader`s Books")
	}
	logger.Debug("Marshal result")
	return c.JSON(http.StatusCreated, books)
}
