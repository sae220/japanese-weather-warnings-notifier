package main

import (
	"errors"
)

// ある地域に出ている気象警報・注意報一つ
type AreaWeatherWarning struct {
	// 気象警報・注意報コード
	code string
	// 気象警報・注意報の状態
	status string
}

// ある地域に出ている気象警報・注意報の配列
type AreaWeatherWarnings []AreaWeatherWarning

// 全国地方公共団体コードに対応する地域に出ている気象警報・注意報を取得する
func FetchAreaWeatherWarnings(areaCode AreaCode) (AreaWeatherWarnings, error) {
	if !areaCode.IsValid() {
		return nil, errors.New("全国地方公共団体コードが無効です")
	}
	if areaCode.IsPrefectureAreaCode() {
		return nil, errors.New("都道府県ではなく市町村のコードを入力してください")
	}
	return nil, nil
}
