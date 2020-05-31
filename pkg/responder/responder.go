package responder

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func Respond(code int, resp interface{}, err error) (*events.APIGatewayProxyResponse, error) {
	var body string
	if err != nil {
		log.Println(err)

		body = err.Error()
	} else {
		js, err := json.Marshal(resp)
		if err != nil {
			js = []byte(err.Error())
		}

		body = string(js)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
		Headers: map[string]string{
			"access-control-allow-origin": "*",
		},
	}, nil
}
