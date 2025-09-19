package server

import (
	"log"

	"minilink/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app    *fiber.App
	config *config.Config
}

func New(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	app.Use(logger.New())

	server := &Server{
		app:    app,
		config: cfg,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	s.app.Get("/:shortcode", s.handleRedirect)
	s.app.Get("/", s.handleHome)
}

func (s *Server) handleHome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "MiniLink URL Shortener",
		"status":  "running",
	})
}

func (s *Server) handleRedirect(c *fiber.Ctx) error {
	shortcode := c.Params("shortcode")
	queryString := string(c.Request().URI().QueryString())

	if url, exists := s.config.FindRoute(shortcode, queryString); exists {
		return c.Redirect(url, 301)
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "Short URL not found",
		"code":  shortcode,
	})
}

func (s *Server) Start(port string) error {
	log.Printf("Starting server on port %s", port)
	return s.app.Listen(":" + port)
}