package server

import (
	"log"
	"log_ingestor/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func StartServer() {
	// Connect to the database
	handlers.ConnectToDB()
	defer handlers.DisconnectFromDB()

	// Start the server
	PORT := os.Getenv("PORT")
	app := fiber.New()
	// app.Use(middleware.Logger())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, world!",
		})
	})
	app.Get("/search", handlers.SearchHandler)
	app.Get("/search-form", handlers.SearchFormHandler)
	app.Post("/ingest", handlers.IngestHandler)
	app.Get("/ingest-form", handlers.IngestFormHandler)

	// Run the server in a goroutine
	go func() {
		err := app.Listen(":" + PORT)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Server is quitting, disconnect from the database
	handlers.DisconnectFromDB()
}
