package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambda() error {
	lineBot, err := InitLineBot()
	if err != nil {
		return fmt.Errorf("line bot failed in initialization: %s", err)
	}

	_, err = FetchAreaWeatherWarnings(os.Getenv("AREA_CODE"))
	if err != nil {
		return fmt.Errorf("failed in fetching weather warnings: %s", err)
	}

	_, err = lineBot.BroadcastSimpleMessage("Hello World!")
	if err != nil {
		return fmt.Errorf("line bot failed in broadcasting: %s", err)
	}
	return nil
}

func main() {
	lambda.Start(HandleLambda)
}
