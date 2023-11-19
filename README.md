# Logs Ingestor and Query Searcher


### How to run:


`go mod tidy`

`go run main.go`


## To Ingest logs or Fetch it

Ingest:

you can make a curl request for example :

`curl -X POST -H "Content-Type: application/json" -d '{
  "level": "info",
  "message": "Log message",
  "resourceId": "123",
  "timestamp": "2023-01-01T12:34:56Z",
  "traceId": "456",
  "spanId": "789",
  "commit": "abc",
  "metadata": {
    "parentResourceId": "789"
  }
}' http://localhost:3000/ingest`


OR

you can visit http://localhost:3000/ingest-form and manually fill in the logs


Fetch:

you can make a curl request for example :

`curl "http://localhost:3000/search?level=warning"`

OR 

you can visit http://localhost:3000/ingest-form and manually search using any parameter you lik.

