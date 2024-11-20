package main

import (
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// LINEのbot
type LineBot struct {
	api *messaging_api.MessagingApiAPI
}

// LINEのbotの初期化
func InitLineBot() (*LineBot, error) {
	api, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return nil, err
	}
	return &LineBot{api}, nil
}

// LINEで単純な文字列を全ての友だち向けに送信する
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
