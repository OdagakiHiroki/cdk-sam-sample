package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"string"`
}

type Event struct {
	events.EventBridgeEvent
	events.CognitoEventUserPoolsPostConfirmation
}

func handler(event Event) (events.LambdaFunctionURLResponse, error) {
	jsonRes, _ := json.Marshal(event)
	fmt.Println(string(jsonRes))
	if event.Source == "aws.scheduler" {
		fmt.Println("called by event bridge scheduler")
		return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "called by event bridge scheduler"}, nil
	}
	msg := "Hello world"
	fmt.Println(msg)
	responseBody, err := json.Marshal(Response{Message: msg})
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: string(responseBody)}, err
}

func main() {
	lambda.Start(handler)
}
