package main

import (
	"os"

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
