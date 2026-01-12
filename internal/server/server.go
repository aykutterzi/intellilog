package server

import (
	"github.com/aykutterzi/intellilog/internal/ai"
	"github.com/aykutterzi/intellilog/internal/handlers"
	"github.com/aykutterzi/intellilog/internal/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo        *echo.Echo
	logHandler  *handlers.LogHandler
	authHandler *handlers.AuthHandler
}

func NewServer() *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Static files
	e.Static("/", "web")

	// Store
	logStore := store.NewInMemoryLogStore()

	// AI
	aiService := ai.NewSimpleRuleBasedAI()

	// Handlers
	logHandler := handlers.NewLogHandler(logStore, aiService)
	authHandler := handlers.NewAuthHandler("super-secret-key")

	s := &Server{
		echo:        e,
		logHandler:  logHandler,
		authHandler: authHandler,
	}

	s.RegisterRoutes()

	return s
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}
