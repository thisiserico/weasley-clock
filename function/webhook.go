package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	acceptedMembersEnvVar   = "ACCEPTED_MEMBERS"
	acceptedLocationsEnvVar = "ACCEPTED_LOCATIONS"
)

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
		return response(http.StatusBadRequest, err)
	}

	if req.Message.Text == "/start" {
		return response(http.StatusOK, nil)
	}

	acceptedMembers := os.Getenv(acceptedMembersEnvVar)
	memberExtractor := regexp.MustCompile(`([a-z]+):(\d+)`)
	if !memberExtractor.MatchString(acceptedMembers) {
		log.Printf("invalid member list %s", acceptedMembers)
		return response(http.StatusOK, nil)
	}

	memberMatches := memberExtractor.FindAllStringSubmatch(acceptedMembers, -1)
	members := make(map[int]string, len(memberMatches))
	for _, match := range memberMatches {
		chatID, err := strconv.Atoi(match[2])
		if err != nil {
			log.Printf("invalid chat ID %s", match[2])
			return response(http.StatusOK, nil)
		}
		members[chatID] = match[1]
	}

	name, allowed := members[req.Message.From.ID]
	if !allowed {
		log.Printf("unrecognized chat ID %d", req.Message.From.ID)
		return response(http.StatusOK, nil)
	}

	locations := strings.Split(os.Getenv(acceptedLocationsEnvVar), " ")
	var allowedLocation bool
	for _, location := range locations {
		if location == req.Message.Text {
			allowedLocation = true
			break
		}
	}
	if !allowedLocation {
		log.Printf("unrecognized location %s", req.Message.Text)
		return response(http.StatusOK, nil)
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
