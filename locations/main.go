package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const firebaseAddressEnvVar = "FIREBASE_ADDRESS"

type locations map[string]string

func main() {
	lambda.Start(handler)
}

func handler(_ events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	locs, err := fetchLocations()
	if err != nil {
		return respond(http.StatusInternalServerError, err)
	}

	return respond(http.StatusOK, locs)
}

func fetchLocations() (locations, error) {
	endpoint, _ := url.Parse(os.Getenv(firebaseAddressEnvVar))
	endpoint.Path = path.Join(endpoint.Path, "location.json")

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locs locations
	if err := json.NewDecoder(resp.Body).Decode(&locs); err != nil {
		return nil, err
	}

	return locs, nil
}

func respond(code int, resp interface{}) (*events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(resp)
	if err != nil {
		body = []byte(err.Error())
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(body),
	}, nil
}
