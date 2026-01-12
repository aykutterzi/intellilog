package server

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() {
	// Public routes
	s.echo.POST("/login", s.authHandler.Login)
	s.echo.GET("/health", s.healthHandler)

	// Protected routes
	api := s.echo.Group("/api")

	// Middleware
	config := echojwt.Config{
		SigningKey: []byte("super-secret-key"),
	}
	api.Use(echojwt.WithConfig(config))
	api.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	api.POST("/logs", s.logHandler.CreateLog)
	api.GET("/logs", s.logHandler.GetLogs)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
