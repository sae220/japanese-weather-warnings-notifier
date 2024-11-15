package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambda() error {
	lineBot, err := InitLineBot()
	if err != nil {
		return fmt.Errorf("line bot failed in initialization: %s", err)
	}

	_, err = BroadcastSimpleMessage(lineBot, "Hello World!")
	if err != nil {
		return fmt.Errorf("line bot failed in broadcasting: %s", err)
	}
	return nil
}

func main() {
	lambda.Start(HandleLambda)
}
