package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Inside copy function")

	return events.APIGatewayProxyResponse{
		Body:       string("responseBody"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
