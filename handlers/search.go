package handlers

import (
	"context"
	"fmt"
	"log_ingestor/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchHandler(c *fiber.Ctx) error {
	// Get query parameters
	q := c.Query("q")
	level := c.Query("level")
	resourceId := c.Query("resourceId")
	traceId := c.Query("traceId")
	spanId := c.Query("spanId")
	commit := c.Query("commit")
	parentResourceId := c.Query("metadata.parentResourceId")

	// Filter logs
	filteredLogs := filterLogs(q, level, resourceId, traceId, spanId, commit, parentResourceId)

	// Return the filtered logs
	return c.JSON(filteredLogs)
}

func filterLogs(q, level, resourceId, traceId, spanId, commit, parentResourceId string) []models.Log {
	var filteredLogs []models.Log
	filter := bson.M{}
	if q != "" {
		filter["message"] = bson.M{"$regex": primitive.Regex{Pattern: q, Options: "i"}}
	}
	if level != "" {
		filter["level"] = level
	}
	if resourceId != "" {
		filter["resourceId"] = resourceId
	}
	if traceId != "" {
		filter["traceId"] = traceId
	}
	if spanId != "" {
		filter["spanId"] = spanId
	}
	if commit != "" {
		filter["commit"] = commit
	}
	if parentResourceId != "" {
		filter["metadata.parentResourceId"] = parentResourceId
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("Error querying logs from MongoDB:", err)
		return filteredLogs
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var log models.Log
		if err := cur.Decode(&log); err != nil {
			fmt.Println("Error decoding log:", err)
			continue
		}
		filteredLogs = append(filteredLogs, log)
	}

	return filteredLogs
}

func SearchFormHandler(c *fiber.Ctx) error {
	html := ` 
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Log Search</title>
	</head>
	<body>
		<h1>Log Search</h1>
		<form action="/search" method="get">
			<label for="q">Search Query:</label>
			<input type="text" name="q"><br>
			<label for="level">Level:</label>
			<input type="text" name="level"><br>
			<label for="resourceId">Resource ID:</label>
			<input type="text" name="resourceId"><br>
			<label for="traceId">Trace ID:</label>
			<input type="text" name="traceId"><br>
			<label for="spanId">Span ID:</label>
			<input type="text" name="spanId"><br>
			<label for="commit">Commit:</label>
			<input type="text" name="commit"><br>
			<label for="parentResourceId">Parent Resource ID:</label>
			<input type="text" name="parentResourceId"><br>
			<button type="submit">Search</button>
		</form>
	</body>
	</html>
`
	c.Set("Content-Type", "text/html")
	return c.Status(fiber.StatusOK).SendString(html)
}
