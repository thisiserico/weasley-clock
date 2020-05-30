package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	acceptedMembersEnvVar   = "ACCEPTED_MEMBERS"
	acceptedLocationsEnvVar = "ACCEPTED_LOCATIONS"

	firebaseAddressEnvVar = "FIREBASE_ADDRESS"
)

type request struct {
	Message struct {
		From struct {
			ID int `json:"id"`
		} `json:"from"`
		Text string `json:"text"`
	} `json:"message"`
}

func (r request) isInitialMessage() bool {
	return r.Message.Text == "/start"
}

type members map[int]string

func (m members) name(req request) (string, error) {
	chatID := req.Message.From.ID
	if name, found := m[chatID]; found {
		return name, nil
	}

	return "", fmt.Errorf("unrecognized chat ID %d", chatID)
}

type locations []string

func (l locations) match(r request) (string, error) {
	givenLocation := r.Message.Text
	for _, location := range l {
		if location == givenLocation {
			return location, nil
		}
	}

	return "", fmt.Errorf("unrecognized location %s", givenLocation)
}

func main() {
	lambda.Start(handler)
}

func handler(r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	req, err := decodeRequest(r)
	if err != nil {
		return respond(http.StatusBadRequest, err)
	}

	if req.isInitialMessage() {
		return respond(http.StatusOK, nil)
	}

	members, err := allowedMembers()
	if err != nil {
		return respond(http.StatusOK, err)
	}

	name, err := members.name(req)
	if err != nil {
		return respond(http.StatusOK, err)
	}

	locations := allowedLocations()
	location, err := locations.match(req)
	if err != nil {
		return respond(http.StatusOK, err)
	}

	return respond(http.StatusOK, updateLocation(name, location))
}

func decodeRequest(r events.APIGatewayProxyRequest) (request, error) {
	var req request
	decoder := json.NewDecoder(bytes.NewBufferString(r.Body))
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	return req, nil
}

func allowedMembers() (members, error) {
	acceptedMembers := os.Getenv(acceptedMembersEnvVar)

	memberExtractor := regexp.MustCompile(`([a-z]+):(\d+)`)
	if !memberExtractor.MatchString(acceptedMembers) {
		return nil, fmt.Errorf("invalid member list %s", acceptedMembers)
	}

	memberMatches := memberExtractor.FindAllStringSubmatch(acceptedMembers, -1)
	mems := make(members, len(memberMatches))
	for _, match := range memberMatches {
		chatID, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("invalid chat ID %s", match[2])
		}

		mems[chatID] = match[1]
	}

	return mems, nil
}

func allowedLocations() locations {
	return strings.Split(os.Getenv(acceptedLocationsEnvVar), " ")
}

func updateLocation(name, location string) error {
	endpoint, _ := url.Parse(os.Getenv(firebaseAddressEnvVar))
	endpoint.Path = path.Join(endpoint.Path, "location.json")

	payload := map[string]string{name: location}
	js, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, endpoint.String(), bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)

	return err
}

func respond(code int, err error) (*events.APIGatewayProxyResponse, error) {
	resp := &events.APIGatewayProxyResponse{StatusCode: code}
	if err != nil {
		log.Println(err)
		resp.Body = err.Error()
	}

	return resp, nil
}
