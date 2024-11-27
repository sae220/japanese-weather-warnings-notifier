package jpweatherwarnings

import (
	"errors"
)

// 気象警報・注意報を取得するAPIのレスポンスの形式
type (
	WeatherWarningApiResponse struct {
		AreaTypes AreaTypes `json:"areaTypes"`
	}

	AreaTypes []AreaType

	AreaType struct {
		Areas Areas `json:"areas"`
	}

	Areas []Area

	Area struct {
		Code     AreaCode `json:"code"`
		Warnings Warnings `json:"warnings"`
	}

	Warnings []Warning

	Warning struct {
		Code   string `json:"code"`
		Status string `json:"status"`
	}
)

// 全国地方公共団体コードに対応する地域に出ている気象警報・注意報を取得する
func FetchAreaWeatherWarnings(areaCode AreaCode) (Warnings, error) {
	if !areaCode.IsValid() {
		return nil, errors.New("全国地方公共団体コードが無効です")
	}
	if areaCode.IsPrefectureAreaCode() {
		return nil, errors.New("都道府県ではなく市町村のコードを入力してください")
	}
	return nil, nil
}
