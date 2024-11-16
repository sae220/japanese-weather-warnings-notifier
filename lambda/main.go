package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func InitLineBot() (*messaging_api.MessagingApiAPI, error) {
	return messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
}

func BroadcastSimpleMessage(lineBot *messaging_api.MessagingApiAPI, text string) (*map[string]interface{}, error) {
	return lineBot.Broadcast(
		&messaging_api.BroadcastRequest{
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: text,
				},
			},
		},
		"",
	)
}

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
