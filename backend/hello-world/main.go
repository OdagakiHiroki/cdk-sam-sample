package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"string"`
}

func handler(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	msg := "Hello world"
	fmt.Println(msg)
	responseBody, err := json.Marshal(Response{Message: msg})
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: string(responseBody)}, err
}

func main() {
	lambda.Start(handler)
}
