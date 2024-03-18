package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type User struct {
	Id    string `json:"Id"` // Use uppercase for exported fields
	email string `json:"email"`
	name  string `json:"name"`
}
type InputForm struct {
	Id string `json:"id"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Lambda function called successfully!")
	var inputForm InputForm
	err := json.Unmarshal([]byte(request.Body), &inputForm)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid request body"}, nil
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Specify your desired AWS region
	})
	fmt.Println("session initiated successfully!")
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to create session: %v", err)
	}
	db := dynamodb.New(sess)
	fmt.Println("session created successfully!", db)
	// Define the item to be inserted
	user := User{
		Id:    string(inputForm.Id),
		email: string(inputForm.Id) + "@gmail.com",
		name:  string(inputForm.Id) + "Happy man",
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String("MyTable"), // Replace with your table name
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(user.Id), // Include the ID attribute in the item
			},
			"email": {
				S: aws.String(user.email),
			},
			"name": {
				S: aws.String(user.name),
			},
		},
	}
	fmt.Println("Input object craeted succesfully", input)

	_, err = db.PutItem(input)
	fmt.Println("After input")

	if err != nil {
		log.Printf("Error calling PutItem: %v", err)
		fmt.Println("Final input loggggg", err)
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println("Final input init", err)

	message := "Successfully added item to DynamoDB table!"

	// Use JSON encoding for a structured response
	responseBody, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseBody),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
