package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
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

// ある地域に出ている気象警報・注意報一つ
type AreaWeatherWarning struct {
	// 気象警報・注意報コード
	code string
	// 気象警報・注意報の状態
	status string
}

// ある地域に出ている気象警報・注意報の配列
type AreaWeatherWarnings []AreaWeatherWarning

// 地域コードに対応する地域に出ている気象警報・注意報を取得する
func FetchAreaWeatherWarnings(areaCode string) (AreaWeatherWarnings, error) {
	return nil, nil
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
