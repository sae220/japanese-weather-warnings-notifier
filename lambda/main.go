package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type LineBot struct {
	api *messaging_api.MessagingApiAPI
}

func InitLineBot() (*LineBot, error) {
	api, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return nil, err
	}
	return &LineBot{api}, nil
}

func (lineBot *LineBot) BroadcastSimpleMessage(text string) (*map[string]interface{}, error) {
	return lineBot.api.Broadcast(
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

	_, err = lineBot.BroadcastSimpleMessage("Hello World!")
	if err != nil {
		return fmt.Errorf("line bot failed in broadcasting: %s", err)
	}
	return nil
}

func main() {
	lambda.Start(HandleLambda)
}
