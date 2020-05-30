package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const acceptedMembersEnvVar = "ACCEPTED_MEMBERS"

type request struct {
	Message struct {
		Date int64 `json:"date"`
		From struct {
			ID int `json:"id"`
		} `json:"from"`
		Text string `json:"text"`
	} `json:"message"`
}

func main() {
	lambda.Start(handler)
}

func handler(r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req request
	decoder := json.NewDecoder(bytes.NewBufferString(r.Body))
	if err := decoder.Decode(&req); err != nil {
		return response(http.StatusInternalServerError, err)
	}

	acceptedMembers := os.Getenv(acceptedMembersEnvVar)
	memberExtractor := regexp.MustCompile(`([a-z]+):(\d+)`)
	if !memberExtractor.MatchString(acceptedMembers) {
		log.Fatalf("invalid member list %s", acceptedMembers)
	}

	matches := memberExtractor.FindAllStringSubmatch(acceptedMembers, -1)
	members := make(map[int]string, len(matches))
	for _, match := range matches {
		chatID, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		members[chatID] = match[1]
	}

	name, allowed := members[req.Message.From.ID]
	if !allowed {
		return response(http.StatusBadRequest, errors.New("unrecognized chat"))
	}

	log.Printf("from %s at %s: %s", name, time.Unix(req.Message.Date, 0), req.Message.Text)

	return response(http.StatusOK, nil)
}

func response(code int, err error) (*events.APIGatewayProxyResponse, error) {
	resp := &events.APIGatewayProxyResponse{StatusCode: code}
	if err != nil {
		resp.Body = err.Error()
	}

	return resp, nil
}
