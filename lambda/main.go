package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type requestBody struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body requestBody
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	j, _ := json.Marshal(Response{Message: "Hello " + body.Name})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(j),
	}, nil
}
