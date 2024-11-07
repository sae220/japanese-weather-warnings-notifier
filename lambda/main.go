package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambda() (*string, error) {
	message := "Hello World!"
	return &message, nil
}

func main() {
	lambda.Start(HandleLambda)
}
