package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambda() error {
	lineBot, err := InitLineBot()
	if err != nil {
		return err
	}

	_, err = BroadcastSimpleMessage(lineBot, "Hello World!")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(HandleLambda)
}
