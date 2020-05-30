package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

type request struct {
	Message struct {
		Date int64 `json:"date"`
		From struct {
			ID int `json:"id"`
		} `json:"from"`
		Text string `json:"text"`
	} `json:"message"`
}

func handler(r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req request
	if err := json.NewDecoder(bytes.NewBufferString(r.Body)).Decode(&req); err != nil {
		return response(http.StatusInternalServerError, err)
	}

	log.Printf(
		"from %d at %s: %s",
		req.Message.From.ID,
		time.Unix(req.Message.Date, 0),
		req.Message.Text,
	)

	return response(http.StatusOK, nil)
}

func response(code int, err error) (*events.APIGatewayProxyResponse, error) {
	resp := &events.APIGatewayProxyResponse{StatusCode: code}
	if err != nil {
		resp.Body = err.Error()
	}

	return resp, nil
}
