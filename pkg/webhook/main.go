package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thisiserico/weasley-clock/pkg/firebase"
	"github.com/thisiserico/weasley-clock/pkg/netlify"
	"github.com/thisiserico/weasley-clock/pkg/responder"
	"github.com/thisiserico/weasley-clock/pkg/telegram"
	"golang.org/x/sync/errgroup"
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
	return r.status() == "/start"
}

func (r request) status() string {
	return strings.ToLower(r.Message.Text)
}

func (r request) chatID() int {
	return r.Message.From.ID
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	req, err := decodeRequest(r)
	if err != nil {
		return responder.Respond(http.StatusBadRequest, nil, err)
	}

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
		return responder.Respond(http.StatusOK, nil, err)
	}

	if req.isInitialMessage() {
		return responder.Respond(
			http.StatusOK,
			nil,
			telegram.RequestNextStatus(req.chatID(), allStatuses),
		)
	}

	name, err := netlify.FetchMembers().NameFromChat(req.chatID())
	if err != nil {
		return responder.Respond(http.StatusOK, nil, err)
	}

	if !allPeople.Exists(name) {
		err := fmt.Errorf("unrecognized member %s", name)
		return responder.Respond(http.StatusOK, nil, err)
	}

	if !allStatuses.Exists(req.status()) {
		err := fmt.Errorf("unrecognized status %s", req.status())
		return responder.Respond(http.StatusOK, nil, err)
	}

	group, ctx = errgroup.WithContext(ctx)
	group.Go(func() error {
		return firebase.UpdateLocation(name, req.status())
	})
	group.Go(func() error {
		return telegram.RequestNextStatus(req.chatID(), allStatuses)
	})
	return responder.Respond(http.StatusOK, nil, group.Wait())
}

func decodeRequest(r events.APIGatewayProxyRequest) (request, error) {
	var req request
	decoder := json.NewDecoder(bytes.NewBufferString(r.Body))
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	return req, nil
}
