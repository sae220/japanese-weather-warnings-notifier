package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func HandleLambda() error {
	lineBot, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return err
	}

	_, err = lineBot.Broadcast(
		&messaging_api.BroadcastRequest{
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: "Hello world!",
				},
			},
		},
		"",
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(HandleLambda)
}
