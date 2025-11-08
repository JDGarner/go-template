package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/JDGarner/go-template/internal/handlers"
	"github.com/JDGarner/go-template/internal/store"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	serverTimeout = 10 * time.Second
)

type Server struct {
	echo *echo.Echo
	port string
}

func New(s *store.Store, port string) *Server {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	h := &handlers.DummyHandler{
		Store: s,
	}
	registerRoutes(e, h)

	return &Server{
		echo: e,
		port: port,
	}
}

func (s *Server) Start() error {
	err := s.echo.Start(":" + s.port)
	if err != nil && err != http.ErrServerClosed {
		slog.Error("server error", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), serverTimeout)
	defer cancel()

	return s.echo.Shutdown(shutdownCtx)
}

func registerRoutes(e *echo.Echo, h *handlers.DummyHandler) {
	e.GET("/item/:id", h.GetItem)
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}
