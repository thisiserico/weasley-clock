package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thisiserico/weasley-clock/pkg/firebase"
	"github.com/thisiserico/weasley-clock/pkg/responder"
	"golang.org/x/sync/errgroup"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, _ events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	group, ctx := errgroup.WithContext(ctx)

	var allStatuses firebase.Statuses
	group.Go(func() error {
		var err error
		allStatuses, err = firebase.FetchStatuses()
		return err
	})

	var allPeople firebase.People
	group.Go(func() error {
		var err error
		allPeople, err = firebase.FetchPeople()
		return err
	})

	if err := group.Wait(); err != nil {
		return responder.Respond(http.StatusInternalServerError, nil, err)
	}

	type response struct {
		Statuses firebase.Statuses `json:"statuses"`
		People   firebase.People   `json:"people"`
	}

	return responder.Respond(http.StatusOK, response{
		Statuses: allStatuses,
		People:   allPeople,
	}, nil)
}
