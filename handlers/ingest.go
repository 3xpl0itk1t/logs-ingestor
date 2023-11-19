package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log_ingestor/models"

	"github.com/gofiber/fiber/v2"
)

func IngestHandler(c *fiber.Ctx) error {
	var log models.Log
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&log)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error decoding log")
	}

	// Save log to MongoDB
	_, err = collection.InsertOne(context.Background(), log)
	if err != nil {
		fmt.Println("Error inserting log into MongoDB:", err)
	}

	return c.SendString("Log ingested successfully")
}
func IngestFormHandler(c *fiber.Ctx) error {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Log Ingestor</title>
</head>
<body>
    <h1>Log Ingestor</h1>
    <form action="/ingest" method="post">
        <label for="level">Level:</label>
        <input type="text" name="level" required><br>
        <label for="message">Message:</label>
        <input type="text" name="message" required><br>
        <label for="resourceId">Resource ID:</label>
        <input type="text" name="resourceId" required><br>
        <label for="traceId">Trace ID:</label>
        <input type="text" name="traceId" required><br>
        <label for="spanId">Span ID:</label>
        <input type="text" name="spanId" required><br>
        <label for="commit">Commit:</label>
        <input type="text" name="commit" required><br>
        <label for="parentResourceId">Parent Resource ID:</label>
        <input type="text" name="parentResourceId" required><br>
        <button type="submit">Submit</button>
    </form>
</body>
</html>
`
	c.Set("Content-Type", "text/html")
	return c.Status(fiber.StatusOK).SendString(html)
}
