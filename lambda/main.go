package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"lambda/jpweatherwarnings"
)

func HandleLambda() error {
	lineBot, err := InitLineBot()
	if err != nil {
		return fmt.Errorf("line bot failed in initialization: %s", err)
	}

	areaCode := jpweatherwarnings.AreaCode(os.Getenv("AREA_CODE"))
	_, err = jpweatherwarnings.FetchAreaWeatherWarnings(areaCode)
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
